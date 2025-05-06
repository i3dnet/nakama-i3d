package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/fleetmanager_config"
	"io"
	"net/http"
	"sync"
	"time"
)

type Authentication interface {
	GetAccessToken() (string, error)
	IsExpired() bool
}

type AuthenticationService struct {
	cfg         *fleetmanager_config.Config
	accessToken string
	expiresAt   time.Time
	mu          sync.Mutex
	httpClient  *http.Client
}

func NewAuthentication(cfg *fleetmanager_config.Config) *AuthenticationService {
	return &AuthenticationService{
		cfg:         cfg,
		accessToken: "",
		expiresAt:   time.Now(),
		mu:          sync.Mutex{},
		httpClient:  http.DefaultClient,
	}
}

func (a *AuthenticationService) GetAccessToken() (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.IsExpired() {
		return a.accessToken, nil
	}

	url := a.cfg.OneApi.BaseUrl + "/oauth/token"

	body := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     a.cfg.OneApi.ClientId,
		"client_secret": a.cfg.OneApi.ClientSecret,
		"audience":      a.cfg.OneApi.Audience,
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return "", err
	}

	req, err := a.httpClient.Post(url, "application/json", bytes.NewBuffer(bodyBytes))

	if err != nil {
		return "", err
	}

	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get access token, with statusCode %s", req.Status)
	}

	responseBody, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	var tokenResp tokenResponse
	err = json.Unmarshal(responseBody, &tokenResp)
	if err != nil {
		return "", err
	}

	a.accessToken = tokenResp.AccessToken
	a.expiresAt = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
	return a.accessToken, nil
}

func (a *AuthenticationService) IsExpired() bool {
	if a.cfg.OneApi.UseBearerAuth == false {
		return false
	}

	if a.accessToken == "" {
		return true
	}

	return time.Now().After(a.expiresAt.Add(-30 * time.Second))
}
