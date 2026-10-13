package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/internal/tests"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// Any authenticated client may write this collection under its own user ID.
const playerUserID = "8f1c2b7e-4d5a-4c3b-9e2f-1a2b3c4d5e6f"

func memoryService(t *testing.T) (*FleetManagerStorageService, *tests.MemoryNakama) {
	t.Helper()
	nk := tests.NewMemoryNakama()
	service, err := NewFleetManagerStorageService(nk, tests.MemoryInitializer{Nakama: nk}, tests.NewMockLogger())
	require.NoError(t, err)
	return service, nk
}
func instanceValue(t *testing.T, instance *runtime.InstanceInfo) string {
	t.Helper()
	data, err := json.Marshal(instance)
	require.NoError(t, err)
	return string(data)
}

type storageNakama struct {
	runtime.NakamaModule
	read  func(context.Context, []*runtime.StorageRead) ([]*api.StorageObject, error)
	index func(context.Context, string, string, string, int, []string, string) (*api.StorageObjects, string, error)
	list  func(context.Context, string, string, string, int, string) ([]*api.StorageObject, string, error)
}

func (n storageNakama) StorageRead(ctx context.Context, r []*runtime.StorageRead) ([]*api.StorageObject, error) {
	return n.read(ctx, r)
}
func (n storageNakama) StorageIndexList(ctx context.Context, caller, index, query string, limit int, order []string, cursor string) (*api.StorageObjects, string, error) {
	return n.index(ctx, caller, index, query, limit, order, cursor)
}
func (n storageNakama) StorageList(ctx context.Context, caller, user, collection string, limit int, cursor string) ([]*api.StorageObject, string, error) {
	return n.list(ctx, caller, user, collection, limit, cursor)
}
func TestReadMissingOrCorruptSessionReturnsError(t *testing.T) {
	for _, value := range []string{"missing", "null", "{", "{}", `{"id":"different"}`} {
		t.Run(value, func(t *testing.T) {
			nk := storageNakama{read: func(context.Context, []*runtime.StorageRead) ([]*api.StorageObject, error) {
				if value == "missing" {
					return nil, nil
				}
				return []*api.StorageObject{{Key: "instance", Value: value}}, nil
			}}
			service := &FleetManagerStorageService{nk: nk, logger: tests.NewMockLogger()}
			require.NotPanics(t, func() {
				got, err := service.GetGameSessionFromStorage(context.Background(), "instance")
				require.Error(t, err)
				require.Nil(t, got)
			})
		})
	}
}
func TestReadPreservesStorageFailure(t *testing.T) {
	failure := errors.New("storage unavailable")
	service := &FleetManagerStorageService{nk: storageNakama{read: func(context.Context, []*runtime.StorageRead) ([]*api.StorageObject, error) { return nil, failure }}, logger: tests.NewMockLogger()}
	_, err := service.GetGameSessionFromStorage(context.Background(), "instance")
	require.ErrorIs(t, err, failure)
}
func TestStorageListingForwardsCursorAndOrder(t *testing.T) {
	data, _ := json.Marshal(&runtime.InstanceInfo{Id: "instance"})
	service := &FleetManagerStorageService{logger: tests.NewMockLogger(), nk: storageNakama{index: func(_ context.Context, caller, index, query string, limit int, order []string, cursor string) (*api.StorageObjects, string, error) {
		require.Empty(t, caller)
		require.Equal(t, StorageI3dIndex, index)
		require.Equal(t, "+value.metadata.mode:arena", query)
		require.Equal(t, 7, limit)
		require.Equal(t, []string{"-value.player_count"}, order)
		require.Equal(t, "previous", cursor)
		return &api.StorageObjects{Objects: []*api.StorageObject{{Key: "instance", Value: string(data)}}}, "next", nil
	}}}
	got, cursor, err := service.ListGameSessionsFromStorage(context.Background(), "+value.metadata.mode:arena", 7, []string{"-value.player_count"}, "previous")
	require.NoError(t, err)
	require.Len(t, got, 1)
	require.Equal(t, "next", cursor)
}

func TestSnapshotIgnoresRecordsNotOwnedBySystem(t *testing.T) {
	ctx := context.Background()
	service, nk := memoryService(t)
	require.NoError(t, service.CreateGameSession(ctx, &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED"}, "123", nil))
	_, err := nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{Collection: StorageI3dInstancesCollection, Key: "corrupt", UserID: playerUserID, Value: "{}"},
		{Collection: StorageI3dInstancesCollection, Key: "instance", UserID: playerUserID, Value: instanceValue(t, &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED"})},
	})
	require.NoError(t, err)
	snapshot, err := service.SnapshotGameSessions(ctx)
	require.NoError(t, err)
	require.Len(t, snapshot, 1)
	require.Equal(t, tests.SystemUserID, snapshot[0].UserId)

	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{{Collection: StorageI3dInstancesCollection, Key: "corrupt", Value: "{}"}})
	require.NoError(t, err)
	_, err = service.SnapshotGameSessions(ctx)
	require.ErrorIs(t, err, ErrCorruptSession, "corrupt system records must still fail the snapshot")
}
func TestSnapshotListsSystemRecordsAndSkipsOtherOwners(t *testing.T) {
	service := &FleetManagerStorageService{logger: tests.NewMockLogger(), nk: storageNakama{list: func(_ context.Context, caller, user, collection string, _ int, _ string) ([]*api.StorageObject, string, error) {
		require.Empty(t, caller)
		require.Equal(t, tests.SystemUserID, user)
		require.Equal(t, StorageI3dInstancesCollection, collection)
		return []*api.StorageObject{
			{Key: "instance", UserId: tests.SystemUserID, Value: instanceValue(t, &runtime.InstanceInfo{Id: "instance"}), Version: "1"},
			{Key: "unowned", Value: instanceValue(t, &runtime.InstanceInfo{Id: "unowned"}), Version: "2"},
			{Key: "instance", UserId: playerUserID, Value: "{}", Version: "3"},
		}, "", nil
	}}}
	snapshot, err := service.SnapshotGameSessions(context.Background())
	require.NoError(t, err)
	require.Len(t, snapshot, 2)
	require.Equal(t, "1", snapshot[0].Version)
	require.Equal(t, "2", snapshot[1].Version)
}
func TestListingIgnoresRecordsNotOwnedBySystem(t *testing.T) {
	ctx := context.Background()
	entries := []*api.StorageObject{
		{Key: "instance", UserId: tests.SystemUserID, Value: instanceValue(t, &runtime.InstanceInfo{Id: "instance"})},
		{Key: "corrupt", UserId: playerUserID, Value: "{}"},
		{Key: "forged", UserId: playerUserID, Value: instanceValue(t, &runtime.InstanceInfo{Id: "forged", ConnectionInfo: &runtime.ConnectionInfo{IpAddress: "203.0.113.1", Port: 7777}})},
	}
	service := &FleetManagerStorageService{logger: tests.NewMockLogger(), nk: storageNakama{index: func(context.Context, string, string, string, int, []string, string) (*api.StorageObjects, string, error) {
		return &api.StorageObjects{Objects: entries}, "", nil
	}}}
	got, _, err := service.ListGameSessionsFromStorage(ctx, "*", 10, nil, "")
	require.NoError(t, err)
	require.Len(t, got, 1)
	require.Equal(t, "instance", got[0].Id)

	entries = append(entries, &api.StorageObject{Key: "broken", UserId: tests.SystemUserID, Value: "{}"})
	_, _, err = service.ListGameSessionsFromStorage(ctx, "*", 10, nil, "")
	require.ErrorIs(t, err, ErrCorruptSession, "corrupt system records must still fail the page")
}
func TestStorageIndexAdmitsOnlySystemWrites(t *testing.T) {
	ctx := context.Background()
	service, nk := memoryService(t)
	require.NoError(t, service.CreateGameSession(ctx, &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED"}, "123", nil))
	_, err := nk.StorageWrite(ctx, []*runtime.StorageWrite{
		{Collection: StorageI3dInstancesCollection, Key: "forged", UserID: playerUserID, Value: instanceValue(t, &runtime.InstanceInfo{Id: "forged"})},
		{Collection: StorageI3dInstancesCollection, Key: "corrupt", UserID: playerUserID, Value: "{}"},
	})
	require.NoError(t, err)
	indexed, _, err := nk.StorageIndexList(ctx, "", StorageI3dIndex, "*", 10, nil, "")
	require.NoError(t, err)
	require.Len(t, indexed.GetObjects(), 1)
	require.Equal(t, tests.SystemUserID, indexed.GetObjects()[0].UserId)
}
func TestUnchangedProviderViewDoesNotRewriteStoredSession(t *testing.T) {
	provider := func() *runtime.InstanceInfo {
		return &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", CreateTime: time.Unix(100, 0).UTC(), ConnectionInfo: &runtime.ConnectionInfo{IpAddress: "192.0.2.1", Port: 7777}, Metadata: map[string]any{"map": "arena", "rank": 3}}
	}
	for name, refresh := range map[string]func(*FleetManagerStorageService, *api.StorageObject, *runtime.InstanceInfo) error{
		"reconcile": func(service *FleetManagerStorageService, snapshot *api.StorageObject, incoming *runtime.InstanceInfo) error {
			return service.ReconcileGameSession(context.Background(), snapshot, incoming, "123", "")
		},
		"provider page": func(service *FleetManagerStorageService, _ *api.StorageObject, incoming *runtime.InstanceInfo) error {
			return service.UpdateStorageGameSession(context.Background(), []*runtime.InstanceInfo{incoming})
		},
	} {
		t.Run(name, func(t *testing.T) {
			service, nk := memoryService(t)
			allocated := provider()
			allocated.Metadata["i3d_max_players"] = 4
			require.NoError(t, service.CreateGameSession(context.Background(), allocated, "123", []string{"user"}))
			snapshot, err := service.GetGameSessionSnapshot(context.Background(), "instance")
			require.NoError(t, err)
			require.Contains(t, snapshot.Value, `": `, "storage returns JSONB text, not Go's encoding")
			before := nk.Writes
			incoming := provider()
			require.NoError(t, refresh(service, snapshot, incoming))
			require.Equal(t, before, nk.Writes, "an unchanged record must keep its version")
			require.EqualValues(t, 4, incoming.Metadata["i3d_max_players"])
		})
	}
}
