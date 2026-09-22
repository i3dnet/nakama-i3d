package fleetmanager

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/config"
	"github.com/i3dnet/nakama-i3d/internal/tests"
	"github.com/i3dnet/nakama-i3d/internal/tests/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type callbackRegistry struct {
	mu          sync.Mutex
	callbacks   map[string]runtime.FmCreateCallbackFn
	next        int
	invocations int
}

func (r *callbackRegistry) GenerateCallbackId() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.next++
	return strconv.Itoa(r.next)
}
func (r *callbackRegistry) SetCallback(id string, fn runtime.FmCreateCallbackFn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.callbacks[id] = fn
}
func (r *callbackRegistry) InvokeCallback(id string, status runtime.FmCreateStatus, instance *runtime.InstanceInfo, sessions []*runtime.SessionInfo, metadata map[string]any, err error) {
	r.mu.Lock()
	r.invocations++
	fn := r.callbacks[id]
	delete(r.callbacks, id)
	r.mu.Unlock()
	if fn != nil {
		fn(status, instance, sessions, metadata, err)
	}
}

type createResult struct {
	status   runtime.FmCreateStatus
	instance *runtime.InstanceInfo
	sessions []*runtime.SessionInfo
	metadata map[string]any
	err      error
}

func resultCallback(ch chan<- createResult) runtime.FmCreateCallbackFn {
	return func(s runtime.FmCreateStatus, i *runtime.InstanceInfo, users []*runtime.SessionInfo, m map[string]any, err error) {
		ch <- createResult{s, i, users, m, err}
	}
}
func awaitCreate(t *testing.T, ch <-chan createResult) createResult {
	t.Helper()
	select {
	case result := <-ch:
		return result
	case <-time.After(time.Second):
		t.Fatal("no terminal callback")
		return createResult{}
	}
}
func createFixture(t *testing.T) (*I3dFleetManager, *mock.MockApplicationInstance, *tests.MockFleetManagerStorage, *callbackRegistry, context.CancelFunc) {
	t.Helper()
	ctrl := gomock.NewController(t)
	client := mock.NewMockApplicationInstance(ctrl)
	cache := tests.NewMockFleetManagerStorage(ctrl)
	lifetime, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	registry := &callbackRegistry{callbacks: map[string]runtime.FmCreateCallbackFn{}}
	fm := &I3dFleetManager{ctx: lifetime, cfg: &config.Config{}, client: client, storage: cache, logger: tests.NewMockLogger(), callbackHandler: registry}
	return fm, client, cache, registry, cancel
}

func TestCreateSurvivesHookCancellationAndPersistsBeforeCallback(t *testing.T) {
	fm, client, cache, registry, _ := createFixture(t)
	started, release, stored := make(chan struct{}), make(chan struct{}), make(chan struct{})
	type key struct{}
	hook, cancel := context.WithCancel(context.WithValue(context.Background(), key{}, "trace"))
	defer cancel()
	instance := &runtime.InstanceInfo{Id: "instance", Metadata: map[string]any{}}
	client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, _ map[string]any, _ string) (*runtime.InstanceInfo, error) {
		close(started)
		<-release
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		if ctx.Value(key{}) != "trace" {
			return nil, errors.New("lost context values")
		}
		return instance, nil
	})
	cache.EXPECT().UpdateStorageGameSession(gomock.Any(), gomock.Any()).DoAndReturn(func(context.Context, []*runtime.InstanceInfo) error { close(stored); return nil })
	result := make(chan createResult, 2)
	_, err := fm.Create(hook, 4, []string{"player"}, nil, map[string]any{"map": "arena"}, func(s runtime.FmCreateStatus, i *runtime.InstanceInfo, users []*runtime.SessionInfo, m map[string]any, e error) {
		select {
		case <-stored:
		default:
			e = errors.New("callback before storage")
		}
		result <- createResult{s, i, users, m, e}
	})
	require.NoError(t, err)
	<-started
	cancel()
	close(release)
	got := awaitCreate(t, result)
	require.NoError(t, got.err)
	require.Equal(t, runtime.CreateSuccess, got.status)
	require.Equal(t, "player", got.sessions[0].UserId)
	require.Equal(t, 4, got.instance.Metadata[MaxPlayers])
	registry.mu.Lock()
	require.Empty(t, registry.callbacks)
	registry.mu.Unlock()
}

func TestCreateRejectsInvalidRequestsBeforeRegisteringCallback(t *testing.T) {
	for _, test := range []struct {
		name     string
		max      int
		users    []string
		canceled bool
	}{
		{"zero capacity", 0, nil, false}, {"too many users", 1, []string{"a", "b"}, false}, {"empty user", 1, []string{""}, false}, {"duplicate user", 2, []string{"a", "a"}, false}, {"canceled", 1, nil, true},
	} {
		t.Run(test.name, func(t *testing.T) {
			fm, _, _, registry, _ := createFixture(t)
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			if test.canceled {
				cancel()
			}
			_, err := fm.Create(ctx, test.max, test.users, nil, nil, resultCallback(make(chan createResult, 1)))
			require.Error(t, err)
			registry.mu.Lock()
			require.Empty(t, registry.callbacks)
			registry.mu.Unlock()
		})
	}
}

func TestCreateStorageFailureIsTerminalError(t *testing.T) {
	fm, client, cache, _, _ := createFixture(t)
	client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(&runtime.InstanceInfo{Id: "instance", Metadata: map[string]any{}}, nil)
	cache.EXPECT().UpdateStorageGameSession(gomock.Any(), gomock.Any()).Return(errors.New("storage unavailable"))
	result := make(chan createResult, 2)
	_, err := fm.Create(context.Background(), 2, nil, nil, nil, resultCallback(result))
	require.NoError(t, err)
	got := awaitCreate(t, result)
	require.Error(t, got.err)
	require.Equal(t, runtime.CreateError, got.status)
	require.Nil(t, got.instance)
	require.Nil(t, got.sessions)
	require.Nil(t, got.metadata)
}

func TestCreateNoUsersReturnsNilSessions(t *testing.T) {
	fm, client, cache, _, _ := createFixture(t)
	client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(&runtime.InstanceInfo{Id: "instance", Metadata: map[string]any{}}, nil)
	cache.EXPECT().UpdateStorageGameSession(gomock.Any(), gomock.Any()).Return(nil)
	result := make(chan createResult, 1)
	_, err := fm.Create(context.Background(), 2, nil, nil, nil, resultCallback(result))
	require.NoError(t, err)
	got := awaitCreate(t, result)
	require.NoError(t, got.err)
	require.Nil(t, got.sessions)
}

func TestCreateShutdownConsumesCallback(t *testing.T) {
	fm, client, _, registry, shutdown := createFixture(t)
	started := make(chan struct{})
	client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, _ map[string]any, _ string) (*runtime.InstanceInfo, error) {
		close(started)
		<-ctx.Done()
		return nil, ctx.Err()
	})
	result := make(chan createResult, 2)
	_, err := fm.Create(context.Background(), 2, nil, nil, nil, resultCallback(result))
	require.NoError(t, err)
	<-started
	shutdown()
	got := awaitCreate(t, result)
	require.ErrorIs(t, got.err, context.Canceled)
	require.Equal(t, runtime.CreateError, got.status)
	registry.mu.Lock()
	require.Empty(t, registry.callbacks)
	registry.mu.Unlock()
}

func TestCreateTimeoutIgnoresLateCompletion(t *testing.T) {
	fm, client, _, registry, _ := createFixture(t)
	fm.cfg.AllocationTimeout = 10 * time.Millisecond
	started, release := make(chan struct{}), make(chan struct{})
	client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(context.Context, map[string]any, string) (*runtime.InstanceInfo, error) {
		close(started)
		<-release
		return &runtime.InstanceInfo{Id: "late"}, nil
	})
	result := make(chan createResult, 2)
	_, err := fm.Create(context.Background(), 2, nil, nil, nil, resultCallback(result))
	require.NoError(t, err)
	<-started
	got := awaitCreate(t, result)
	require.ErrorIs(t, got.err, context.DeadlineExceeded)
	close(release)
	fm.operations.Wait()
	require.Empty(t, result)
	registry.mu.Lock()
	require.Empty(t, registry.callbacks)
	require.Equal(t, 1, registry.invocations)
	registry.mu.Unlock()
}

func TestCreateNilCallbackDoesNotRegister(t *testing.T) {
	fm, client, cache, registry, _ := createFixture(t)
	client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(&runtime.InstanceInfo{Id: "instance"}, nil)
	cache.EXPECT().UpdateStorageGameSession(gomock.Any(), gomock.Any()).Return(nil)
	_, err := fm.Create(context.Background(), 2, nil, nil, nil, nil)
	require.NoError(t, err)
	fm.operations.Wait()
	registry.mu.Lock()
	require.Empty(t, registry.callbacks)
	require.Zero(t, registry.invocations)
	registry.mu.Unlock()
}

func TestShutdownRejectsFurtherCreate(t *testing.T) {
	fm, _, _, registry, cancel := createFixture(t)
	fm.cancel = cancel
	fm.shutdown(context.Background())
	_, err := fm.Create(context.Background(), 2, nil, nil, nil, resultCallback(make(chan createResult, 1)))
	require.Error(t, err)
	registry.mu.Lock()
	require.Empty(t, registry.callbacks)
	registry.mu.Unlock()
}
