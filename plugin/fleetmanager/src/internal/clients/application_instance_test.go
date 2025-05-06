package clients

import (
	"strconv"
	"testing"

	"bytes"
	"context"
	"encoding/json"
	"github.com/bxcodec/faker/v4"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/suite"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/config"
	openapi "gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/openapi"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/tests"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
)

func newOneApiTestClient(cfg *config.Config, authentication Authentication, httpClient *http.Client, logger runtime.Logger) *OneApiClient {
	return &OneApiClient{
		cfg:            cfg,
		logger:         logger,
		authentication: authentication,
		httpClient:     httpClient,
	}
}

type mockRoundTripper struct {
	response *http.Response
	err      error
}

func newMockRoundTripper(response *http.Response, err error) *mockRoundTripper {
	return &mockRoundTripper{
		response: response,
		err:      err,
	}
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.response, m.err
}

type ApplicationInstanceTestSuite struct {
	suite.Suite
	logger runtime.Logger
	cfg    *config.Config
}

func (suite *ApplicationInstanceTestSuite) SetupTest() {

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

func (suite *ApplicationInstanceTestSuite) NewApiClient(roundTrip *mockRoundTripper, auth Authentication) *OneApiClient {
	suite.logger = tests.NewMockLogger()
	httpClient := &http.Client{
		Transport: roundTrip,
	}
	return newOneApiTestClient(suite.cfg, auth, httpClient, suite.logger)
}

func (suite *ApplicationInstanceTestSuite) TestGetApplicationInstance_success() {
	// arrange
	ctrl := gomock.NewController(suite.T())
	auth := tests.NewMockAuthentication(ctrl)
	defer ctrl.Finish()

	// setting up mock
	auth.EXPECT().IsExpired().Return(false).AnyTimes()
	auth.EXPECT().GetAccessToken().Times(0)
	const port = 1234
	const ipAddress = "127.0.0.1"
	// create a mock response
	instance := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: ipAddress,
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: strconv.Itoa(port),
			PropertyKey:   "port",
			Id:            strconv.Itoa(port),
			PropertyType:  1,
		},
	})
	response := ToHttpResponse([]openapi.ApplicationInstance{instance})

	client := suite.NewApiClient(newMockRoundTripper(response, nil), auth)

	// act
	appInstance, err := client.GetApplicationInstance(context.TODO(), instance.Id)

	// assert
	suite.NoError(err)
	suite.NotNil(appInstance)
	suite.Equal(instance.Id, appInstance.Id)
	suite.Equal(port, appInstance.ConnectionInfo.Port)
	suite.Equal(ipAddress, appInstance.ConnectionInfo.IpAddress)
}

func (suite *ApplicationInstanceTestSuite) TestGetApplicationInstance_whenPrivateAndPublicIp_ShouldReturnOnlyPublic() {
	// arrange
	ctrl := gomock.NewController(suite.T())
	auth := tests.NewMockAuthentication(ctrl)
	defer ctrl.Finish()

	// setting up mock
	auth.EXPECT().IsExpired().Return(false).AnyTimes()
	auth.EXPECT().GetAccessToken().Times(0)
	const port = 1234
	const ipAddress = "90.80.10.12"
	// create a mock response
	instance := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: ipAddress,
			IpVersion: 4,
		},
		{
			Private:   1,
			IpAddress: "127.0.0.1",
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: strconv.Itoa(port),
			PropertyKey:   "port",
			Id:            strconv.Itoa(port),
			PropertyType:  1,
		},
	})
	response := ToHttpResponse([]openapi.ApplicationInstance{instance})

	client := suite.NewApiClient(newMockRoundTripper(response, nil), auth)

	// act
	appInstance, err := client.GetApplicationInstance(context.TODO(), instance.Id)

	// assert
	suite.NoError(err)
	suite.NotNil(appInstance)
	suite.Equal(instance.Id, appInstance.Id)
	suite.Equal(port, appInstance.ConnectionInfo.Port)
	suite.Equal(ipAddress, appInstance.ConnectionInfo.IpAddress)
}

func (suite *ApplicationInstanceTestSuite) TestGetApplicationInstance_failed() {
	// arrange
	ctrl := gomock.NewController(suite.T())
	auth := tests.NewMockAuthentication(ctrl)
	defer ctrl.Finish()

	// setting up mock
	auth.EXPECT().IsExpired().Return(false).AnyTimes()
	auth.EXPECT().GetAccessToken().Times(0)

	// create a mock response
	instance := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: "127.0.0.1",
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: "1245",
			PropertyKey:   "port",
			Id:            "1234",
			PropertyType:  1,
		},
	})

	for _, statusCode := range suite.StatusCodes() {
		response := ToHttpResponse([]openapi.ApplicationInstance{instance})
		response.StatusCode = statusCode
		client := suite.NewApiClient(newMockRoundTripper(response, nil), auth)

		// act
		appInstance, err := client.GetApplicationInstance(context.TODO(), instance.Id)

		// assert
		suite.NotNil(err)
		suite.Nil(appInstance)
	}
}

func (suite *ApplicationInstanceTestSuite) TestListApplicationInstances_success() {
	// arrange
	ctrl := gomock.NewController(suite.T())
	auth := tests.NewMockAuthentication(ctrl)
	defer ctrl.Finish()

	// setting up mock
	auth.EXPECT().IsExpired().Return(false).AnyTimes()
	auth.EXPECT().GetAccessToken().Times(0)
	const port1 = 1234
	const ipAddress1 = "127.0.0.1"
	const port2 = 1235
	const ipAddress2 = "127.0.0.2"
	// create a mock response
	instance1 := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: ipAddress1,
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: strconv.Itoa(port1),
			PropertyKey:   "port",
			Id:            strconv.Itoa(port1),
			PropertyType:  1,
		},
	})

	instance2 := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: ipAddress2,
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: strconv.Itoa(port2),
			PropertyKey:   "port",
			Id:            strconv.Itoa(port2),
			PropertyType:  1,
		},
	})
	response := ToHttpResponse([]openapi.ApplicationInstance{instance1, instance2})

	client := suite.NewApiClient(newMockRoundTripper(response, nil), auth)

	// act
	appResponse, err := client.ListApplicationInstances(context.TODO(), "", 2, "")

	// assert
	suite.NoError(err)
	suite.NotNil(appResponse)
	suite.Equal(2, len(appResponse.Instances))
	suite.Equal(instance1.Id, appResponse.Instances[0].Id)
	suite.Equal(port1, appResponse.Instances[0].ConnectionInfo.Port)
	suite.Equal(ipAddress1, appResponse.Instances[0].ConnectionInfo.IpAddress)

	suite.Equal(instance2.Id, appResponse.Instances[1].Id)
	suite.Equal(port2, appResponse.Instances[1].ConnectionInfo.Port)
	suite.Equal(ipAddress2, appResponse.Instances[1].ConnectionInfo.IpAddress)
}

func (suite *ApplicationInstanceTestSuite) TestListApplicationInstances_failed() {
	// arrange
	ctrl := gomock.NewController(suite.T())
	auth := tests.NewMockAuthentication(ctrl)
	defer ctrl.Finish()

	// setting up mock
	auth.EXPECT().IsExpired().Return(false).AnyTimes()
	auth.EXPECT().GetAccessToken().Times(0)
	const port1 = 1234
	const ipAddress1 = "127.0.0.1"
	const port2 = 1235
	const ipAddress2 = "127.0.0.2"
	// create a mock response
	instance1 := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: ipAddress1,
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: strconv.Itoa(port1),
			PropertyKey:   "port",
			Id:            strconv.Itoa(port1),
			PropertyType:  1,
		},
	})

	instance2 := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: ipAddress2,
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: strconv.Itoa(port2),
			PropertyKey:   "port",
			Id:            strconv.Itoa(port2),
			PropertyType:  1,
		},
	})
	response := ToHttpResponse([]openapi.ApplicationInstance{instance1, instance2})

	// act
	for _, statusCode := range suite.StatusCodes() {
		response.StatusCode = statusCode
		client := suite.NewApiClient(newMockRoundTripper(response, nil), auth)

		// act
		appResponse, err := client.ListApplicationInstances(context.TODO(), "", 2, "")

		// assert
		suite.NotNil(err)
		suite.Nil(appResponse)
	}
}

func (suite *ApplicationInstanceTestSuite) TestAllocateApplicationInstance_success() {
	// arrange
	ctrl := gomock.NewController(suite.T())
	auth := tests.NewMockAuthentication(ctrl)
	defer ctrl.Finish()

	// setting up mock
	auth.EXPECT().IsExpired().Return(false).AnyTimes()
	auth.EXPECT().GetAccessToken().Times(0)
	const port = 1234
	const ipAddress = "127.0.0.1"
	// create a mock response
	instance := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: ipAddress,
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: strconv.Itoa(port),
			PropertyKey:   "port",
			Id:            strconv.Itoa(port),
			PropertyType:  1,
		},
	})

	// setting up mock metadata
	const key = "key"
	const value = "value"
	instance.Metadata = []openapi.Metadata{
		{
			Key:   key,
			Value: value,
		},
	}

	response := ToHttpResponse([]openapi.ApplicationInstance{instance})
	client := suite.NewApiClient(newMockRoundTripper(response, nil), auth)
	metaData := make(map[string]any)
	metaData[key] = value
	// act
	appInstance, err := client.AllocateApplicationInstance(context.TODO(), metaData, "")

	// assert
	suite.NoError(err)
	suite.NotNil(appInstance)
	suite.Equal(instance.Id, appInstance.Id)
	suite.Equal(port, appInstance.ConnectionInfo.Port)
	suite.Equal(ipAddress, appInstance.ConnectionInfo.IpAddress)
	suite.Equal(appInstance.Metadata, metaData)
}

func (suite *ApplicationInstanceTestSuite) TestAllocateApplicationInstance_failed() {
	// arrange
	ctrl := gomock.NewController(suite.T())
	auth := tests.NewMockAuthentication(ctrl)
	defer ctrl.Finish()

	// setting up mock
	auth.EXPECT().IsExpired().Return(false).AnyTimes()
	auth.EXPECT().GetAccessToken().Times(0)
	const port = 1234
	const ipAddress = "127.0.0.1"

	// create a mock response
	instance := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: ipAddress,
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: strconv.Itoa(port),
			PropertyKey:   "port",
			Id:            strconv.Itoa(port),
			PropertyType:  1,
		},
	})
	response := ToHttpResponse([]openapi.ApplicationInstance{instance})
	const key = "key"
	const value = "value"
	metaData := make(map[string]any)
	metaData[key] = value
	// act
	for _, statusCode := range suite.StatusCodes() {
		response.StatusCode = statusCode
		client := suite.NewApiClient(newMockRoundTripper(response, nil), auth)

		// act
		appResponse, err := client.AllocateApplicationInstance(context.TODO(), metaData, "")

		// assert
		suite.NotNil(err)
		suite.Nil(appResponse)
	}
}

func (suite *ApplicationInstanceTestSuite) TestRestartApplicationInstance_success() {
	// arrange
	ctrl := gomock.NewController(suite.T())
	auth := tests.NewMockAuthentication(ctrl)
	defer ctrl.Finish()

	// setting up mock
	auth.EXPECT().IsExpired().Return(false).AnyTimes()
	auth.EXPECT().GetAccessToken().Times(0)
	const port = 1234
	const ipAddress = "90.80.10.12"
	// create a mock response
	instance := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: ipAddress,
			IpVersion: 4,
		},
		{
			Private:   1,
			IpAddress: "127.0.0.1",
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: strconv.Itoa(port),
			PropertyKey:   "port",
			Id:            strconv.Itoa(port),
			PropertyType:  1,
		},
	})
	response := ToHttpResponse([]openapi.ApplicationInstance{instance})

	client := suite.NewApiClient(newMockRoundTripper(response, nil), auth)

	// act
	err := client.RestartApplicationInstance(context.TODO(), instance.Id)

	// assert
	suite.NoError(err)
}

func (suite *ApplicationInstanceTestSuite) TestRestartApplicationInstance_failed() {
	// arrange
	ctrl := gomock.NewController(suite.T())
	auth := tests.NewMockAuthentication(ctrl)
	defer ctrl.Finish()

	// setting up mock
	auth.EXPECT().IsExpired().Return(false).AnyTimes()
	auth.EXPECT().GetAccessToken().Times(0)

	// create a mock response
	instance := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: "127.0.0.1",
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: "1245",
			PropertyKey:   "port",
			Id:            "1234",
			PropertyType:  1,
		},
	})

	for _, statusCode := range suite.StatusCodes() {
		response := ToHttpResponse([]openapi.ApplicationInstance{instance})
		response.StatusCode = statusCode
		client := suite.NewApiClient(newMockRoundTripper(response, nil), auth)

		// act
		err := client.RestartApplicationInstance(context.TODO(), instance.Id)

		// assert
		suite.NotNil(err)
	}
}

func (suite *ApplicationInstanceTestSuite) TestUpdateApplicationInstance_success() {
	// arrange
	ctrl := gomock.NewController(suite.T())
	auth := tests.NewMockAuthentication(ctrl)
	defer ctrl.Finish()

	// setting up mock
	auth.EXPECT().IsExpired().Return(false).AnyTimes()
	auth.EXPECT().GetAccessToken().Times(0)
	const port = 1234
	const ipAddress = "127.0.0.1"
	// create a mock response
	instance := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: ipAddress,
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: strconv.Itoa(port),
			PropertyKey:   "port",
			Id:            strconv.Itoa(port),
			PropertyType:  1,
		},
	})

	// setting up mock metadata
	const key = "key"
	const value = "value"
	instance.Metadata = []openapi.Metadata{
		{
			Key:   key,
			Value: value,
		},
	}

	response := ToHttpResponse([]openapi.ApplicationInstance{instance})
	client := suite.NewApiClient(newMockRoundTripper(response, nil), auth)
	metaData := make(map[string]any)
	metaData[key] = value

	// act
	appInstance, err := client.UpdateApplicationInstance(context.TODO(), instance.Id, metaData)

	// assert
	suite.NoError(err)
	suite.NotNil(appInstance)
	suite.Equal(instance.Id, appInstance.Id)
	suite.Equal(port, appInstance.ConnectionInfo.Port)
	suite.Equal(ipAddress, appInstance.ConnectionInfo.IpAddress)
	suite.Equal(appInstance.Metadata, metaData)
}

func (suite *ApplicationInstanceTestSuite) TestUpdateApplicationInstance_failed() {
	// arrange
	ctrl := gomock.NewController(suite.T())
	auth := tests.NewMockAuthentication(ctrl)
	defer ctrl.Finish()

	// setting up mock
	auth.EXPECT().IsExpired().Return(false).AnyTimes()
	auth.EXPECT().GetAccessToken().Times(0)
	const port = 1234
	const ipAddress = "127.0.0.1"

	// create a mock response
	instance := applicationInstanceMock([]openapi.ApplicationInstanceIP{
		{
			Private:   0,
			IpAddress: ipAddress,
			IpVersion: 4,
		},
	}, []openapi.ApplicationInstanceProperty{
		{
			PropertyValue: strconv.Itoa(port),
			PropertyKey:   "port",
			Id:            strconv.Itoa(port),
			PropertyType:  1,
		},
	})

	response := ToHttpResponse([]openapi.ApplicationInstance{instance})
	const key = "key"
	const value = "value"
	metaData := make(map[string]any)
	metaData[key] = value
	// act
	for _, statusCode := range suite.StatusCodes() {
		response.StatusCode = statusCode
		client := suite.NewApiClient(newMockRoundTripper(response, nil), auth)

		// act
		appInstance, err := client.UpdateApplicationInstance(context.TODO(), instance.Id, metaData)

		// assert
		suite.NotNil(err)
		suite.Nil(appInstance)
	}
}

func applicationInstanceMock(ipAddresses []openapi.ApplicationInstanceIP, properties []openapi.ApplicationInstanceProperty) openapi.ApplicationInstance {
	var instance openapi.ApplicationInstance
	err := faker.FakeData(&instance)
	if err != nil {
		panic(err)
	}

	instance.IpAddress = ipAddresses
	instance.Properties = properties

	return instance
}

func ToHttpResponse(data interface{}) *http.Response {
	// Convert the data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(jsonData)),
		Header:     make(http.Header),
	}
	resp.Header.Set("Content-Type", "application/json")
	return resp
}

func (suite *ApplicationInstanceTestSuite) StatusCodes() []int {
	return []int{http.StatusInternalServerError, http.StatusBadGateway, http.StatusBadRequest, http.StatusNotFound, http.StatusConflict}
}

func TestApplicationInstance(t *testing.T) {
	suite.Run(t, new(ApplicationInstanceTestSuite))
}
