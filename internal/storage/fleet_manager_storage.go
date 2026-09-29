package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	"strings"
	"time"
)

const (
	StorageI3dIndex               = "_i3D_instances_idx"
	StorageI3dInstancesCollection = "_i3D_instances"
)

var ErrSessionNotFound = errors.New("session not found")
var ErrCorruptSession = errors.New("corrupt stored session")
var ErrStaleSession = errors.New("stored allocation has been invalidated")

var errNoSessionUpdates = errors.New("no local sessions to refresh")

// staleProviderGeneration prevents admission while retaining cleanup ownership.
const staleProviderGeneration = "STALE"

type FleetManagerStorage interface {
	GetGameSessionSnapshot(context.Context, string) (*api.StorageObject, error)
	SnapshotGameSessions(context.Context) ([]*api.StorageObject, error)
	ReconcileGameSession(context.Context, *api.StorageObject, *runtime.InstanceInfo, string, string) error
	CreateGameSession(ctx context.Context, instance *runtime.InstanceInfo, applicationID string, userIDs []string) error
	MutateGameSession(ctx context.Context, id string, create bool, fn func(*runtime.InstanceInfo, map[string]bool) error) (*runtime.InstanceInfo, error)
	GetGameSessionFromStorage(ctx context.Context, id string) (*runtime.InstanceInfo, error)
	ListGameSessionsFromStorage(ctx context.Context, query string, limit int, order []string, cursor string) ([]*runtime.InstanceInfo, string, error)
	UpdateStorageGameSession(ctx context.Context, instances []*runtime.InstanceInfo) error
	DeleteStorageGameSession(ctx context.Context, ids []string) error
}

type FleetManagerStorageService struct {
	nk          runtime.NakamaModule
	initializer runtime.Initializer
	logger      runtime.Logger
}

func NewFleetManagerStorageService(nk runtime.NakamaModule, initializer runtime.Initializer, logger runtime.Logger) (*FleetManagerStorageService, error) {
	if err := initializer.RegisterStorageIndex(StorageI3dIndex, StorageI3dInstancesCollection, "", []string{"id", "create_time", "player_count", "metadata"}, []string{"create_time", "player_count"}, 1_000_000, false); err != nil {
		return nil, err
	}

	return &FleetManagerStorageService{
		nk:          nk,
		logger:      logger,
		initializer: initializer,
	}, nil
}

func (fms *FleetManagerStorageService) GetGameSessionFromStorage(ctx context.Context, id string) (*runtime.InstanceInfo, error) {
	objects, err := fms.nk.StorageRead(ctx, []*runtime.StorageRead{{
		Collection: StorageI3dInstancesCollection,
		Key:        id,
	}})

	if err != nil {
		fms.logger.WithField("error", err.Error()).Error("failed to read storage")
		return nil, err
	}

	if len(objects) == 0 {
		return nil, ErrSessionNotFound
	}
	return decodeInstance(objects[0].Value, id)
}

func decodeInstance(value, id string) (*runtime.InstanceInfo, error) {
	var instance *runtime.InstanceInfo
	if err := json.Unmarshal([]byte(value), &instance); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCorruptSession, err)
	}
	if instance == nil || instance.Id == "" || instance.Id != id {
		return nil, fmt.Errorf("%w: invalid instance ID", ErrCorruptSession)
	}
	return instance, nil
}

func (fms *FleetManagerStorageService) ListGameSessionsFromStorage(ctx context.Context, query string, limit int, order []string, cursor string) ([]*runtime.InstanceInfo, string, error) {
	if limit == 0 {
		limit = 100
	}
	entries, nextCursor, err := fms.nk.StorageIndexList(ctx, "", StorageI3dIndex, query, limit, order, cursor)
	if err != nil {
		return nil, "", err
	}
	results := make([]*runtime.InstanceInfo, 0, len(entries.GetObjects()))
	for _, entry := range entries.GetObjects() {
		info, err := decodeInstance(entry.Value, entry.Key)
		if err != nil {
			return nil, "", err
		}
		results = append(results, info)
	}
	return results, nextCursor, nil
}

// localState is stored beside InstanceInfo, keeping existing indexed fields intact.
type localState struct {
	Generation    string          `json:"generation,omitempty"`
	ApplicationID string          `json:"application_id,omitempty"`
	AllocatedAt   time.Time       `json:"allocated_at,omitempty"`
	JoinedUsers   map[string]bool `json:"joined_users,omitempty"`
}
type sessionRecord struct {
	*runtime.InstanceInfo
	Local localState `json:"_i3d,omitempty"`
}

func decodeRecord(value, id string) (*sessionRecord, error) {
	instance, err := decodeInstance(value, id)
	if err != nil {
		return nil, err
	}
	var record sessionRecord
	if err = json.Unmarshal([]byte(value), &record); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCorruptSession, err)
	}
	record.InstanceInfo = instance
	if record.Local.JoinedUsers == nil {
		record.Local.JoinedUsers = make(map[string]bool)
	}
	return &record, nil
}
func cloneInstance(instance *runtime.InstanceInfo) (*runtime.InstanceInfo, error) {
	if instance == nil || instance.Id == "" {
		return nil, fmt.Errorf("instance ID is required")
	}
	data, err := json.Marshal(instance)
	if err != nil {
		return nil, err
	}
	return decodeInstance(string(data), instance.Id)
}

// MergeProviderInstance refreshes provider fields while retaining local capacity/admissions.
// A changed provider creation time identifies a new instance generation.
func MergeProviderInstance(stored, incoming *runtime.InstanceInfo) error {
	if stored.Status == staleProviderGeneration {
		return ErrStaleSession
	}
	refreshed, err := cloneInstance(incoming)
	if err != nil {
		return err
	}
	if refreshed.Metadata == nil {
		refreshed.Metadata = make(map[string]any)
	}
	for key := range refreshed.Metadata {
		if strings.HasPrefix(key, "i3d_") {
			delete(refreshed.Metadata, key)
		}
	}
	sameGeneration := stored.CreateTime.IsZero() || refreshed.CreateTime.IsZero() || stored.CreateTime.Equal(refreshed.CreateTime)
	if sameGeneration {
		for key, value := range stored.Metadata {
			if strings.HasPrefix(key, "i3d_") {
				refreshed.Metadata[key] = value
			}
		}
		if _, owned := stored.Metadata["i3d_max_players"]; owned {
			refreshed.PlayerCount = stored.PlayerCount
		}
	}
	*stored = *refreshed
	return nil
}

func (fms *FleetManagerStorageService) mutate(ctx context.Context, id string, create bool, fn func(*sessionRecord) error) (*runtime.InstanceInfo, error) {
	if id == "" {
		return nil, fmt.Errorf("instance ID is required")
	}
	var committed *runtime.InstanceInfo
	_, err := fms.nk.StorageWriteRetry(ctx, []*runtime.StorageRead{{Collection: StorageI3dInstancesCollection, Key: id}}, func(objects []*api.StorageObject) ([]*runtime.StorageWrite, error) {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		version := "*"
		record := &sessionRecord{InstanceInfo: &runtime.InstanceInfo{Id: id, Metadata: map[string]any{}}, Local: localState{JoinedUsers: map[string]bool{}}}
		if len(objects) == 0 {
			if !create {
				return nil, ErrSessionNotFound
			}
		} else {
			var err error
			record, err = decodeRecord(objects[0].Value, id)
			if err != nil {
				return nil, err
			}
			version = objects[0].Version
			if version == "" {
				return nil, fmt.Errorf("%w: missing storage version", ErrCorruptSession)
			}
		}
		if err := fn(record); err != nil {
			return nil, err
		}
		if record.InstanceInfo == nil || record.Id != id {
			return nil, fmt.Errorf("mutation changed instance ID")
		}
		value, err := json.Marshal(record)
		if err != nil {
			return nil, err
		}
		committed = record.InstanceInfo
		return []*runtime.StorageWrite{{Collection: StorageI3dInstancesCollection, Key: id, Value: string(value), Version: version, PermissionRead: 0, PermissionWrite: 0}}, nil
	}, 5)
	if err != nil {
		return nil, err
	}
	return committed, nil
}

// MutateGameSession reruns fn on a fresh snapshot after storage version conflicts.
// fn must be free of external side effects and replace any result derived per attempt.
func (fms *FleetManagerStorageService) MutateGameSession(ctx context.Context, id string, create bool, fn func(*runtime.InstanceInfo, map[string]bool) error) (*runtime.InstanceInfo, error) {
	return fms.mutate(ctx, id, create, func(record *sessionRecord) error { return fn(record.InstanceInfo, record.Local.JoinedUsers) })
}
func (fms *FleetManagerStorageService) CreateGameSession(ctx context.Context, instance *runtime.InstanceInfo, applicationID string, userIDs []string) error {
	copied, err := cloneInstance(instance)
	if err != nil {
		return err
	}
	generation := uuid.NewString()
	allocatedAt := time.Now().UTC()
	_, err = fms.mutate(ctx, instance.Id, true, func(record *sessionRecord) error {
		record.InstanceInfo = copied
		record.Local = localState{Generation: generation, ApplicationID: applicationID, AllocatedAt: allocatedAt, JoinedUsers: map[string]bool{}}
		for _, id := range userIDs {
			record.Local.JoinedUsers[id] = true
		}
		return nil
	})
	return err
}

func (fms *FleetManagerStorageService) UpdateStorageGameSession(ctx context.Context, instances []*runtime.InstanceInfo) error {
	if len(instances) == 0 {
		return nil
	}
	reads := make([]*runtime.StorageRead, 0, len(instances))
	ids := make(map[string]bool, len(instances))
	for _, incoming := range instances {
		if incoming == nil || incoming.Id == "" || ids[incoming.Id] {
			return fmt.Errorf("provider page has a missing or duplicate instance ID")
		}
		ids[incoming.Id] = true
		reads = append(reads, &runtime.StorageRead{Collection: StorageI3dInstancesCollection, Key: incoming.Id})
	}
	var committed []*runtime.InstanceInfo
	_, err := fms.nk.StorageWriteRetry(ctx, reads, func(objects []*api.StorageObject) ([]*runtime.StorageWrite, error) {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		byID := make(map[string]*api.StorageObject, len(objects))
		for _, object := range objects {
			byID[object.Key] = object
		}
		writes := make([]*runtime.StorageWrite, 0, len(instances))
		committed = make([]*runtime.InstanceInfo, 0, len(instances))
		for _, incoming := range instances {
			// Provider metadata cannot reconstruct local admission capacity.
			// Return a sanitized provider view, but never create a cache record.
			providerView := &runtime.InstanceInfo{Id: incoming.Id}
			if err := MergeProviderInstance(providerView, incoming); err != nil {
				return nil, err
			}
			object := byID[incoming.Id]
			if object == nil {
				committed = append(committed, providerView)
				continue
			}
			record, err := decodeRecord(object.Value, incoming.Id)
			if err != nil {
				return nil, err
			}
			if object.Version == "" {
				return nil, fmt.Errorf("%w: missing storage version", ErrCorruptSession)
			}
			changedGeneration := !record.CreateTime.IsZero() && !incoming.CreateTime.IsZero() && !record.CreateTime.Equal(incoming.CreateTime)
			if record.Status == staleProviderGeneration {
				committed = append(committed, providerView)
				continue
			}
			if changedGeneration && incoming.CreateTime.Before(record.CreateTime) {
				committed = append(committed, record.InstanceInfo)
				continue
			}
			if changedGeneration {
				// Invalidate the old admission state atomically with the page.
				// Keep its identity/ownership for conditional reconciliation;
				// do not adopt the unknown replacement's capacity or generation.
				record.Status = staleProviderGeneration
				committed = append(committed, providerView)
			} else {
				if err := MergeProviderInstance(record.InstanceInfo, incoming); err != nil {
					return nil, err
				}
				committed = append(committed, record.InstanceInfo)
			}
			value, err := json.Marshal(record)
			if err != nil {
				return nil, err
			}
			writes = append(writes, &runtime.StorageWrite{Collection: StorageI3dInstancesCollection, Key: incoming.Id, Value: string(value), Version: object.Version, PermissionRead: 0, PermissionWrite: 0})
		}
		if len(writes) == 0 {
			// Avoid issuing an empty write transaction for provider-only pages.
			return nil, errNoSessionUpdates
		}
		return writes, nil
	}, 5)
	if err != nil && !errors.Is(err, errNoSessionUpdates) {
		return err
	}
	// Publish the merged page only after the complete transaction succeeds.
	for i, incoming := range instances {
		*incoming = *committed[i]
	}
	return nil
}
func (fms *FleetManagerStorageService) DeleteStorageGameSession(ctx context.Context, ids []string) error {
	for _, id := range ids {
		objects, err := fms.nk.StorageRead(ctx, []*runtime.StorageRead{{Collection: StorageI3dInstancesCollection, Key: id}})
		if err != nil {
			return err
		}
		if len(objects) == 0 {
			continue
		}
		if objects[0].Version == "" {
			return fmt.Errorf("%w: missing storage version", ErrCorruptSession)
		}
		if err = fms.nk.StorageDelete(ctx, []*runtime.StorageDelete{{Collection: StorageI3dInstancesCollection, Key: id, Version: objects[0].Version}}); err != nil {
			return err
		}
	}
	return nil
}
