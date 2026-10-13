package tests

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/protobuf/proto"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

// SystemUserID owns runtime writes that omit a user ID, as in Nakama.
const SystemUserID = "00000000-0000-0000-0000-000000000000"

type indexFilter func(context.Context, runtime.Logger, *sql.DB, runtime.NakamaModule, *runtime.StorageWrite) bool

// MemoryNakama keeps the real JSON/version/ownership boundary for session tests.
// Objects are identified by collection, key and owner, and values are returned
// in PostgreSQL JSONB text form. The retry method models Nakama 3.41's
// ErrStorageRejectedVersion contract.
type MemoryNakama struct {
	runtime.NakamaModule
	mu           sync.Mutex
	objects      map[string]*api.StorageObject
	indexes      map[string]string
	indexFilters map[string]indexFilter
	sequence     int
	AfterRead    func()
	WriteError   error
	Writes       int
}

func NewMemoryNakama() *MemoryNakama {
	return &MemoryNakama{objects: map[string]*api.StorageObject{}, indexes: map[string]string{}, indexFilters: map[string]indexFilter{}}
}

func owner(userID string) string {
	if userID == "" {
		return SystemUserID
	}
	return userID
}

func storageID(collection, key, userID string) string {
	return collection + "\x00" + key + "\x00" + owner(userID)
}

// jsonbText mimics PostgreSQL JSONB output: object keys ordered by length then
// bytes, with ": " and ", " separators. Invalid JSON is kept to model corruption.
func jsonbText(value string) string {
	decoder := json.NewDecoder(strings.NewReader(value))
	decoder.UseNumber()
	var parsed any
	if err := decoder.Decode(&parsed); err != nil || decoder.More() {
		return value
	}
	var b strings.Builder
	writeJSONB(&b, parsed)
	return b.String()
}

func writeJSONB(b *strings.Builder, value any) {
	switch value := value.(type) {
	case map[string]any:
		keys := make([]string, 0, len(value))
		for key := range value {
			keys = append(keys, key)
		}
		sort.Slice(keys, func(i, j int) bool {
			if len(keys[i]) != len(keys[j]) {
				return len(keys[i]) < len(keys[j])
			}
			return keys[i] < keys[j]
		})
		b.WriteByte('{')
		for i, key := range keys {
			if i > 0 {
				b.WriteString(", ")
			}
			writeJSONB(b, key)
			b.WriteString(": ")
			writeJSONB(b, value[key])
		}
		b.WriteByte('}')
	case []any:
		b.WriteByte('[')
		for i, element := range value {
			if i > 0 {
				b.WriteString(", ")
			}
			writeJSONB(b, element)
		}
		b.WriteByte(']')
	default:
		data, _ := json.Marshal(value)
		b.Write(data)
	}
}

func (n *MemoryNakama) StorageRead(ctx context.Context, reads []*runtime.StorageRead) ([]*api.StorageObject, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	n.mu.Lock()
	out := make([]*api.StorageObject, 0, len(reads))
	for _, r := range reads {
		if obj := n.objects[storageID(r.Collection, r.Key, r.UserID)]; obj != nil {
			clone := proto.Clone(obj).(*api.StorageObject)
			out = append(out, clone)
		}
	}
	hook := n.AfterRead
	n.mu.Unlock()
	if hook != nil {
		hook()
	}
	return out, nil
}
func (n *MemoryNakama) StorageWrite(ctx context.Context, writes []*runtime.StorageWrite) ([]*api.StorageObjectAck, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Writes++
	if n.WriteError != nil {
		return nil, n.WriteError
	}
	for _, w := range writes {
		existing := n.objects[storageID(w.Collection, w.Key, w.UserID)]
		if (w.Version == "*" && existing != nil) || (w.Version != "" && w.Version != "*" && (existing == nil || existing.Version != w.Version)) {
			return nil, runtime.ErrStorageRejectedVersion
		}
	}
	acks := make([]*api.StorageObjectAck, 0, len(writes))
	for _, w := range writes {
		n.sequence++
		version := strconv.Itoa(n.sequence)
		n.objects[storageID(w.Collection, w.Key, w.UserID)] = &api.StorageObject{Collection: w.Collection, Key: w.Key, UserId: owner(w.UserID), Value: jsonbText(w.Value), Version: version}
		acks = append(acks, &api.StorageObjectAck{Collection: w.Collection, Key: w.Key, UserId: owner(w.UserID), Version: version})
	}
	return acks, nil
}
func (n *MemoryNakama) StorageWriteRetry(ctx context.Context, reads []*runtime.StorageRead, fn func([]*api.StorageObject) ([]*runtime.StorageWrite, error), maxRetries int) ([]*api.StorageObjectAck, error) {
	for attempt := 0; attempt <= maxRetries; attempt++ {
		objects, err := n.StorageRead(ctx, reads)
		if err != nil {
			return nil, err
		}
		writes, err := fn(objects)
		if err != nil {
			return nil, err
		}
		acks, err := n.StorageWrite(ctx, writes)
		if errors.Is(err, runtime.ErrStorageRejectedVersion) {
			continue
		}
		return acks, err
	}
	return nil, runtime.ErrStorageWriteExhaustedRetries
}
func (n *MemoryNakama) StorageDelete(ctx context.Context, deletes []*runtime.StorageDelete) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	n.mu.Lock()
	defer n.mu.Unlock()
	for _, d := range deletes {
		obj := n.objects[storageID(d.Collection, d.Key, d.UserID)]
		if d.Version != "" && (obj == nil || obj.Version != d.Version) {
			return errors.New("Storage delete rejected - not found, version check failed, or permission denied.")
		}
	}
	for _, d := range deletes {
		delete(n.objects, storageID(d.Collection, d.Key, d.UserID))
	}
	return nil
}

// page returns matching objects in ID order after cursor.
func (n *MemoryNakama) page(match func(*api.StorageObject) bool, limit int, cursor string) ([]*api.StorageObject, string) {
	ids := make([]string, 0, len(n.objects))
	for id, obj := range n.objects {
		if id > cursor && match(obj) {
			ids = append(ids, id)
		}
	}
	sort.Strings(ids)
	next := ""
	if len(ids) > limit {
		ids = ids[:limit]
		next = ids[len(ids)-1]
	}
	out := make([]*api.StorageObject, 0, len(ids))
	for _, id := range ids {
		out = append(out, proto.Clone(n.objects[id]).(*api.StorageObject))
	}
	return out, next
}

// StorageList models the runtime listing: an empty userID lists every owner.
func (n *MemoryNakama) StorageList(ctx context.Context, callerID, userID, collection string, limit int, cursor string) ([]*api.StorageObject, string, error) {
	if err := ctx.Err(); err != nil {
		return nil, "", err
	}
	n.mu.Lock()
	defer n.mu.Unlock()
	out, next := n.page(func(obj *api.StorageObject) bool {
		return obj.Collection == collection && (userID == "" || obj.UserId == userID)
	}, limit, cursor)
	return out, next, nil
}

// StorageIndexList lists indexed objects of every owner admitted by the index
// filter. Query and order are not modeled.
func (n *MemoryNakama) StorageIndexList(ctx context.Context, callerID, indexName, query string, limit int, order []string, cursor string) (*api.StorageObjects, string, error) {
	if err := ctx.Err(); err != nil {
		return nil, "", err
	}
	n.mu.Lock()
	defer n.mu.Unlock()
	collection, ok := n.indexes[indexName]
	if !ok {
		return nil, "", fmt.Errorf("index %q not found", indexName)
	}
	filter := n.indexFilters[indexName]
	out, next := n.page(func(obj *api.StorageObject) bool {
		if obj.Collection != collection {
			return false
		}
		return filter == nil || filter(ctx, nil, nil, n, &runtime.StorageWrite{Collection: obj.Collection, Key: obj.Key, UserID: obj.UserId, Value: obj.Value, Version: obj.Version})
	}, limit, cursor)
	return &api.StorageObjects{Objects: out}, next, nil
}

func (n *MemoryNakama) MetricsCounterAdd(string, map[string]string, int64)          {}
func (n *MemoryNakama) MetricsTimerRecord(string, map[string]string, time.Duration) {}

// MemoryInitializer registers storage indexes and their filters on a MemoryNakama.
type MemoryInitializer struct {
	runtime.Initializer
	Nakama *MemoryNakama
}

func (i MemoryInitializer) RegisterStorageIndex(name, collection, _ string, _, _ []string, _ int, _ bool) error {
	i.Nakama.mu.Lock()
	defer i.Nakama.mu.Unlock()
	i.Nakama.indexes[name] = collection
	return nil
}
func (i MemoryInitializer) RegisterStorageIndexFilter(name string, fn func(context.Context, runtime.Logger, *sql.DB, runtime.NakamaModule, *runtime.StorageWrite) bool) error {
	i.Nakama.mu.Lock()
	defer i.Nakama.mu.Unlock()
	i.Nakama.indexFilters[name] = fn
	return nil
}
