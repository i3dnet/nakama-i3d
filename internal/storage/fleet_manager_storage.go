package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	StorageI3dIndex               = "_i3D_instances_idx"
	StorageI3dInstancesCollection = "_i3D_instances"
)

var ErrSessionNotFound = errors.New("session not found")
var ErrCorruptSession = errors.New("corrupt stored session")

type FleetManagerStorage interface {
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

func (fms *FleetManagerStorageService) UpdateStorageGameSession(ctx context.Context, instances []*runtime.InstanceInfo) error {
	storageWrites := make([]*runtime.StorageWrite, 0, len(instances))
	for _, instance := range instances {
		v, err := json.Marshal(instance)
		if err != nil {
			fms.logger.WithField("error", err.Error()).Error("failed to marshal instance")
			return err
		}

		storageWrites = append(storageWrites, &runtime.StorageWrite{
			Collection: StorageI3dInstancesCollection,
			Key:        instance.Id,
			Value:      string(v),
		})
	}

	if _, err := fms.nk.StorageWrite(ctx, storageWrites); err != nil {
		fms.logger.WithField("error", err.Error()).Error("failed to write storage")
		return err
	}
	return nil
}

func (fms *FleetManagerStorageService) DeleteStorageGameSession(ctx context.Context, ids []string) error {
	storageDeletes := make([]*runtime.StorageDelete, 0, len(ids))
	for _, key := range ids {
		storageDeletes = append(storageDeletes, &runtime.StorageDelete{
			Collection: StorageI3dInstancesCollection,
			Key:        key,
		})
	}

	if err := fms.nk.StorageDelete(ctx, storageDeletes); err != nil {
		fms.logger.WithField("error", err.Error()).Error("failed to delete storage")
		return err
	}

	return nil
}
