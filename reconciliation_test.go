package fleetmanager

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/config"
	"github.com/i3dnet/nakama-i3d/internal/clients"
	"github.com/i3dnet/nakama-i3d/internal/storage"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sync/atomic"
	"testing"
	"time"
)

func ownedSession(t *testing.T, fm *I3dFleetManager, id string) {
	t.Helper()
	require.NoError(t, fm.storage.CreateGameSession(context.Background(), &runtime.InstanceInfo{Id: id, Status: "ALLOCATED", Metadata: map[string]any{MaxPlayers: 4}}, fm.cfg.ApplicationId, nil))
}
func TestReconciliationRequiresTwoCompleteScans(t *testing.T) {
	fm, _, client := sessionFixture(t, 0, 4)
	fm.cfg.ApplicationId = "123"
	ownedSession(t, fm, "instance")
	ctx := context.Background()
	now := time.Now().Add(time.Hour)
	client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 100, "").Return(&clients.ApplicationInstanceListResponse{NextCursor: "second"}, nil).Times(2)
	client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 100, "second").Return(&clients.ApplicationInstanceListResponse{}, nil).Times(2)
	missing, err := fm.reconcileOnce(ctx, nil, now)
	require.NoError(t, err)
	_, err = fm.storage.GetGameSessionFromStorage(ctx, "instance")
	require.NoError(t, err)
	_, err = fm.reconcileOnce(ctx, missing, now)
	require.NoError(t, err)
	_, err = fm.storage.GetGameSessionFromStorage(ctx, "instance")
	require.ErrorIs(t, err, storage.ErrSessionNotFound)
}
func TestFailedLaterPageCannotDeleteOrRefresh(t *testing.T) {
	fm, _, client := sessionFixture(t, 1, 4)
	fm.cfg.ApplicationId = "123"
	ownedSession(t, fm, "instance")
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{{Id: "new", Status: "ALLOCATED"}}, NextCursor: "second"}, nil)
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "second").Return(nil, errors.New("provider unavailable"))
	_, err := fm.reconcileOnce(context.Background(), map[string]string{"instance": "some-version"}, time.Now().Add(time.Hour))
	require.Error(t, err)
	_, err = fm.storage.GetGameSessionFromStorage(context.Background(), "instance")
	require.NoError(t, err)
	_, err = fm.storage.GetGameSessionFromStorage(context.Background(), "new")
	require.ErrorIs(t, err, storage.ErrSessionNotFound)
}
func TestReconciliationProtectsConcurrentJoinAndReallocation(t *testing.T) {
	for _, action := range []string{"join", "reallocate"} {
		t.Run(action, func(t *testing.T) {
			fm, _, client := sessionFixture(t, 0, 4)
			fm.cfg.ApplicationId = "123"
			ownedSession(t, fm, "instance")
			ctx := context.Background()
			now := time.Now().Add(time.Hour)
			client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").Return(&clients.ApplicationInstanceListResponse{}, nil)
			missing, err := fm.reconcileOnce(ctx, nil, now)
			require.NoError(t, err)
			client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").DoAndReturn(func(context.Context, string, int, string) (*clients.ApplicationInstanceListResponse, error) {
				if action == "join" {
					_, err := fm.Join(ctx, "instance", []string{"user"}, nil)
					require.NoError(t, err)
				} else {
					ownedSession(t, fm, "instance")
				}
				return &clients.ApplicationInstanceListResponse{}, nil
			})
			_, err = fm.reconcileOnce(ctx, missing, now)
			require.NoError(t, err)
			got, err := fm.storage.GetGameSessionFromStorage(ctx, "instance")
			require.NoError(t, err)
			if action == "join" {
				require.Equal(t, 1, got.PlayerCount)
			}
		})
	}
}
func TestReconciliationPreservesCapacityAndScopesAbsence(t *testing.T) {
	fm, _, client := sessionFixture(t, 1, 4)
	fm.cfg.ApplicationId = "123"
	fm.cfg.FleetId = "7"
	require.NoError(t, fm.storage.CreateGameSession(context.Background(), &runtime.InstanceInfo{Id: "scoped", Status: "ALLOCATED", PlayerCount: 1, Metadata: map[string]any{MaxPlayers: 4, "i3d_fleet_id": "7"}}, "123", []string{"user"}))
	ownedSession(t, fm, "outside-fleet")
	require.NoError(t, fm.storage.CreateGameSession(context.Background(), &runtime.InstanceInfo{Id: "other-app", Status: "ALLOCATED"}, "999", nil))
	ctx := context.Background()
	now := time.Now().Add(time.Hour)
	client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123 and fleetId=7", 100, "").Return(&clients.ApplicationInstanceListResponse{NextCursor: "second"}, nil)
	client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123 and fleetId=7", 100, "second").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{{Id: "scoped", Status: "ALLOCATED", PlayerCount: 0, Metadata: map[string]any{"mode": "arena"}}}}, nil)
	missing, err := fm.reconcileOnce(ctx, nil, now)
	require.NoError(t, err)
	require.Empty(t, missing)
	got, err := fm.storage.GetGameSessionFromStorage(ctx, "scoped")
	require.NoError(t, err)
	require.Equal(t, 1, got.PlayerCount)
	capacity, err := getMaxPlayers(got)
	require.NoError(t, err)
	require.Equal(t, 4, capacity)
	for _, id := range []string{"instance", "outside-fleet", "other-app"} {
		_, err = fm.storage.GetGameSessionFromStorage(ctx, id)
		require.NoError(t, err)
	}
}
func TestReconciliationGraceAndCursorLoop(t *testing.T) {
	fm, _, client := sessionFixture(t, 0, 4)
	fm.cfg.ApplicationId = "123"
	fm.cfg.ReconcileGracePeriod = time.Hour
	ownedSession(t, fm, "instance")
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").Return(&clients.ApplicationInstanceListResponse{}, nil).Times(2)
	missing, err := fm.reconcileOnce(context.Background(), nil, time.Now())
	require.NoError(t, err)
	_, err = fm.reconcileOnce(context.Background(), missing, time.Now())
	require.NoError(t, err)
	_, err = fm.storage.GetGameSessionFromStorage(context.Background(), "instance")
	require.NoError(t, err)
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").Return(&clients.ApplicationInstanceListResponse{NextCursor: "loop"}, nil)
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "loop").Return(&clients.ApplicationInstanceListResponse{NextCursor: "loop"}, nil)
	_, err = fm.reconcileOnce(context.Background(), nil, time.Now())
	require.Error(t, err)
}
func TestReconciliationProtectsNewAllocationDuringScan(t *testing.T) {
	fm, _, client := sessionFixture(t, 0, 4)
	fm.cfg.ApplicationId = "123"
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").DoAndReturn(func(context.Context, string, int, string) (*clients.ApplicationInstanceListResponse, error) {
		ownedSession(t, fm, "new")
		return &clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{{Id: "new", Status: "ALLOCATED", PlayerCount: 3}}}, nil
	})
	_, err := fm.reconcileOnce(context.Background(), nil, time.Now())
	require.NoError(t, err)
	got, err := fm.storage.GetGameSessionFromStorage(context.Background(), "new")
	require.NoError(t, err)
	require.Equal(t, 0, got.PlayerCount)
	capacity, err := getMaxPlayers(got)
	require.NoError(t, err)
	require.Equal(t, 4, capacity)
}
func TestReconciliationSnapshotsEveryStoragePage(t *testing.T) {
	fm, _, client := sessionFixture(t, 0, 4)
	fm.cfg.ApplicationId = "123"
	for i := 0; i < 105; i++ {
		ownedSession(t, fm, fmt.Sprintf("owned-%03d", i))
	}
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").Return(&clients.ApplicationInstanceListResponse{}, nil).Times(2)
	ctx := context.Background()
	now := time.Now().Add(time.Hour)
	missing, err := fm.reconcileOnce(ctx, nil, now)
	require.NoError(t, err)
	require.Len(t, missing, 105)
	_, err = fm.reconcileOnce(ctx, missing, now)
	require.NoError(t, err)
	snapshots, err := fm.storage.SnapshotGameSessions(ctx)
	require.NoError(t, err)
	require.Len(t, snapshots, 1) // legacy is retained
}
func TestReconciliationWorkerRecoversAndStops(t *testing.T) {
	fm, nk, client := sessionFixture(t, 0, 4)
	fm.cfg.ApplicationId = "123"
	fm.cfg.ReconcileInterval = 10 * time.Millisecond
	fm.cfg.ReconcileTimeout = time.Second
	fm.ctx, fm.cancel = context.WithCancel(context.Background())
	t.Cleanup(fm.cancel)
	var calls atomic.Int32
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").DoAndReturn(func(context.Context, string, int, string) (*clients.ApplicationInstanceListResponse, error) {
		if calls.Add(1) == 1 {
			return nil, errors.New("startup failure")
		}
		return &clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{{Id: "recovered", Status: "ALLOCATED"}}}, nil
	}).AnyTimes()
	require.NoError(t, fm.Init(nk, &callbackRegistry{callbacks: map[string]runtime.FmCreateCallbackFn{}}))
	require.Eventually(t, func() bool {
		_, err := fm.storage.GetGameSessionFromStorage(context.Background(), "recovered")
		return err == nil
	}, time.Second, time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fm.shutdown(ctx)
	count := calls.Load()
	require.GreaterOrEqual(t, count, int32(2))
	select {
	case <-time.After(3 * fm.cfg.ReconcileInterval):
	}
	require.Equal(t, count, calls.Load())
}
func TestConfiguredFleetScopesAllocation(t *testing.T) {
	fm, client, cache, _, _ := createFixture(t)
	fm.cfg.FleetId = "7"
	client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), "(dcLocationId=1) and fleetId=7").Return(&runtime.InstanceInfo{Id: "one"}, nil)
	cache.EXPECT().CreateGameSession(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, instance *runtime.InstanceInfo, _ string, _ []string) error {
		require.Equal(t, "7", instance.Metadata["i3d_fleet_id"])
		return nil
	})
	ch := make(chan createResult, 1)
	_, err := fm.Create(context.Background(), 4, nil, nil, map[string]any{clients.I3dFilters: "dcLocationId=1"}, resultCallback(ch))
	require.NoError(t, err)
	require.NoError(t, awaitCreate(t, ch).err)
}
func TestProviderResponseCannotDeleteReallocatedSession(t *testing.T) {
	for _, operation := range []string{"get", "delete"} {
		t.Run(operation, func(t *testing.T) {
			fm, _, client := sessionFixture(t, 0, 4)
			fm.cfg.ApplicationId = "123"
			ownedSession(t, fm, "instance")
			if operation == "get" {
				client.EXPECT().GetApplicationInstance(gomock.Any(), "instance").DoAndReturn(func(context.Context, string) (*runtime.InstanceInfo, error) {
					ownedSession(t, fm, "instance")
					return &runtime.InstanceInfo{Id: "instance", Status: "ONLINE"}, nil
				}).Times(4)
				_, _ = fm.Get(context.Background(), "instance")
			} else {
				client.EXPECT().RestartApplicationInstance(gomock.Any(), "instance").DoAndReturn(func(context.Context, string) error { ownedSession(t, fm, "instance"); return nil })
				_ = fm.Delete(context.Background(), "instance")
			}
			got, err := fm.storage.GetGameSessionFromStorage(context.Background(), "instance")
			require.NoError(t, err)
			require.Equal(t, "ALLOCATED", got.Status)
		})
	}
}

type allocationDuringSnapshot struct {
	storage.FleetManagerStorage
	create func()
}

func (s allocationDuringSnapshot) SnapshotGameSessions(ctx context.Context) ([]*api.StorageObject, error) {
	s.create()
	return s.FleetManagerStorage.SnapshotGameSessions(ctx)
}
func TestReconciliationExcludesAllocationCreatedDuringStorageScan(t *testing.T) {
	fm, _, client := sessionFixture(t, 0, 4)
	fm.cfg.ApplicationId = "123"
	base := fm.storage
	fm.storage = allocationDuringSnapshot{FleetManagerStorage: base, create: func() {
		require.NoError(t, base.CreateGameSession(context.Background(), &runtime.InstanceInfo{Id: "new", Status: "ALLOCATED", CreateTime: time.Unix(200, 0), Metadata: map[string]any{MaxPlayers: 4, "map": "new-game"}}, "123", nil))
	}}
	// The provider scan can briefly still return the previous generation.
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{{Id: "new", Status: "ALLOCATED", CreateTime: time.Unix(100, 0), Metadata: map[string]any{"map": "old-game"}}}}, nil)
	_, err := fm.reconcileOnce(context.Background(), nil, time.Now())
	require.NoError(t, err)
	got, err := base.GetGameSessionFromStorage(context.Background(), "new")
	require.NoError(t, err)
	require.Equal(t, "new-game", got.Metadata["map"])
	capacity, err := getMaxPlayers(got)
	require.NoError(t, err)
	require.Equal(t, 4, capacity)
}

type recentlyWrittenSnapshot struct{ storage.FleetManagerStorage }

func (s recentlyWrittenSnapshot) SnapshotGameSessions(ctx context.Context) ([]*api.StorageObject, error) {
	objects, err := s.FleetManagerStorage.SnapshotGameSessions(ctx)
	for _, object := range objects {
		object.UpdateTime = timestamppb.Now()
	}
	return objects, err
}
func TestReconciliationExcludesRecordUpdatedDuringStorageScan(t *testing.T) {
	fm, _, client := sessionFixture(t, 0, 4)
	fm.cfg.ApplicationId = "123"
	require.NoError(t, fm.storage.CreateGameSession(context.Background(), &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", CreateTime: time.Unix(200, 0), Metadata: map[string]any{MaxPlayers: 4, "map": "new-game"}}, "123", nil))
	// The allocation began earlier, but storage reports an update during this pass.
	fm.storage = recentlyWrittenSnapshot{fm.storage}
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{{Id: "instance", Status: "ALLOCATED", CreateTime: time.Unix(100, 0), Metadata: map[string]any{"map": "old-game"}}}}, nil)
	_, err := fm.reconcileOnce(context.Background(), nil, time.Now())
	require.NoError(t, err)
	got, err := fm.storage.GetGameSessionFromStorage(context.Background(), "instance")
	require.NoError(t, err)
	require.Equal(t, "new-game", got.Metadata["map"])
}

func TestReconciliationProtectsWritesWithBoundedClockSkew(t *testing.T) {
	for _, kind := range []string{"allocation", "storage-update"} {
		t.Run(kind, func(t *testing.T) {
			fm, nk, client := sessionFixture(t, 0, 4)
			ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_ENV, map[string]string{"I3D_APPLICATION_ID": "123", "I3D_ACCESS_TOKEN": "test-only", "I3D_RECONCILE_CLOCK_SKEW": "10s"})
			cfg, configErr := config.NewConfigFromRuntime(ctx)
			require.Nil(t, configErr)
			fm.cfg = cfg
			base := fm.storage
			if kind == "allocation" {
				fm.storage = allocationDuringSnapshot{FleetManagerStorage: base, create: func() {
					require.NoError(t, base.CreateGameSession(ctx, &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", CreateTime: time.Unix(200, 0), Metadata: map[string]any{MaxPlayers: 4, "map": "new-game"}}, "123", nil))
				}}
			} else {
				require.NoError(t, base.CreateGameSession(ctx, &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", CreateTime: time.Unix(200, 0), Metadata: map[string]any{MaxPlayers: 4, "map": "new-game"}}, "123", nil))
				objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{{Collection: storage.StorageI3dInstancesCollection, Key: "instance"}})
				require.NoError(t, err)
				require.Len(t, objects, 1)
				var value map[string]any
				require.NoError(t, json.Unmarshal([]byte(objects[0].Value), &value))
				value["_i3d"].(map[string]any)["allocated_at"] = time.Now().Add(-time.Hour).UTC().Format(time.RFC3339Nano)
				encoded, err := json.Marshal(value)
				require.NoError(t, err)
				_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{{Collection: storage.StorageI3dInstancesCollection, Key: "instance", Version: objects[0].Version, Value: string(encoded)}})
				require.NoError(t, err)
				fm.storage = recentlyWrittenSnapshot{base}
			}
			client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{{Id: "instance", Status: "ALLOCATED", CreateTime: time.Unix(100, 0), Metadata: map[string]any{"map": "old-game"}}}}, nil)
			_, err := fm.reconcileOnce(ctx, nil, time.Now().Add(7*time.Second))
			require.NoError(t, err)
			got, err := base.GetGameSessionFromStorage(ctx, "instance")
			require.NoError(t, err)
			require.Equal(t, "new-game", got.Metadata["map"])
			capacity, err := getMaxPlayers(got)
			require.NoError(t, err)
			require.Equal(t, 4, capacity)
		})
	}
}

func TestProgrammaticClockSkewCannotBypassMinimum(t *testing.T) {
	fm, _, client := sessionFixture(t, 0, 4)
	fm.cfg.ApplicationId = "123"
	fm.cfg.ReconcileClockSkew = 500 * time.Millisecond
	require.NoError(t, fm.storage.CreateGameSession(context.Background(), &runtime.InstanceInfo{
		Id: "instance", Status: "ALLOCATED", CreateTime: time.Unix(200, 0), Metadata: map[string]any{MaxPlayers: 4, "map": "new-game"},
	}, "123", nil))
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").Return(&clients.ApplicationInstanceListResponse{
		Instances: []*runtime.InstanceInfo{{Id: "instance", Status: "ALLOCATED", CreateTime: time.Unix(100, 0), Metadata: map[string]any{"map": "old-game"}}},
	}, nil)
	_, err := fm.reconcileOnce(context.Background(), nil, time.Now().Add(750*time.Millisecond))
	require.NoError(t, err)
	got, err := fm.storage.GetGameSessionFromStorage(context.Background(), "instance")
	require.NoError(t, err)
	require.Equal(t, "new-game", got.Metadata["map"])
}
