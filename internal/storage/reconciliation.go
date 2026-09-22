package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	"time"
)

// SnapshotGameSessions uses authoritative storage, not the eventually indexed view.
// No caller may act on a partial or corrupt snapshot.
func (fms *FleetManagerStorageService) SnapshotGameSessions(ctx context.Context) ([]*api.StorageObject, error) {
	var result []*api.StorageObject
	cursor := ""
	seen := map[string]bool{}
	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		objects, next, err := fms.nk.StorageList(ctx, "", "", StorageI3dInstancesCollection, 100, cursor)
		if err != nil {
			return nil, err
		}
		for _, obj := range objects {
			if _, err := decodeRecord(obj.Value, obj.Key); err != nil {
				return nil, err
			}
			if obj.Version == "" {
				return nil, fmt.Errorf("%w: missing version", ErrCorruptSession)
			}
		}
		result = append(result, objects...)
		if next == "" {
			return result, nil
		}
		if seen[next] {
			return nil, fmt.Errorf("storage cursor repeated")
		}
		seen[next] = true
		cursor = next
	}
}

// SnapshotInScope reports explicit ownership; legacy/foreign records are not
// removed based only on absence from a differently scoped provider query.
func SnapshotInScope(obj *api.StorageObject, applicationID, fleetID string) (bool, time.Time, error) {
	record, err := decodeRecord(obj.Value, obj.Key)
	if err != nil {
		return false, time.Time{}, err
	}
	fleet, _ := record.Metadata["i3d_fleet_id"].(string)
	return record.Local.ApplicationID != "" && record.Local.ApplicationID == applicationID && (fleetID == "" || fleet == fleetID), record.Local.AllocatedAt, nil
}

// ReconcileGameSession applies exactly the version observed before the provider
// scan. It deliberately does not retry conflicts on a newer allocation/admission.
// A nil incoming value removes only that snapshot; a nil snapshot inserts only.
func (fms *FleetManagerStorageService) ReconcileGameSession(ctx context.Context, snapshot *api.StorageObject, incoming *runtime.InstanceInfo, applicationID, fleetID string) error {
	if snapshot == nil && incoming == nil {
		return nil
	}
	if incoming == nil {
		if snapshot.Version == "" {
			return fmt.Errorf("%w: missing version", ErrCorruptSession)
		}
		err := fms.nk.StorageDelete(ctx, []*runtime.StorageDelete{{Collection: StorageI3dInstancesCollection, Key: snapshot.Key, Version: snapshot.Version}})
		if err != nil {
			// Nakama 3.41 delete conflicts return a plain error, unlike writes.
			// Prove the snapshot changed rather than parsing an error string.
			current, readErr := fms.GetGameSessionSnapshot(ctx, snapshot.Key)
			if readErr == nil && (current == nil || current.Version != snapshot.Version) {
				return runtime.ErrStorageRejectedVersion
			}
		}
		return err
	}
	version := "*"
	record := &sessionRecord{InstanceInfo: &runtime.InstanceInfo{Id: incoming.Id}}
	if snapshot != nil {
		var err error
		record, err = decodeRecord(snapshot.Value, snapshot.Key)
		if err != nil {
			return err
		}
		if snapshot.Version == "" || snapshot.Key != incoming.Id {
			return fmt.Errorf("%w: invalid snapshot", ErrCorruptSession)
		}
		version = snapshot.Version
	}
	newGeneration := record.Local.Generation == "" || (!record.CreateTime.IsZero() && !incoming.CreateTime.IsZero() && !record.CreateTime.Equal(incoming.CreateTime))
	if err := MergeProviderInstance(record.InstanceInfo, incoming); err != nil {
		return err
	}
	if newGeneration {
		record.Local = localState{Generation: uuid.NewString(), ApplicationID: applicationID, AllocatedAt: time.Now().UTC(), JoinedUsers: map[string]bool{}}
	}
	if applicationID != "" {
		record.Local.ApplicationID = applicationID
	}
	if fleetID != "" {
		record.Metadata["i3d_fleet_id"] = fleetID
	}
	value, err := json.Marshal(record)
	if err != nil {
		return err
	}
	if snapshot != nil && string(value) == snapshot.Value {
		*incoming = *record.InstanceInfo
		return nil
	}
	_, err = fms.nk.StorageWrite(ctx, []*runtime.StorageWrite{{Collection: StorageI3dInstancesCollection, Key: incoming.Id, Version: version, Value: string(value), PermissionRead: 0, PermissionWrite: 0}})
	if err == nil {
		*incoming = *record.InstanceInfo
	}
	return err
}
func (fms *FleetManagerStorageService) GetGameSessionSnapshot(ctx context.Context, id string) (*api.StorageObject, error) {
	objects, err := fms.nk.StorageRead(ctx, []*runtime.StorageRead{{Collection: StorageI3dInstancesCollection, Key: id}})
	if err != nil {
		return nil, err
	}
	if len(objects) == 0 {
		return nil, nil
	}
	obj := objects[0]
	if _, err := decodeRecord(obj.Value, id); err != nil {
		return nil, err
	}
	if obj.Version == "" {
		return nil, fmt.Errorf("%w: missing version", ErrCorruptSession)
	}
	return obj, nil
}

// SameAllocation reports a proven allocation identity, independently of storage
// versions changed by joins/refreshes. Legacy records cannot prove this identity.
func SameAllocation(left, right *api.StorageObject) (bool, error) {
	if left == nil || right == nil || left.Key != right.Key {
		return false, nil
	}
	a, err := decodeRecord(left.Value, left.Key)
	if err != nil {
		return false, err
	}
	b, err := decodeRecord(right.Value, right.Key)
	if err != nil {
		return false, err
	}
	return a.Local.Generation != "" && a.Local.Generation == b.Local.Generation, nil
}
