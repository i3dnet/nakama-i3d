package fleetmanager

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/internal/clients"
	"github.com/i3dnet/nakama-i3d/internal/storage"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestReconciliationIncludesApplicationOverrides(t *testing.T) {
	for _, mode := range []string{"missing", "new-generation", "refresh"} {
		t.Run(mode, func(t *testing.T) {
			fm, _, client := sessionFixture(t, 0, 4)
			fm.cfg.ApplicationId, fm.cfg.FleetId = "123", "7"
			fm.callbackHandler = &callbackRegistry{callbacks: map[string]runtime.FmCreateCallbackFn{}}
			ctx := context.Background()
			created := time.Unix(100, 0)
			client.EXPECT().AllocateApplicationInstance(gomock.Any(), map[string]any{clients.ApplicationId: "456"}, "fleetId=7").
				Return(&runtime.InstanceInfo{Id: "override", Status: "ALLOCATED", CreateTime: created}, nil)
			done := make(chan createResult, 1)
			_, err := fm.Create(ctx, 4, []string{"player"}, nil, map[string]any{clients.ApplicationId: "456"}, resultCallback(done))
			require.NoError(t, err)
			require.NoError(t, awaitCreate(t, done).err)
			fm.operations.Wait()
			// An application owned by another fleet must not expand this manager's scan.
			require.NoError(t, fm.storage.CreateGameSession(ctx, &runtime.InstanceInfo{
				Id: "foreign", Status: "ALLOCATED", Metadata: map[string]any{MaxPlayers: 4, "i3d_fleet_id": "8"},
			}, "789", nil))
			scans := 1
			if mode == "missing" {
				scans = 2
			}
			for i := 0; i < scans; i++ {
				// Cursor values are scoped to each application, not shared between scans.
				client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123 and fleetId=7", 100, "").
					Return(&clients.ApplicationInstanceListResponse{NextCursor: "second"}, nil)
				client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123 and fleetId=7", 100, "second").
					Return(&clients.ApplicationInstanceListResponse{}, nil)
				client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=456 and fleetId=7", 100, "").
					Return(&clients.ApplicationInstanceListResponse{NextCursor: "second"}, nil)
				response := &clients.ApplicationInstanceListResponse{}
				if mode != "missing" {
					incoming := &runtime.InstanceInfo{Id: "override", Status: "ALLOCATED", CreateTime: created, Metadata: map[string]any{"mode": "refreshed"}}
					if mode == "new-generation" {
						incoming.CreateTime = created.Add(time.Hour)
					}
					response.Instances = []*runtime.InstanceInfo{incoming}
				}
				client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=456 and fleetId=7", 100, "second").
					Return(response, nil)
			}
			now := time.Now().Add(time.Hour)
			missing, err := fm.reconcileOnce(ctx, nil, now)
			require.NoError(t, err)
			if mode == "missing" {
				require.Contains(t, missing, "override")
				_, err = fm.storage.GetGameSessionFromStorage(ctx, "override")
				require.NoError(t, err, "one complete scan must not remove the session")
				_, err = fm.reconcileOnce(ctx, missing, now)
				require.NoError(t, err)
			}
			got, err := fm.storage.GetGameSessionFromStorage(ctx, "override")
			if mode == "refresh" {
				require.NoError(t, err)
				require.Equal(t, 1, got.PlayerCount)
				capacity, err := getMaxPlayers(got)
				require.NoError(t, err)
				require.Equal(t, 4, capacity)
				require.Equal(t, "refreshed", got.Metadata["mode"])
				snapshot, err := fm.storage.GetGameSessionSnapshot(ctx, "override")
				require.NoError(t, err)
				scoped, _, err := storage.SnapshotInScope(snapshot, "456", "7")
				require.NoError(t, err)
				require.True(t, scoped, "refresh must preserve the overridden application's ownership")
			} else {
				require.ErrorIs(t, err, storage.ErrSessionNotFound)
			}
			for _, id := range []string{"instance", "foreign"} {
				_, err := fm.storage.GetGameSessionFromStorage(ctx, id)
				require.NoError(t, err, "legacy and other-fleet records must survive")
			}
		})
	}
}

func TestFailedOverrideScanCannotMutateAnyApplication(t *testing.T) {
	fm, _, client := sessionFixture(t, 0, 4)
	fm.cfg.ApplicationId, fm.cfg.FleetId = "123", "7"
	ctx := context.Background()
	for id, app := range map[string]string{"default": "123", "override": "456"} {
		require.NoError(t, fm.storage.CreateGameSession(ctx, &runtime.InstanceInfo{
			Id: id, Status: "ALLOCATED", CreateTime: time.Unix(100, 0),
			Metadata: map[string]any{MaxPlayers: 4, "i3d_fleet_id": "7"},
		}, app, nil))
	}
	before, err := fm.storage.SnapshotGameSessions(ctx)
	require.NoError(t, err)
	previous := map[string]string{}
	for _, obj := range before {
		previous[obj.Key] = obj.Version
	}
	client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123 and fleetId=7", 100, "").
		Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{
			{Id: "default", Status: "ALLOCATED", CreateTime: time.Unix(200, 0)},
		}}, nil)
	client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=456 and fleetId=7", 100, "").
		Return(&clients.ApplicationInstanceListResponse{NextCursor: "second"}, nil)
	failure := errors.New("override scan failed")
	client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=456 and fleetId=7", 100, "second").
		Return(nil, failure)
	_, err = fm.reconcileOnce(ctx, previous, time.Now().Add(time.Hour))
	require.ErrorIs(t, err, failure)
	after, err := fm.storage.SnapshotGameSessions(ctx)
	require.NoError(t, err)
	require.Equal(t, before, after, "no refresh or removal is safe until every application scan completes")
}
