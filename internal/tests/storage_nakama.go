package tests

import (
	"context"
	"errors"
	"google.golang.org/protobuf/proto"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
)

// MemoryNakama keeps the real JSON/version boundary for session tests.
// The retry method models Nakama 3.41's ErrStorageRejectedVersion contract.
type MemoryNakama struct {
	runtime.NakamaModule
	mu         sync.Mutex
	objects    map[string]*api.StorageObject
	sequence   int
	AfterRead  func()
	WriteError error
	Writes     int
}

func NewMemoryNakama() *MemoryNakama { return &MemoryNakama{objects: map[string]*api.StorageObject{}} }
func (n *MemoryNakama) StorageRead(ctx context.Context, reads []*runtime.StorageRead) ([]*api.StorageObject, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	n.mu.Lock()
	out := make([]*api.StorageObject, 0, len(reads))
	for _, r := range reads {
		if obj := n.objects[r.Key]; obj != nil {
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
		existing := n.objects[w.Key]
		if (w.Version == "*" && existing != nil) || (w.Version != "" && w.Version != "*" && (existing == nil || existing.Version != w.Version)) {
			return nil, runtime.ErrStorageRejectedVersion
		}
	}
	acks := make([]*api.StorageObjectAck, 0, len(writes))
	for _, w := range writes {
		n.sequence++
		version := strconv.Itoa(n.sequence)
		n.objects[w.Key] = &api.StorageObject{Collection: w.Collection, Key: w.Key, Value: w.Value, Version: version}
		acks = append(acks, &api.StorageObjectAck{Collection: w.Collection, Key: w.Key, Version: version})
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
		obj := n.objects[d.Key]
		if d.Version != "" && (obj == nil || obj.Version != d.Version) {
			return errors.New("Storage delete rejected - not found, version check failed, or permission denied.")
		}
	}
	for _, d := range deletes {
		delete(n.objects, d.Key)
	}
	return nil
}
func (n *MemoryNakama) StorageList(ctx context.Context, callerID, userID, collection string, limit int, cursor string) ([]*api.StorageObject, string, error) {
	if err := ctx.Err(); err != nil {
		return nil, "", err
	}
	n.mu.Lock()
	defer n.mu.Unlock()
	keys := make([]string, 0, len(n.objects))
	for key, obj := range n.objects {
		if obj.Collection == collection && key > cursor {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)
	next := ""
	if len(keys) > limit {
		keys = keys[:limit]
		next = keys[len(keys)-1]
	}
	out := make([]*api.StorageObject, 0, len(keys))
	for _, key := range keys {
		out = append(out, proto.Clone(n.objects[key]).(*api.StorageObject))
	}
	return out, next, nil
}

func (n *MemoryNakama) MetricsCounterAdd(string, map[string]string, int64)          {}
func (n *MemoryNakama) MetricsTimerRecord(string, map[string]string, time.Duration) {}
