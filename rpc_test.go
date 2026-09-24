package fleetmanager

import (
	"context"
	"errors"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestLifecycleRPCRejectsPlayerBeforeDecoding(t *testing.T) {
	for _, rpc := range []string{"update", "delete"} {
		for _, payload := range []string{"not json", `{"id":"instance","player_count":1}`} {
			t.Run(rpc+"/"+payload, func(t *testing.T) {
				fm, _, _, _, _ := createFixture(t)
				ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_USER_ID, "ordinary-player")
				handler := fm.UpdateInstanceInfo
				if rpc == "delete" {
					handler = fm.DeleteInstanceInfo
				}
				_, err := handler(ctx, fm.logger, nil, nil, payload)
				var runtimeErr *runtime.Error
				require.ErrorAs(t, err, &runtimeErr)
				require.Equal(t, PERMISSION_DENIED, runtimeErr.Code)
			})
		}
	}
}
func TestLifecycleRPCRejectsInvalidPayloadWithoutSideEffects(t *testing.T) {
	for _, test := range []struct{ rpc, payload string }{
		{"update", "null"}, {"update", "["}, {"update", `{"id":"","player_count":0}`},
		{"update", `{"id":"instance","player_count":-1}`},
		{"update", `{"id":"instance","playerCount":1}`},
		{"update", `{"id":"instance","metadata":{"i3d_max_players":999}}`},
		{"update", `{"id":"instance","metadata":{"overwriteApplicationId":"other"}}`},
		{"update", `{"id":"instance","metadata":{"i3dFilters":"other"}}`},
		{"delete", "null"}, {"delete", "["}, {"delete", `{"id":" "}`}, {"delete", `{"id":"instance"} {}`},
	} {
		t.Run(test.rpc+"/"+test.payload, func(t *testing.T) {
			fm, _, _, _, _ := createFixture(t)
			handler := fm.UpdateInstanceInfo
			if test.rpc == "delete" {
				handler = fm.DeleteInstanceInfo
			}
			_, err := handler(context.Background(), fm.logger, nil, nil, test.payload)
			var runtimeErr *runtime.Error
			require.ErrorAs(t, err, &runtimeErr)
			require.Equal(t, INVALID_ARGUMENT, runtimeErr.Code)
		})
	}
}
func TestServerUpdateRPCUsesDocumentedPayload(t *testing.T) {
	fm, client, cache, _, _ := createFixture(t)
	metadata := map[string]any{"map": "arena"}
	client.EXPECT().UpdateApplicationInstance(gomock.Any(), "instance", metadata).Return(&runtime.InstanceInfo{Id: "instance", Metadata: metadata}, nil)
	cache.EXPECT().MutateGameSession(gomock.Any(), "instance", true, gomock.Any()).DoAndReturn(func(_ context.Context, _ string, _ bool, fn func(*runtime.InstanceInfo, map[string]bool) error) (*runtime.InstanceInfo, error) {
		instance := &runtime.InstanceInfo{Id: "instance"}
		err := fn(instance, map[string]bool{})
		require.NoError(t, err)
		require.Equal(t, 3, instance.PlayerCount)
		require.Equal(t, metadata, instance.Metadata)
		return instance, err
	})
	_, err := fm.UpdateInstanceInfo(context.Background(), fm.logger, nil, nil, `{"id":"instance","player_count":3,"metadata":{"map":"arena"}}`)
	require.NoError(t, err)
}
func TestServerDeleteRPCUsesDocumentedPayload(t *testing.T) {
	fm, client, cache, _, _ := createFixture(t)
	client.EXPECT().RestartApplicationInstance(gomock.Any(), "instance").Return(nil)
	cache.EXPECT().GetGameSessionSnapshot(gomock.Any(), "instance").Return(nil, nil)
	cache.EXPECT().ReconcileGameSession(gomock.Any(), nil, nil, "", "").Return(nil)
	_, err := fm.DeleteInstanceInfo(context.Background(), fm.logger, nil, nil, `{"id":"instance"}`)
	require.NoError(t, err)
}
func TestLifecycleRPCControlsProviderErrors(t *testing.T) {
	for _, rpc := range []string{"update", "delete"} {
		t.Run(rpc, func(t *testing.T) {
			fm, client, cache, _, _ := createFixture(t)
			handler := fm.UpdateInstanceInfo
			if rpc == "delete" {
				handler = fm.DeleteInstanceInfo
				cache.EXPECT().GetGameSessionSnapshot(gomock.Any(), "instance").Return(nil, nil)
				client.EXPECT().RestartApplicationInstance(gomock.Any(), "instance").Return(errors.New("provider failed"))
			} else {
				client.EXPECT().UpdateApplicationInstance(gomock.Any(), "instance", gomock.Any()).Return(nil, errors.New("provider failed"))
			}
			_, err := handler(context.Background(), fm.logger, nil, nil, `{"id":"instance"}`)
			var runtimeErr *runtime.Error
			require.ErrorAs(t, err, &runtimeErr)
			require.Equal(t, INTERNAL, runtimeErr.Code)
		})
	}
}
