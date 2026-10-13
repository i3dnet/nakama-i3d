package clients

import (
	"context"
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/config"
	"github.com/i3dnet/nakama-i3d/internal/openapi"
	"net/http"
	"time"
)

type OneApiClient struct {
	httpClient     *http.Client
	cfg            *config.Config
	authentication Authentication
	RetryExecutor  *RetryExecutor
	logger         runtime.Logger
}
type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func NewOneApiClient(cfg *config.Config, authentication Authentication, logger runtime.Logger) *OneApiClient {
	snapshot := *cfg
	if snapshot.ProviderTimeout == 0 {
		snapshot.ProviderTimeout = 90 * time.Second
	}
	if snapshot.Attempts == 0 {
		snapshot.Attempts = 3
	}
	return &OneApiClient{cfg: &snapshot, logger: logger, authentication: authentication, httpClient: &http.Client{Timeout: snapshot.ProviderTimeout}, RetryExecutor: NewRetryExecutor(logger, &snapshot, retryableRead)}
}

// GetClient creates an immutable request client while reusing the HTTP connection
// pool. Refreshed credentials never mutate a client already in use by another call.
func (o *OneApiClient) GetClient(ctx context.Context) (*openapi.APIClient, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	conf := openapi.NewConfiguration()
	conf.HTTPClient = o.httpClient
	conf.Servers = openapi.ServerConfigurations{{URL: o.cfg.BaseUrl}}
	if o.cfg.UseBearerAuth {
		if o.authentication == nil {
			return nil, fmt.Errorf("OAuth authentication is unavailable")
		}
		token, err := o.authentication.GetAccessToken(ctx)
		if err != nil {
			return nil, err
		}
		conf.DefaultHeader["Authorization"] = "Bearer " + token
	} else {
		conf.DefaultHeader["PRIVATE-TOKEN"] = o.cfg.Token
	}
	return openapi.NewAPIClient(conf), nil
}
func executeRead[T any](ctx context.Context, o *OneApiClient, request func() (T, *http.Response, error)) (T, *http.Response, error) {
	var value T
	var response *http.Response
	err := o.RetryExecutor.Run(ctx, o.cfg.Attempts, func() error {
		var err error
		value, response, err = request()
		if err != nil && response != nil {
			return &providerHTTPError{code: response.StatusCode, err: err}
		}
		return err
	})
	return value, response, err
}
