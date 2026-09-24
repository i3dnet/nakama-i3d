package fleetmanager

import (
	"context"
	"errors"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/internal/tests"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"strings"
	"testing"
	"time"
)

func TestCreatePersistenceHasFreshBudgetAfterProviderCompletion(t *testing.T) {
	fm, client, cache, _, _ := createFixture(t)
	fm.cfg.AllocationTimeout = 25 * time.Millisecond
	type traceKey struct{}
	hook := context.WithValue(context.Background(), traceKey{}, "request-trace")
	var providerCtx context.Context
	client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, _ map[string]any, _ string) (*runtime.InstanceInfo, error) {
		providerCtx = ctx
		return &runtime.InstanceInfo{Id: "instance"}, nil
	})
	cache.EXPECT().CreateGameSession(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, _ *runtime.InstanceInfo, _ string, _ []string) error {
		<-providerCtx.Done()
		if ctx.Value(traceKey{}) != "request-trace" {
			return errors.New("lost request context values")
		}
		return ctx.Err()
	})
	done := make(chan createResult, 1)
	_, err := fm.Create(hook, 2, nil, nil, nil, resultCallback(done))
	require.NoError(t, err)
	result := awaitCreate(t, done)
	fm.operations.Wait()
	require.NoError(t, result.err)
	require.Equal(t, runtime.CreateSuccess, result.status)
}

func TestCreateKnownAllocationErrorRestartsOnce(t *testing.T) {
	fm, client, _, registry, _ := createFixture(t)
	failure := errors.New("allocation response incomplete")
	client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(&runtime.InstanceInfo{Id: "allocated"}, failure)
	cleanupContext := make(chan context.Context, 1)
	client.EXPECT().RestartApplicationInstance(gomock.Any(), "allocated").DoAndReturn(func(ctx context.Context, _ string) error {
		cleanupContext <- ctx
		if ctx.Err() != nil {
			return ctx.Err()
		}
		return nil
	})
	done := make(chan createResult, 2)
	_, err := fm.Create(context.Background(), 2, nil, nil, nil, resultCallback(done))
	require.NoError(t, err)
	require.ErrorIs(t, awaitCreate(t, done).err, failure)
	fm.operations.Wait()
	select {
	case ctx := <-cleanupContext:
		_, bounded := ctx.Deadline()
		require.True(t, bounded)
	default:
		t.Error("known allocation was not restarted")
	}
	require.Empty(t, done)
	registry.mu.Lock()
	defer registry.mu.Unlock()
	require.Equal(t, 1, registry.invocations)
}

func TestCreateReportsDefaultAllocationTimeoutOnce(t *testing.T) {
	fm, client, cache, _, _ := createFixture(t)
	client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(&runtime.InstanceInfo{Id: "instance"}, nil).Times(2)
	cache.EXPECT().CreateGameSession(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(2)
	for i := 0; i < 2; i++ {
		done := make(chan createResult, 1)
		_, err := fm.Create(context.Background(), 2, nil, nil, nil, resultCallback(done))
		require.NoError(t, err)
		require.NoError(t, awaitCreate(t, done).err)
	}
	fm.operations.Wait()
	warnings := 0
	for _, entry := range fm.logger.(*tests.MockLogger).Entries() {
		if entry.Level == "WARN" && strings.Contains(entry.Message, "I3D_ALLOCATION_TIMEOUT") && strings.Contains(entry.Message, "120s") {
			warnings++
		}
	}
	require.Equal(t, 1, warnings, "defaulting must be visible without warning on every match")
}

func TestPublishedAllocationResultWinsOverExpiredContext(t *testing.T) {
	for _, failure := range []error{nil, errors.New("original failure")} {
		result := &allocationResult{done: make(chan struct{})}
		instance := &runtime.InstanceInfo{Id: "instance"}
		require.True(t, result.complete(allocationOutcome{instance, failure}))
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		got := result.wait(ctx)
		require.Same(t, instance, got.instance)
		require.Equal(t, failure, got.err)
	}
}

func TestCreateStorageTimeoutReclaimsLateCommitOnce(t *testing.T) {
	fm, client, cache, registry, _ := createFixture(t)
	fm.cfg.AllocationFinalizeTimeout = 10 * time.Millisecond
	release := make(chan struct{})
	client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(&runtime.InstanceInfo{Id: "late-storage"}, nil)
	cache.EXPECT().CreateGameSession(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(context.Context, *runtime.InstanceInfo, string, []string) error {
		<-release
		return nil // A storage implementation may commit despite cancellation.
	})
	client.EXPECT().RestartApplicationInstance(gomock.Any(), "late-storage").DoAndReturn(func(ctx context.Context, _ string) error {
		require.NoError(t, ctx.Err())
		deadline, ok := ctx.Deadline()
		require.True(t, ok)
		require.LessOrEqual(t, time.Until(deadline), fm.cfg.AllocationFinalizeTimeout)
		return nil
	})
	done := make(chan createResult, 2)
	_, err := fm.Create(context.Background(), 2, nil, nil, nil, resultCallback(done))
	require.NoError(t, err)
	require.ErrorIs(t, awaitCreate(t, done).err, context.DeadlineExceeded)
	close(release)
	fm.operations.Wait()
	require.Empty(t, done)
	registry.mu.Lock()
	defer registry.mu.Unlock()
	require.Equal(t, 1, registry.invocations)
}

func TestCreateCleanupFailurePreservesAllocationError(t *testing.T) {
	fm, client, _, _, _ := createFixture(t)
	failure := errors.New("provider readiness failed")
	client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(&runtime.InstanceInfo{Id: "instance"}, failure)
	client.EXPECT().RestartApplicationInstance(gomock.Any(), "instance").Return(errors.New("restart failed"))
	done := make(chan createResult, 1)
	_, err := fm.Create(context.Background(), 2, nil, nil, nil, resultCallback(done))
	require.NoError(t, err)
	require.ErrorIs(t, awaitCreate(t, done).err, failure)
	fm.operations.Wait()
	found := false
	for _, entry := range fm.logger.(*tests.MockLogger).Entries() {
		if entry.Level == "ERROR" && strings.Contains(entry.Message, "failed to reclaim allocation") {
			found = true
		}
	}
	require.True(t, found, "cleanup failure must be visible to operators")
}

func TestCreateUnknownAllocationDoesNotGuessCleanupID(t *testing.T) {
	for _, instance := range []*runtime.InstanceInfo{nil, {}} {
		fm, client, _, _, _ := createFixture(t)
		client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(instance, errors.New("ambiguous allocation"))
		done := make(chan createResult, 1)
		_, err := fm.Create(context.Background(), 2, nil, nil, nil, resultCallback(done))
		require.NoError(t, err)
		require.Error(t, awaitCreate(t, done).err)
		fm.operations.Wait()
	}
}
