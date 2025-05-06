package clients

import (
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/suite"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/fleetmanager_config"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/tests"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func newTestAuthentication(cfg *fleetmanager_config.Config, httpClient *http.Client) *AuthenticationService {
	return &AuthenticationService{
		cfg:         cfg,
		accessToken: "",
		expiresAt:   time.Now(),
		httpClient:  httpClient,
	}
}

type AuthenticationServiceTestSuite struct {
	suite.Suite
	logger runtime.Logger
	cfg    *fleetmanager_config.Config
}

const (
	clientId          = "clientid"
	clientSecret      = "clientsecret"
	audience          = "audience"
	authenticationUrl = "http://localhost:8080"
)

func (suite *AuthenticationServiceTestSuite) SetupTest() {
	suite.cfg = &fleetmanager_config.Config{
		App: fleetmanager_config.App{
			Name:    "test",
			Version: "1.0.1",
		},
		OneApi: fleetmanager_config.OneApi{
			BaseUrl:           "http://localhost:8080",
			ApplicationId:     "1235",
			UseBearerAuth:     true,
			ClientId:          clientId,
			ClientSecret:      clientSecret,
			Audience:          audience,
			AuthenticationUrl: authenticationUrl,
		},
	}
}

func (suite *AuthenticationServiceTestSuite) NewAuthenticationService(roundTrip *mockRoundTripper) *AuthenticationService {
	suite.logger = tests.NewMockLogger()
	httpClient := &http.Client{
		Transport: roundTrip,
	}
	return newTestAuthentication(suite.cfg, httpClient)
}

func (suite *AuthenticationServiceTestSuite) authenticationResponse(token string, expiresIn int) *http.Response {
	body := fmt.Sprintf(`{"access_token":"%s","expires_in":%d}`, token, expiresIn)
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
}

func (suite *AuthenticationServiceTestSuite) TestGetAccessToken_success() {
	// Arrange
	const token = "testtoken"
	const expiresIn = 3600
	response := suite.authenticationResponse(token, expiresIn)
	client := suite.NewAuthenticationService(newMockRoundTripper(response, nil))

	// Act
	result, err := client.GetAccessToken()

	// Assert
	suite.NoError(err)
	suite.Equal(token, result)
}

func (suite *AuthenticationServiceTestSuite) TestGetAccessToken_WhenTokenIsNotExpired_GivesBackOldToken() {
	// Arrange
	const token = "testtoken1"
	const expiresIn = 3600
	response := suite.authenticationResponse("testToken", expiresIn)
	client := suite.NewAuthenticationService(newMockRoundTripper(response, nil))
	client.accessToken = token
	client.expiresAt = time.Now().Add(time.Duration(expiresIn) * time.Second)

	// Act
	result, err := client.GetAccessToken()

	// Assert
	suite.NoError(err)
	suite.Equal(token, result)
}

func (suite *AuthenticationServiceTestSuite) TestGetAccessToken_failed() {
	// Arrange
	const token = "testtoken"
	const expiresIn = 3600
	response := suite.authenticationResponse(token, expiresIn)
	response.StatusCode = http.StatusUnauthorized
	client := suite.NewAuthenticationService(newMockRoundTripper(response, nil))

	// Act
	result, err := client.GetAccessToken()

	// Assert
	suite.NotNil(err)
	suite.Equal("", result)
}

func (suite *AuthenticationServiceTestSuite) TestIsExpired_whenUseBearerTokenIsFalse_shouldReturnFalse() {
	// Arrange
	suite.cfg.OneApi.UseBearerAuth = false
	client := suite.NewAuthenticationService(newMockRoundTripper(nil, nil))

	// Act
	result := client.IsExpired()

	// Assert
	suite.False(result)
}

func (suite *AuthenticationServiceTestSuite) TestIsExpired_whenAccessTokenIsEmpty_shouldReturnTrue() {
	// Arrange
	client := suite.NewAuthenticationService(newMockRoundTripper(nil, nil))

	// Act
	result := client.IsExpired()

	// Assert
	suite.True(result)
}

func (suite *AuthenticationServiceTestSuite) TestIsExpired_whenAccessTokenIsNotEmptyAndExpiresAtIsInThePast_shouldReturnTrue() {
	client := suite.NewAuthenticationService(newMockRoundTripper(nil, nil))
	client.expiresAt = time.Now().Add(-1 * time.Hour)
	client.accessToken = "testtoken"

	// Act
	result := client.IsExpired()

	// Assert
	suite.True(result)
}

func (suite *AuthenticationServiceTestSuite) TestIsExpired_whenAccessTokenIsNotEmptyAndExpiresAtIsInTheFuture_shouldReturnFalse() {
	client := suite.NewAuthenticationService(newMockRoundTripper(nil, nil))
	client.expiresAt = time.Now().Add(1 * time.Hour)
	client.accessToken = "testtoken"

	// Act
	result := client.IsExpired()

	// Assert
	suite.False(result)
}

func TestAuthenticationService(t *testing.T) {
	suite.Run(t, new(AuthenticationServiceTestSuite))
}
