package fleetmanager

import (
	"context"
	"errors"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/internal/clients"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestFilterBuilderProducesRawQuotedExpression(t *testing.T) {
	require.Equal(t, `regionName="EU \"West\"" and fleetId=123`, NewFilterBuilder().Add(RegionName, `EU "West"`).Add(FleetId, "123").Query())
}
func TestFilterBuilderInitializesNilMetadata(t *testing.T) {
	got := NewFilterBuilder().Add(FleetId, "123").AddFiltersToMetaData(nil)
	require.Equal(t, "fleetId=123", GetFilters(got))
}

func TestDefaultListFiltersStatusLocallyAndRetainsProviderCursor(t *testing.T) {
	fm, client, cache, _, _ := createFixture(t)
	fm.cfg.ApplicationId = "123"
	allocated := &runtime.InstanceInfo{Id: "allocated", Status: "ALLOCATED"}
	online := &runtime.InstanceInfo{Id: "online", Status: "ONLINE"}
	client.EXPECT().ListApplicationInstances(gomock.Any(), "applicationId=123", 7, "previous").Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{allocated, online}, NextCursor: "next"}, nil)
	cache.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{allocated}).Return(nil)
	got, cursor, err := fm.List(context.Background(), "", 7, "previous")
	require.NoError(t, err)
	require.Equal(t, []*runtime.InstanceInfo{allocated}, got)
	require.Equal(t, "next", cursor)
}
func TestDefaultListPropagatesCacheFailure(t *testing.T) {
	fm, client, cache, _, _ := createFixture(t)
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&clients.ApplicationInstanceListResponse{Instances: []*runtime.InstanceInfo{{Id: "instance", Status: "ALLOCATED"}}}, nil)
	cache.EXPECT().UpdateStorageGameSession(gomock.Any(), gomock.Any()).Return(errors.New("cache failed"))
	_, _, err := fm.List(context.Background(), "", 7, "")
	require.Error(t, err)
}
