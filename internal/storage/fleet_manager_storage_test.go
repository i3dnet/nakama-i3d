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
)

type storageNakama struct {
	runtime.NakamaModule
	read  func(context.Context, []*runtime.StorageRead) ([]*api.StorageObject, error)
	index func(context.Context, string, string, string, int, []string, string) (*api.StorageObjects, string, error)
}

func (n storageNakama) StorageRead(ctx context.Context, r []*runtime.StorageRead) ([]*api.StorageObject, error) {
	return n.read(ctx, r)
}
func (n storageNakama) StorageIndexList(ctx context.Context, caller, index, query string, limit int, order []string, cursor string) (*api.StorageObjects, string, error) {
	return n.index(ctx, caller, index, query, limit, order, cursor)
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
