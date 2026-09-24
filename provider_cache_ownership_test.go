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

func TestProviderOnlyAllocationIsNotImported(t *testing.T) {
	for _, method := range []string{"list", "get", "reconcile"} {
		t.Run(method, func(t *testing.T) {
			fm, nk, client := sessionFixture(t, 0, 4)
			fm.cfg.ApplicationId = "123"
			provider := &runtime.InstanceInfo{Id: "provider-only", Status: "ALLOCATED", Metadata: map[string]any{"mode": "arena", MaxPlayers: "999"}}
			nk.WriteError = errors.New("provider-only reads must not attempt a storage write")
			switch method {
			case "list":
				client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 5, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{provider}, NextCursor: "next"}, nil)
				got, next, err := fm.List(context.Background(), "", 5, "")
				require.NoError(t, err)
				require.Len(t, got, 1)
				require.Equal(t, "next", next)
				require.NotContains(t, got[0].Metadata, MaxPlayers)
			case "get":
				client.EXPECT().GetApplicationInstance(gomock.Any(), "provider-only").Return(provider, nil)
				got, err := fm.Get(context.Background(), "provider-only")
				require.NoError(t, err)
				require.Equal(t, "arena", got.Metadata["mode"])
				require.NotContains(t, got.Metadata, MaxPlayers)
			case "reconcile":
				client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 100, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{provider}}, nil)
				_, err := fm.reconcileOnce(context.Background(), nil, time.Now().Add(time.Hour))
				require.NoError(t, err)
			}
			_, err := fm.storage.GetGameSessionFromStorage(context.Background(), "provider-only")
			require.ErrorIs(t, err, storage.ErrSessionNotFound)
			_, err = fm.Join(context.Background(), "provider-only", []string{"player"}, nil)
			require.ErrorIs(t, err, storage.ErrSessionNotFound)
		})
	}
}

func TestProviderListRefreshesOwnedRecordsWithoutCachingUnknownOnes(t *testing.T) {
	fm, _, client := sessionFixture(t, 1, 4)
	fm.cfg.ApplicationId = "123"
	client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 5, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{
		{Id: "provider-only", Status: "ALLOCATED", Metadata: map[string]any{"mode": "remote"}},
		{Id: "instance", Status: "ALLOCATED", Metadata: map[string]any{"mode": "updated"}},
	}}, nil)
	got, _, err := fm.List(context.Background(), "", 5, "")
	require.NoError(t, err)
	require.Len(t, got, 2)
	require.Equal(t, "remote", got[0].Metadata["mode"])
	require.Equal(t, "updated", got[1].Metadata["mode"])
	require.Equal(t, 1, got[1].PlayerCount)
	capacity, err := getMaxPlayers(got[1])
	require.NoError(t, err)
	require.Equal(t, 4, capacity)
	_, err = fm.storage.GetGameSessionFromStorage(context.Background(), "provider-only")
	require.ErrorIs(t, err, storage.ErrSessionNotFound)
}

func TestProviderGenerationChangeInvalidatesCachedAdmission(t *testing.T) {
	for _, method := range []string{"list", "get", "reconcile"} {
		t.Run(method, func(t *testing.T) {
			fm, _, client := sessionFixture(t, 0, 4)
			fm.cfg.ApplicationId = "123"
			first := time.Unix(100, 0)
			require.NoError(t, fm.storage.CreateGameSession(context.Background(), &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", CreateTime: first, Metadata: map[string]any{MaxPlayers: 4}}, "123", nil))
			provider := &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", CreateTime: first.Add(time.Hour), Metadata: map[string]any{"mode": "replacement"}}
			now := time.Now().Add(time.Hour)
			switch method {
			case "list":
				client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 5, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{provider}}, nil)
				_, _, err := fm.List(context.Background(), "", 5, "")
				require.NoError(t, err)
				snapshot, err := fm.storage.GetGameSessionSnapshot(context.Background(), "instance")
				require.NoError(t, err)
				require.NotNil(t, snapshot)
				scoped, _, err := storage.SnapshotInScope(snapshot, "123", "")
				require.NoError(t, err)
				require.True(t, scoped, "invalidated generations must remain owned for cleanup")
				client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 5, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{provider}}, nil)
				_, _, err = fm.List(context.Background(), "", 5, "")
				require.NoError(t, err)
				repeated, err := fm.storage.GetGameSessionSnapshot(context.Background(), "instance")
				require.NoError(t, err)
				require.Equal(t, snapshot.Version, repeated.Version, "repeated discovery must not keep postponing cleanup")

			case "get":
				client.EXPECT().GetApplicationInstance(gomock.Any(), "instance").Return(provider, nil)
				_, err := fm.Get(context.Background(), "instance")
				require.NoError(t, err)
			case "reconcile":
				client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 100, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{provider}}, nil)
				_, err := fm.reconcileOnce(context.Background(), nil, now)
				require.NoError(t, err)
			}
			_, err := fm.Join(context.Background(), "instance", []string{"player"}, nil)
			require.Error(t, err, "an unrelated provider generation must not inherit local admission capacity")
			if method == "list" {
				client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 100, "").Return(&clients.ApplicationInstanceListResponse{}, nil).Times(2)
				missing, err := fm.reconcileOnce(context.Background(), nil, now)
				require.NoError(t, err)
				_, err = fm.reconcileOnce(context.Background(), missing, now)
				require.NoError(t, err)
			}
			_, err = fm.storage.GetGameSessionFromStorage(context.Background(), "instance")
			require.ErrorIs(t, err, storage.ErrSessionNotFound)
		})
	}
}

func TestOlderProviderGenerationDoesNotReplaceLocalAdmission(t *testing.T) {
	for _, method := range []string{"list", "get", "reconcile"} {
		t.Run(method, func(t *testing.T) {
			fm, _, client := sessionFixture(t, 0, 4)
			fm.cfg.ApplicationId = "123"
			created := time.Unix(1000, 0)
			require.NoError(t, fm.storage.CreateGameSession(context.Background(), &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", CreateTime: created, Metadata: map[string]any{MaxPlayers: 4}}, "123", nil))
			provider := &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", CreateTime: time.Unix(100, 0)}
			switch method {
			case "list":
				client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 5, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{provider}}, nil)
				_, _, err := fm.List(context.Background(), "", 5, "")
				require.NoError(t, err)
			case "get":
				client.EXPECT().GetApplicationInstance(gomock.Any(), "instance").Return(provider, nil)
				_, err := fm.Get(context.Background(), "instance")
				require.NoError(t, err)
			case "reconcile":
				client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 100, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{provider}}, nil)
				_, err := fm.reconcileOnce(context.Background(), nil, time.Now().Add(time.Hour))
				require.NoError(t, err)
			}
			joined, err := fm.Join(context.Background(), "instance", []string{"player"}, nil)
			require.NoError(t, err)
			require.Len(t, joined.SessionInfo, 1)
			require.WithinDuration(t, created, joined.InstanceInfo.CreateTime, 0)
		})
	}
}

func TestDelayedUpdateCannotRestoreInvalidatedAdmission(t *testing.T) {
	fm, _, client := sessionFixture(t, 0, 4)
	fm.cfg.ApplicationId = "123"
	first := time.Unix(100, 0)
	require.NoError(t, fm.storage.CreateGameSession(context.Background(), &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", CreateTime: first, Metadata: map[string]any{MaxPlayers: 4}}, "123", nil))
	ready, resume := make(chan struct{}), make(chan struct{})
	client.EXPECT().UpdateApplicationInstance(gomock.Any(), "instance", 0, gomock.Any()).DoAndReturn(func(context.Context, string, int, map[string]any) (*runtime.InstanceInfo, error) {
		close(ready)
		<-resume
		return &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", CreateTime: first}, nil
	})
	finished := make(chan error, 1)
	go func() { finished <- fm.Update(context.Background(), "instance", 0, nil) }()
	<-ready
	client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 5, "").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{{Id: "instance", Status: "ALLOCATED", CreateTime: first.Add(time.Hour)}}}, nil)
	_, _, listErr := fm.List(context.Background(), "", 5, "")
	before, readErr := fm.storage.GetGameSessionSnapshot(context.Background(), "instance")
	close(resume)
	updateErr := <-finished
	require.NoError(t, listErr)
	require.NoError(t, readErr)
	require.Error(t, updateErr)
	after, err := fm.storage.GetGameSessionSnapshot(context.Background(), "instance")
	require.NoError(t, err)
	require.Equal(t, before.Version, after.Version)
	_, err = fm.Join(context.Background(), "instance", []string{"player"}, nil)
	require.Error(t, err)
}

func TestUpdateDoesNotImportProviderOnlyAllocation(t *testing.T) {
	fm, _, _ := sessionFixture(t, 0, 4)
	require.ErrorIs(t, fm.Update(context.Background(), "provider-only", 1, nil), storage.ErrSessionNotFound)
	_, err := fm.storage.GetGameSessionFromStorage(context.Background(), "provider-only")
	require.ErrorIs(t, err, storage.ErrSessionNotFound)
}
