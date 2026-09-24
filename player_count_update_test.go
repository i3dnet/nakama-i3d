package fleetmanager

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/config"
	"github.com/i3dnet/nakama-i3d/internal/clients"
	"github.com/i3dnet/nakama-i3d/internal/openapi"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestPlayerCountUpdateRejectsOverCapacityBeforeProvider(t *testing.T) {
	for _, rpc := range []bool{false, true} {
		name := "Update"
		if rpc {
			name = "RPC"
		}
		t.Run(name, func(t *testing.T) {
			fm, _, client := sessionFixture(t, 1, 3)
			calls := 0
			client.EXPECT().UpdateApplicationInstance(gomock.Any(), "instance", gomock.Any(), gomock.Any()).DoAndReturn(func(context.Context, string, int, map[string]any) (*runtime.InstanceInfo, error) {
				calls++
				return &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED"}, nil
			}).AnyTimes()
			before, err := fm.storage.GetGameSessionSnapshot(context.Background(), "instance")
			require.NoError(t, err)
			if rpc {
				_, err = fm.UpdateInstanceInfo(context.Background(), fm.logger, nil, nil, `{"id":"instance","player_count":4}`)
			} else {
				err = fm.Update(context.Background(), "instance", 4, nil)
			}
			require.ErrorIs(t, err, ErrInvalidInput)
			require.Zero(t, calls)
			after, err := fm.storage.GetGameSessionSnapshot(context.Background(), "instance")
			require.NoError(t, err)
			require.Equal(t, before, after)
			joined, err := fm.Join(context.Background(), "instance", []string{"new-player"}, nil)
			require.NoError(t, err)
			require.Equal(t, 2, joined.InstanceInfo.PlayerCount)
		})
	}
}

func TestPlayerCountUpdateRechecksCapacityAfterConcurrentChange(t *testing.T) {
	for _, duringRetry := range []bool{false, true} {
		name := "provider request"
		if duringRetry {
			name = "storage version retry"
		}
		t.Run(name, func(t *testing.T) {
			fm, nk, client := sessionFixture(t, 1, 4)
			shrinkCapacity := func() {
				_, err := fm.storage.MutateGameSession(context.Background(), "instance", false, func(instance *runtime.InstanceInfo, _ map[string]bool) error {
					instance.Metadata[MaxPlayers] = 2
					return nil
				})
				require.NoError(t, err)
			}
			client.EXPECT().UpdateApplicationInstance(gomock.Any(), "instance", gomock.Any(), gomock.Any()).DoAndReturn(func(context.Context, string, int, map[string]any) (*runtime.InstanceInfo, error) {
				if duringRetry {
					nk.AfterRead = func() {
						nk.AfterRead = nil
						shrinkCapacity()
					}
				} else {
					shrinkCapacity()
				}
				return &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED"}, nil
			})
			require.ErrorIs(t, fm.Update(context.Background(), "instance", 3, nil), ErrInvalidInput)
			stored, err := fm.storage.GetGameSessionFromStorage(context.Background(), "instance")
			require.NoError(t, err)
			capacity, err := getMaxPlayers(stored)
			require.NoError(t, err)
			require.Equal(t, 2, capacity)
			require.Equal(t, 1, stored.PlayerCount)
		})
	}
}

func TestPlayerCountUpdateDoesNotOverwriteNewProviderGeneration(t *testing.T) {
	fm, _, client := sessionFixture(t, 1, 4)
	first := time.Unix(100, 0)
	latest := first.Add(time.Hour)
	client.EXPECT().UpdateApplicationInstance(gomock.Any(), "instance", gomock.Any(), gomock.Any()).DoAndReturn(func(context.Context, string, int, map[string]any) (*runtime.InstanceInfo, error) {
		err := fm.storage.CreateGameSession(context.Background(), &runtime.InstanceInfo{Id: "instance", CreateTime: latest, Status: "ALLOCATED", PlayerCount: 1, Metadata: map[string]any{MaxPlayers: 4}}, "123", []string{"player"})
		require.NoError(t, err)
		return &runtime.InstanceInfo{Id: "instance", CreateTime: first, Status: "ALLOCATED"}, nil
	})
	require.Error(t, fm.Update(context.Background(), "instance", 2, nil))
	stored, err := fm.storage.GetGameSessionFromStorage(context.Background(), "instance")
	require.NoError(t, err)
	require.WithinDuration(t, latest, stored.CreateTime, 0)
	require.Equal(t, 1, stored.PlayerCount)
	capacity, err := getMaxPlayers(stored)
	require.NoError(t, err)
	require.Equal(t, 4, capacity)
}

func TestPlayerCountUpdateSendsReportedCountToProvider(t *testing.T) {
	for _, count := range []int{0, 3} {
		t.Run(strconv.Itoa(count), func(t *testing.T) {
			fm, _, _ := sessionFixture(t, 1, 3)
			updates := make(chan openapi.ApplicationInstance, 1)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				instance := openapi.ApplicationInstance{Id: "instance", ApplicationId: "123", Status: 5, NumPlayers: 1,
					IpAddress:  []openapi.ApplicationInstanceIP{{IpAddress: "203.0.113.10", IpVersion: 4}},
					Properties: []openapi.ApplicationInstanceProperty{{PropertyKey: "port", PropertyValue: "7777"}}}
				if r.Method == http.MethodPut {
					if err := json.NewDecoder(r.Body).Decode(&instance); err != nil {
						t.Error(err)
						w.WriteHeader(http.StatusBadRequest)
						return
					}
					updates <- instance
				}
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode([]openapi.ApplicationInstance{instance})
			}))
			defer server.Close()
			cfg := &config.Config{OneApi: config.OneApi{BaseUrl: server.URL, ApplicationId: "123", Token: "test-only"}, Retry: config.Retry{Attempts: 1}}
			fm.client = clients.NewOneApiClient(cfg, clients.NewAuthentication(cfg), fm.logger)
			require.NoError(t, fm.Update(context.Background(), "instance", count, map[string]any{"map": "arena"}))
			update := <-updates
			require.EqualValues(t, count, update.NumPlayers)
			require.Equal(t, []openapi.Metadata{{Key: "map", Value: "arena"}}, update.Metadata)
			stored, err := fm.storage.GetGameSessionFromStorage(context.Background(), "instance")
			require.NoError(t, err)
			require.Equal(t, count, stored.PlayerCount)
		})
	}
}

func TestPlayerCountUpdateRejectsProviderRangeOverflow(t *testing.T) {
	count := int(math.MaxInt32) + 1
	fm, _, client := sessionFixture(t, 0, count)
	calls := 0
	client.EXPECT().UpdateApplicationInstance(gomock.Any(), "instance", count, gomock.Any()).DoAndReturn(func(context.Context, string, int, map[string]any) (*runtime.InstanceInfo, error) {
		calls++
		return &runtime.InstanceInfo{Id: "instance", Status: "ALLOCATED"}, nil
	}).AnyTimes()
	require.ErrorIs(t, fm.Update(context.Background(), "instance", count, nil), ErrInvalidInput)
	require.Zero(t, calls)
}
