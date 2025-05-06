package fleetmanager

import (
	"context"
	"errors"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/suite"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/config"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/clients"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/storage"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/tests"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/tests/mock"
	"go.uber.org/mock/gomock"
	"sync"
	"testing"
	"time"
)

type FleetManagerSuite struct {
	suite.Suite
	cfg             *config.Config
	logger          *tests.MockLogger
	fleetManager    *I3dFleetManager
	callBackHandler *tests.CallbackHandlerMock
	ctx             context.Context
}

func (suite *FleetManagerSuite) SetupTest() {
	suite.cfg = &config.Config{
		App: config.App{
			Name:    "test",
			Version: "1.0.1",
		},
		OneApi: config.OneApi{
			BaseUrl:       "http://localhost:8080",
			ApplicationId: "1235",
			Token:         "1234567890",
			UseBearerAuth: false,
		},
	}

}

func (suite *FleetManagerSuite) newTestFleetManager(client clients.ApplicationInstance, storage storage.FleetManagerStorage) *I3dFleetManager {
	suite.logger = tests.NewMockLogger()
	return &I3dFleetManager{
		cfg:             suite.cfg,
		client:          client,
		storage:         storage,
		logger:          suite.logger,
		ctx:             context.TODO(),
		callbackHandler: tests.NewCallbackHandlerMock(suite.T()),
	}
}

func (suite *FleetManagerSuite) TestGet_shouldReturnInstance() {
	// assert
	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0)
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().GetApplicationInstance(gomock.Any(), gomock.Any()).Return(expected, nil).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(nil).Times(1)
	storageService.EXPECT().DeleteStorageGameSession(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	// act
	result, err := suite.newTestFleetManager(client, storageService).Get(suite.ctx, expected.Id)

	// assert
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(expected.Id, result.Id)

}

func (suite *FleetManagerSuite) TestGet_GivenInstanceNotStatusAllocated_shouldReturnInstanceAndDeleteFromStorage() {

	for _, status := range clients.ApplicationInstanceStatus {
		// assert
		if status == clients.ApplicationInstanceStatus[5] {
			continue
		}
		expected := InstanceInfoMock(status, 0)
		ctrl := gomock.NewController(suite.T())
		client := mock.NewMockApplicationInstance(ctrl)
		storageService := tests.NewMockFleetManagerStorage(ctrl)
		defer ctrl.Finish()

		client.EXPECT().GetApplicationInstance(gomock.Any(), gomock.Any()).Return(expected, nil).Times(1)
		storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(nil).Times(0)
		storageService.EXPECT().DeleteStorageGameSession(gomock.Any(), gomock.Any()).Return(nil).Times(1)

		// act
		result, err := suite.newTestFleetManager(client, storageService).Get(suite.ctx, expected.Id)

		// assert
		suite.NoError(err)
		suite.NotNil(result)
		suite.Equal(expected.Id, result.Id)
	}
}

func (suite *FleetManagerSuite) TestGet_GivenAnApiError_ShouldReturnError() {

	// assert
	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0)
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().GetApplicationInstance(gomock.Any(), gomock.Any()).Return(nil, errors.New("Api Error")).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(nil).Times(0)
	storageService.EXPECT().DeleteStorageGameSession(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	// act
	result, err := suite.newTestFleetManager(client, storageService).Get(suite.ctx, expected.Id)

	// assert
	suite.NotNil(err)
	suite.Nil(result)
}

func (suite *FleetManagerSuite) TestGet_GivenAnStorageUpdateError_ShouldReturnError() {

	// assert
	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0)
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().GetApplicationInstance(gomock.Any(), gomock.Any()).Return(expected, nil).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(errors.New("storage update error")).Times(1)
	storageService.EXPECT().DeleteStorageGameSession(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	// act
	result, err := suite.newTestFleetManager(client, storageService).Get(suite.ctx, expected.Id)

	// assert
	suite.NotNil(err)
	suite.Nil(result)
}

func (suite *FleetManagerSuite) TestGet_GivenAnStorageDeleteError_ShouldReturnError() {

	// assert
	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[4], 0)
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().GetApplicationInstance(gomock.Any(), gomock.Any()).Return(expected, nil).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(nil).Times(0)
	storageService.EXPECT().DeleteStorageGameSession(gomock.Any(), gomock.Any()).Return(errors.New("storage delete error")).Times(1)

	// act
	result, err := suite.newTestFleetManager(client, storageService).Get(suite.ctx, expected.Id)

	// assert
	suite.NotNil(err)
	suite.Nil(result)
}

func (suite *FleetManagerSuite) TestList_shouldReturnInstances() {
	// assert
	expected := []*runtime.InstanceInfo{
		InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0),
		InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0),
	}
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	response := &clients.ApplicationInstanceListResponse{
		Instances:  expected,
		NextCursor: "",
	}
	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(response, nil).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), expected).Return(nil).Times(1)
	storageService.EXPECT().DeleteStorageGameSession(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	// act
	result, _, err := suite.newTestFleetManager(client, storageService).List(suite.ctx, "", 0, "")

	// assert
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(len(expected), len(result))
}

func (suite *FleetManagerSuite) TestList_GivenAnApiError_ShouldReturnError() {
	// assert
	expected := []*runtime.InstanceInfo{
		InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0),
		InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0),
	}
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("Api Error")).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), expected).Return(nil).Times(0)
	storageService.EXPECT().DeleteStorageGameSession(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	// act
	result, _, err := suite.newTestFleetManager(client, storageService).List(suite.ctx, "", 0, "")

	// assert
	suite.NotNil(err)
	suite.Nil(result)
}

func (suite *FleetManagerSuite) TestList_GivenQuery_ShouldReturnFromStorage() {
	// assert
	const query = "test"
	expected := []*runtime.InstanceInfo{
		InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0),
		InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0),
	}
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil).Times(0)
	storageService.EXPECT().ListGameSessionsFromStorage(gomock.Any(), query, gomock.Any(), gomock.Any(), gomock.Any()).Return(expected, nil).Times(1)

	// act
	result, _, err := suite.newTestFleetManager(client, storageService).List(suite.ctx, "test", 0, "")

	// assert
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(len(expected), len(result))
}

func (suite *FleetManagerSuite) TestList_GivenAnStorageError_ShouldReturnError() {
	// assert
	const query = "test"
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil).Times(0)
	storageService.EXPECT().ListGameSessionsFromStorage(gomock.Any(), query, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("storage error")).Times(1)

	// act
	result, _, err := suite.newTestFleetManager(client, storageService).List(suite.ctx, query, 0, "")

	// assert
	suite.NotNil(err)
	suite.Nil(result)
}

func (suite *FleetManagerSuite) TestCreate_GivenCallback_ShouldCallCallback() {
	// assert
	const maxPlayerCount = 10
	const userId = "1"
	userIds := []string{userId}
	metaData := map[string]any{}
	metaData["key"] = "value"

	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0)
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().AllocateApplicationInstance(gomock.Any(), metaData, gomock.Any()).Return(expected, nil).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(nil).Times(1)

	var wg sync.WaitGroup
	wg.Add(1)

	var callback runtime.FmCreateCallbackFn = func(status runtime.FmCreateStatus, instanceInfo *runtime.InstanceInfo, sessionInfo []*runtime.SessionInfo, metadata map[string]any, err error) {
		suite.Equal(runtime.CreateSuccess, status)
		suite.Equal(expected.Id, instanceInfo.Id)
		suite.Equal(metaData, metadata)
		suite.Equal(userId, sessionInfo[0].UserId)
		suite.Nil(err)
		wg.Done()
	}

	var latency []runtime.FleetUserLatencies
	// act
	err := suite.newTestFleetManager(client, storageService).Create(suite.ctx, maxPlayerCount, userIds, latency, metaData, callback)

	// assert
	suite.NoError(err)

	// 👇 Wait for callback to complete
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		suite.T().Fatal("Callback was not invoked in time")
	}
}

func (suite *FleetManagerSuite) TestCreate_GivenAllocationFails_ShouldGiveError() {
	// assert
	const maxPlayerCount = 10
	const userId = "1"
	userIds := []string{userId}
	metaData := map[string]any{}
	metaData["key"] = "value"

	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0)
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().AllocateApplicationInstance(gomock.Any(), metaData, gomock.Any()).Return(nil, errors.New("allocation failed")).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(nil).Times(0)

	var wg sync.WaitGroup
	wg.Add(1)

	var callback runtime.FmCreateCallbackFn = func(status runtime.FmCreateStatus, instanceInfo *runtime.InstanceInfo, sessionInfo []*runtime.SessionInfo, metadata map[string]any, err error) {
		suite.Equal(runtime.CreateError, status)
		suite.Error(err)
		wg.Done()
	}

	var latency []runtime.FleetUserLatencies
	// act
	err := suite.newTestFleetManager(client, storageService).Create(suite.ctx, maxPlayerCount, userIds, latency, metaData, callback)

	// assert
	suite.NoError(err)

	// 👇 Wait for callback to complete
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		suite.T().Fatal("Callback was not invoked in time")
	}
}

func (suite *FleetManagerSuite) TestCreate_GivenWithOutCallback_ShouldSucceed() {
	// assert
	const maxPlayerCount = 10
	const userId = "1"
	userIds := []string{userId}
	metaData := map[string]any{}
	metaData["key"] = "value"

	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0)
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().AllocateApplicationInstance(gomock.Any(), metaData, gomock.Any()).Return(expected, nil).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(nil).Times(1)

	var wg sync.WaitGroup
	wg.Add(1)

	var latency []runtime.FleetUserLatencies
	// act
	err := suite.newTestFleetManager(client, storageService).Create(suite.ctx, maxPlayerCount, userIds, latency, metaData, nil)

	// assert
	suite.NoError(err)

	// 👇 Wait for callback to complete
	done := make(chan struct{})
	go func() {
		time.Sleep(500 * time.Millisecond)
		wg.Done()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(2 * time.Second):
		suite.T().Fatal("Callback was not invoked in time")
	}
}

func (suite *FleetManagerSuite) TestCreate_GivenUpdatingStorageFails_ShouldSucceed() {
	// assert
	const maxPlayerCount = 10
	const userId = "1"
	const expectedErrorMessage = "error writing to Nakama storage after starting session"
	userIds := []string{userId}
	metaData := map[string]any{}
	metaData["key"] = "value"

	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0)
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().AllocateApplicationInstance(gomock.Any(), metaData, gomock.Any()).Return(expected, nil).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(errors.New("failed to save to storage")).Times(1)

	var wg sync.WaitGroup
	wg.Add(1)

	var callback runtime.FmCreateCallbackFn = func(status runtime.FmCreateStatus, instanceInfo *runtime.InstanceInfo, sessionInfo []*runtime.SessionInfo, metadata map[string]any, err error) {
		suite.Equal(runtime.CreateSuccess, status)
		suite.Equal(expected.Id, instanceInfo.Id)
		suite.Equal(metaData, metadata)
		suite.Equal(userId, sessionInfo[0].UserId)
		suite.Nil(err)
		wg.Done()
	}

	var latency []runtime.FleetUserLatencies
	// act
	err := suite.newTestFleetManager(client, storageService).Create(suite.ctx, maxPlayerCount, userIds, latency, metaData, callback)

	// assert
	suite.NoError(err)

	// 👇 Wait for callback to complete
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		suite.True(suite.logger.HasError(expectedErrorMessage))
	case <-time.After(2 * time.Second):
		suite.T().Fatal("Callback was not invoked in time")
	}
}

func (suite *FleetManagerSuite) TestJoin_GivenAnExistingInstanceIdInCache_ShouldAddPlayersAndReturnInstance() {
	// assert
	const playerCount = 5
	const userId = "1"
	userIds := []string{userId}

	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], playerCount)
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	storageService.EXPECT().GetGameSessionFromStorage(gomock.Any(), expected.Id).Return(expected, nil).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(nil).Times(1)

	// act
	result, err := suite.newTestFleetManager(client, storageService).Join(suite.ctx, expected.Id, userIds, nil)

	// assert
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(expected.Id, result.InstanceInfo.Id)
	suite.Equal(playerCount+len(userIds), result.InstanceInfo.PlayerCount)
	suite.Equal(len(userIds), len(result.SessionInfo))
	suite.Equal(userId, result.SessionInfo[0].UserId)
}

func (suite *FleetManagerSuite) TestJoin_GivenAnNotExistingInstanceIdInCache_ShouldReturnError() {
	// assert
	const playerCount = 10
	const userId = "1"
	userIds := []string{userId}

	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], playerCount)
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	storageService.EXPECT().GetGameSessionFromStorage(gomock.Any(), expected.Id).Return(nil, errors.New("Not Found")).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(nil).Times(0)

	// act
	result, err := suite.newTestFleetManager(client, storageService).Join(suite.ctx, expected.Id, userIds, nil)

	// assert
	suite.Error(err)
	suite.Nil(result)
}

func (suite *FleetManagerSuite) TestJoin_GivenUpdatingStorageFailed_ShouldReturnError() {
	// assert
	const playerCount = 10
	const userId = "1"
	userIds := []string{userId}

	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], playerCount)
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	storageService.EXPECT().GetGameSessionFromStorage(gomock.Any(), expected.Id).Return(expected, nil).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(errors.New("failed to save to storage")).Times(1)

	// act
	result, err := suite.newTestFleetManager(client, storageService).Join(suite.ctx, expected.Id, userIds, nil)

	// assert
	suite.Error(err)
	suite.Nil(result)
	suite.True(suite.logger.HasError("failed to update storage"))
}

func (suite *FleetManagerSuite) TestJoin_GivenAllSlotsOccupied_ShouldReturnInstanceWithoutSessions() {
	// assert
	const playerCount = 10
	const userId = "1"
	userIds := []string{userId}

	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], playerCount)
	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	storageService.EXPECT().GetGameSessionFromStorage(gomock.Any(), expected.Id).Return(expected, nil).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(nil).Times(1)

	// act
	result, err := suite.newTestFleetManager(client, storageService).Join(suite.ctx, expected.Id, userIds, nil)

	// assert
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(expected.Id, result.InstanceInfo.Id)
	suite.Equal(playerCount, result.InstanceInfo.PlayerCount)
	suite.Equal(0, len(result.SessionInfo))
}

func (suite *FleetManagerSuite) TestJoin_GivenNoMaxPlayersSet_ShouldReturnError() {
	// assert
	const playerCount = 10
	const userId = "1"
	userIds := []string{userId}

	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], playerCount)
	expected.Metadata = map[string]any{"key": "value"}

	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	storageService.EXPECT().GetGameSessionFromStorage(gomock.Any(), expected.Id).Return(expected, nil).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), []*runtime.InstanceInfo{expected}).Return(nil).Times(0)

	// act
	result, err := suite.newTestFleetManager(client, storageService).Join(suite.ctx, expected.Id, userIds, nil)

	// assert
	suite.Error(err)
	suite.Nil(result)
	suite.True(suite.logger.HasError("failed to get max players"))
}

func InstanceInfoMock(status string, playerCount int) *runtime.InstanceInfo {
	instance := runtime.InstanceInfo{
		Id: "1234567890",
		ConnectionInfo: &runtime.ConnectionInfo{
			IpAddress: "127.0.0.1",
			DnsName:   "127.0.0.1",
			Port:      8080,
		},
		PlayerCount: playerCount,
		Status:      status,
		Metadata:    map[string]any{MaxPlayers: 10},
	}

	instance.Status = status

	return &instance
}

func (suite *FleetManagerSuite) TestUpdate_ShouldUpdate() {
	// assert
	const playerCount = 10

	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], playerCount)
	metaData := map[string]any{"newKey": "newValue"}

	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().UpdateApplicationInstance(gomock.Any(), expected.Id, metaData).DoAndReturn(func(ctx context.Context, instanceID string, meta map[string]any) (*runtime.InstanceInfo, error) {
		cp := *expected
		cp.Metadata = meta
		return &cp, nil
	}).Times(1)

	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, inst []*runtime.InstanceInfo) error {
		suite.Equal(1, len(inst))
		suite.Equal(len(metaData), len(inst[0].Metadata))
		suite.Equal(expected.Id, inst[0].Id)
		suite.Equal(expected.PlayerCount, inst[0].PlayerCount)
		suite.Equal(expected.Status, inst[0].Status)
		suite.Equal(metaData, inst[0].Metadata)
		return nil
	}).Times(1)

	// act
	err := suite.newTestFleetManager(client, storageService).Update(suite.ctx, expected.Id, expected.PlayerCount, metaData)

	// assert
	suite.NoError(err)
}

func (suite *FleetManagerSuite) TestUpdateGivenApiError_ShouldReturnError() {
	// assert
	const playerCount = 10

	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], playerCount)
	metaData := map[string]any{"newKey": "newValue"}

	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().UpdateApplicationInstance(gomock.Any(), expected.Id, metaData).Return(nil, errors.New("failed")).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), gomock.Any()).Return(nil).Times(0)

	// act
	err := suite.newTestFleetManager(client, storageService).Update(suite.ctx, expected.Id, expected.PlayerCount, metaData)

	// assert
	suite.Error(err)
}

func (suite *FleetManagerSuite) TestUpdateGivenStorageError_ShouldReturnError() {
	// assert
	const playerCount = 10

	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], playerCount)
	metaData := map[string]any{"newKey": "newValue"}

	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().UpdateApplicationInstance(gomock.Any(), expected.Id, metaData).Return(expected, nil).Times(1)
	storageService.EXPECT().UpdateStorageGameSession(gomock.Any(), gomock.Any()).Return(errors.New("failed")).Times(1)

	// act
	err := suite.newTestFleetManager(client, storageService).Update(suite.ctx, expected.Id, expected.PlayerCount, metaData)

	// assert
	suite.Error(err)
}

func (suite *FleetManagerSuite) TestDelete_ShouldDelete() {
	// assert
	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0)

	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().RestartApplicationInstance(gomock.Any(), expected.Id).Return(nil).Times(1)
	storageService.EXPECT().DeleteStorageGameSession(gomock.Any(), []string{expected.Id}).Return(nil).Times(1)

	// act
	err := suite.newTestFleetManager(client, storageService).Delete(suite.ctx, expected.Id)

	// assert
	suite.NoError(err)
}

func (suite *FleetManagerSuite) TestDelete_GivenAnApiError_ShouldReturnError() {
	// assert
	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0)

	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().RestartApplicationInstance(gomock.Any(), expected.Id).Return(errors.New("failed")).Times(1)
	storageService.EXPECT().DeleteStorageGameSession(gomock.Any(), []string{expected.Id}).Return(nil).Times(0)

	// act
	err := suite.newTestFleetManager(client, storageService).Delete(suite.ctx, expected.Id)

	// assert
	suite.Error(err)
}

func (suite *FleetManagerSuite) TestDelete_GivenAnStorageError_ShouldReturnError() {
	// assert
	expected := InstanceInfoMock(clients.ApplicationInstanceStatus[5], 0)

	ctrl := gomock.NewController(suite.T())
	client := mock.NewMockApplicationInstance(ctrl)
	storageService := tests.NewMockFleetManagerStorage(ctrl)
	defer ctrl.Finish()

	client.EXPECT().RestartApplicationInstance(gomock.Any(), expected.Id).Return(nil).Times(1)
	storageService.EXPECT().DeleteStorageGameSession(gomock.Any(), []string{expected.Id}).Return(errors.New("failed")).Times(1)

	// act
	err := suite.newTestFleetManager(client, storageService).Delete(suite.ctx, expected.Id)

	// assert
	suite.Error(err)
}

func TestFleetManager(t *testing.T) {
	suite.Run(t, new(FleetManagerSuite))
}
