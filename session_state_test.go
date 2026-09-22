package fleetmanager

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/config"
	"github.com/i3dnet/nakama-i3d/internal/storage"
	"github.com/i3dnet/nakama-i3d/internal/tests"
	"github.com/i3dnet/nakama-i3d/internal/tests/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type storageInitializer struct{ runtime.Initializer }

func (storageInitializer) RegisterStorageIndex(string, string, string, []string, []string, int, bool) error {
	return nil
}
func sessionFixture(t *testing.T, count, capacity int) (*I3dFleetManager, *tests.MemoryNakama, *mock.MockApplicationInstance) {
	t.Helper()
	nk := tests.NewMemoryNakama()
	logger := tests.NewMockLogger()
	cache, err := storage.NewFleetManagerStorageService(nk, storageInitializer{}, logger)
	require.NoError(t, err)
	instance := &runtime.InstanceInfo{Id: "instance", PlayerCount: count, Status: "ALLOCATED", Metadata: map[string]any{MaxPlayers: capacity}}
	data, err := json.Marshal(instance)
	require.NoError(t, err)
	_, err = nk.StorageWrite(context.Background(), []*runtime.StorageWrite{{Collection: storage.StorageI3dInstancesCollection, Key: "instance", Value: string(data)}})
	require.NoError(t, err)
	client := mock.NewMockApplicationInstance(gomock.NewController(t))
	fm := &I3dFleetManager{ctx: context.Background(), cfg: &config.Config{}, storage: cache, logger: logger, client: client}
	return fm, nk, client
}
func TestConcurrentJoinCannotOverbookLastSlot(t *testing.T) {
	fm, nk, _ := sessionFixture(t, 0, 1)
	// Separate manager objects represent two Nakama nodes sharing storage.
	other := &I3dFleetManager{storage: fm.storage, logger: fm.logger}
	ready := make(chan struct{})
	var reads atomic.Int32
	nk.AfterRead = func() {
		n := reads.Add(1)
		if n == 2 {
			close(ready)
		}
		if n <= 2 {
			<-ready
		}
	}
	results := make(chan *runtime.JoinInfo, 2)
	failures := make(chan error, 2)
	var wg sync.WaitGroup
	for i, manager := range []*I3dFleetManager{fm, other} {
		wg.Add(1)
		go func(i int, manager *I3dFleetManager) {
			defer wg.Done()
			result, err := manager.Join(context.Background(), "instance", []string{[]string{"a", "b"}[i]}, nil)
			results <- result
			failures <- err
		}(i, manager)
	}
	wg.Wait()
	close(results)
	close(failures)
	for err := range failures {
		require.NoError(t, err)
	}
	admitted := 0
	for result := range results {
		admitted += len(result.SessionInfo)
	}
	require.Equal(t, 1, admitted)
	stored, err := fm.storage.GetGameSessionFromStorage(context.Background(), "instance")
	require.NoError(t, err)
	require.Equal(t, 1, stored.PlayerCount)
}
func TestRepeatedJoinDoesNotCountUsersTwice(t *testing.T) {
	fm, _, _ := sessionFixture(t, 0, 3)
	first, err := fm.Join(context.Background(), "instance", []string{"a", "a"}, nil)
	require.NoError(t, err)
	require.Len(t, first.SessionInfo, 1)
	require.Equal(t, 1, first.InstanceInfo.PlayerCount)
	again, err := fm.Join(context.Background(), "instance", []string{"a", "b"}, nil)
	require.NoError(t, err)
	require.Len(t, again.SessionInfo, 2)
	require.Equal(t, 2, again.InstanceInfo.PlayerCount)
}
func TestProviderRefreshPreservesCapacityAndAdmissions(t *testing.T) {
	fm, _, client := sessionFixture(t, 1, 3)
	client.EXPECT().GetApplicationInstance(gomock.Any(), "instance").Return(&runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", PlayerCount: 0, Metadata: map[string]any{"map": "new"}}, nil)
	refreshed, err := fm.Get(context.Background(), "instance")
	require.NoError(t, err)
	capacity, err := getMaxPlayers(refreshed)
	require.NoError(t, err)
	require.Equal(t, 3, capacity)
	require.Equal(t, 1, refreshed.PlayerCount)
	require.Equal(t, "new", refreshed.Metadata["map"])
}

func TestNewAllocationResetsReusedInstanceAdmissions(t *testing.T) {
	fm, _, _ := sessionFixture(t, 0, 3)
	first := &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", PlayerCount: 1, Metadata: map[string]any{MaxPlayers: 3}}
	require.NoError(t, fm.storage.CreateGameSession(context.Background(), first, "123", []string{"a"}))
	joined, err := fm.Join(context.Background(), "instance", []string{"a"}, nil)
	require.NoError(t, err)
	require.Equal(t, 1, joined.InstanceInfo.PlayerCount)
	next := &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", Metadata: map[string]any{MaxPlayers: 4}}
	require.NoError(t, fm.storage.CreateGameSession(context.Background(), next, "123", nil))
	joined, err = fm.Join(context.Background(), "instance", []string{"a"}, nil)
	require.NoError(t, err)
	require.Equal(t, 1, joined.InstanceInfo.PlayerCount)
	capacity, err := getMaxPlayers(joined.InstanceInfo)
	require.NoError(t, err)
	require.Equal(t, 4, capacity)
}
func TestJoinRetriesOnlyVersionConflictsWithinBound(t *testing.T) {
	for _, test := range []struct {
		name    string
		failure error
		writes  int
		want    error
	}{
		{"version conflicts", runtime.ErrStorageRejectedVersion, 6, runtime.ErrStorageWriteExhaustedRetries},
		{"other failure", errors.New("storage unavailable"), 1, nil},
	} {
		t.Run(test.name, func(t *testing.T) {
			fm, nk, _ := sessionFixture(t, 0, 3)
			before := nk.Writes
			nk.WriteError = test.failure
			result, err := fm.Join(context.Background(), "instance", []string{"a"}, nil)
			require.Error(t, err)
			require.Nil(t, result)
			require.Equal(t, test.writes, nk.Writes-before)
			if test.want != nil {
				require.ErrorIs(t, err, test.want)
			} else {
				require.ErrorIs(t, err, test.failure)
			}
		})
	}
}
func TestCanceledJoinDoesNotWrite(t *testing.T) {
	fm, nk, _ := sessionFixture(t, 0, 3)
	before := nk.Writes
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	result, err := fm.Join(ctx, "instance", []string{"a"}, nil)
	require.ErrorIs(t, err, context.Canceled)
	require.Nil(t, result)
	require.Equal(t, before, nk.Writes)
}
func TestJoinReturnsOnlyAdmittedUsersWhenCapacityIsPartial(t *testing.T) {
	fm, _, _ := sessionFixture(t, 2, 3)
	result, err := fm.Join(context.Background(), "instance", []string{"a", "b"}, nil)
	require.NoError(t, err)
	require.Equal(t, 3, result.InstanceInfo.PlayerCount)
	require.Len(t, result.SessionInfo, 1)
	require.Equal(t, "a", result.SessionInfo[0].UserId)
	require.Empty(t, result.SessionInfo[0].SessionId)
}
func TestAuthoritativePlayerUpdatePreservesCapacity(t *testing.T) {
	fm, _, client := sessionFixture(t, 1, 3)
	client.EXPECT().UpdateApplicationInstance(gomock.Any(), "instance", map[string]any{"map": "arena"}).Return(&runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", Metadata: map[string]any{"map": "arena"}}, nil)
	require.NoError(t, fm.Update(context.Background(), "instance", 2, map[string]any{"map": "arena"}))
	stored, err := fm.storage.GetGameSessionFromStorage(context.Background(), "instance")
	require.NoError(t, err)
	capacity, err := getMaxPlayers(stored)
	require.NoError(t, err)
	require.Equal(t, 3, capacity)
	require.Equal(t, 2, stored.PlayerCount)
}
func TestRefreshCannotOverwriteConcurrentJoin(t *testing.T) {
	fm, nk, client := sessionFixture(t, 0, 3)
	client.EXPECT().GetApplicationInstance(gomock.Any(), "instance").Return(&runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED", Metadata: map[string]any{"map": "new"}}, nil).MinTimes(1).MaxTimes(2)
	ready := make(chan struct{})
	var reads atomic.Int32
	nk.AfterRead = func() {
		n := reads.Add(1)
		if n == 2 {
			close(ready)
		}
		if n <= 2 {
			<-ready
		}
	}
	failures := make(chan error, 2)
	go func() { _, err := fm.Join(context.Background(), "instance", []string{"a"}, nil); failures <- err }()
	go func() { _, err := fm.Get(context.Background(), "instance"); failures <- err }()
	require.NoError(t, <-failures)
	require.NoError(t, <-failures)
	stored, err := fm.storage.GetGameSessionFromStorage(context.Background(), "instance")
	require.NoError(t, err)
	require.Equal(t, 1, stored.PlayerCount)
	require.Equal(t, "new", stored.Metadata["map"])
}
