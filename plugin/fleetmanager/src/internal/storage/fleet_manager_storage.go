package storage

import (
	"context"
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
)

const (
	StorageI3dIndex               = "_i3D_instances_idx"
	StorageI3dInstancesCollection = "_i3D_instances"
)

type FleetManagerStorage interface {
	GetGameSessionFromStorage(ctx context.Context, id string) (*runtime.InstanceInfo, error)
	ListGameSessionsFromStorage(ctx context.Context, query string, limit int, order []string, cursor string) ([]*runtime.InstanceInfo, error)
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
		fms.logger.WithField("objects", objects).WithField("id", id).Error("session not found in storage")
	}

	object := objects[0]
	var instance *runtime.InstanceInfo
	if err = json.Unmarshal([]byte(object.Value), &instance); err != nil {
		fms.logger.WithField("error", err.Error()).Error("failed to unmarshal instance")
	}

	return instance, nil
}

func (fms *FleetManagerStorageService) ListGameSessionsFromStorage(ctx context.Context, query string, limit int, order []string, cursor string) ([]*runtime.InstanceInfo, error) {
	entries, _, err := fms.nk.StorageIndexList(ctx, "", StorageI3dIndex, query, limit, []string{"player_count", "-create_time"}, "")
	if err != nil {
		fms.logger.WithField("error", err.Error()).Error("failed to list instances from storage")
		return nil, err
	}

	results := make([]*runtime.InstanceInfo, 0)
	for _, entry := range entries.GetObjects() {
		var info *runtime.InstanceInfo
		if err = json.Unmarshal([]byte(entry.Value), &info); err != nil {
			return nil, err
		}
		results = append(results, info)
	}

	return results, nil
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
