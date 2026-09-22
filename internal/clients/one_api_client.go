package clients

import (
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/config"
	openapi "github.com/i3dnet/nakama-i3d/internal/openapi"
	"net/http"
)

type OneApiClient struct {
	client         *openapi.APIClient
	httpClient     *http.Client
	cfg            *config.Config
	authentication Authentication
	RetryExecutor  *RetryExecutor
	logger         runtime.Logger
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"` // seconds
	TokenType   string `json:"token_type"`
}

func NewOneApiClient(cfg *config.Config, authentication Authentication, logger runtime.Logger) *OneApiClient {
	return &OneApiClient{
		cfg:            cfg,
		logger:         logger,
		authentication: authentication,
		httpClient:     http.DefaultClient,
		RetryExecutor: NewRetryExecutor(logger, cfg, func(err error) bool {

			return true
		}),
	}
}

func (o *OneApiClient) GetClient() *openapi.APIClient {

	if o.client == nil || o.authentication.IsExpired() {
		o.client = o.createClient()
	}

	return o.client
}

func (o *OneApiClient) createClient() *openapi.APIClient {
	conf := openapi.NewConfiguration()
	conf.HTTPClient = o.httpClient
	conf.Servers = openapi.ServerConfigurations{
		{
			URL: o.cfg.OneApi.BaseUrl,
		},
	}

	if o.cfg.UseBearerAuth {
		accToken, err := o.authentication.GetAccessToken()
		if err != nil {
			panic("failed to get access token: " + err.Error())
		}

		conf.DefaultHeader["Authorization"] = "Bearer " + accToken
	} else {
		conf.DefaultHeader["PRIVATE-TOKEN"] = o.cfg.OneApi.Token
	}

	return openapi.NewAPIClient(conf)
}
