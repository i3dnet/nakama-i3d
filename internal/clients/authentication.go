package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/i3dnet/nakama-i3d/config"
	"io"
	"math"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Authentication interface {
	GetAccessToken(context.Context) (string, error)
	IsExpired() bool
}
type tokenRefresh struct {
	done  chan struct{}
	token string
	err   error
}
type AuthenticationService struct {
	cfg         *config.Config
	accessToken string
	expiresAt   time.Time
	expirySkew  time.Duration
	mu          sync.Mutex
	refresh     *tokenRefresh
	httpClient  *http.Client
}

func NewAuthentication(cfg *config.Config) *AuthenticationService {
	snapshot := *cfg
	return &AuthenticationService{cfg: &snapshot, httpClient: &http.Client{Timeout: 30 * time.Second}}
}
func (a *AuthenticationService) expiredLocked() bool {
	if !a.cfg.UseBearerAuth {
		return false
	}
	return a.accessToken == "" || !time.Now().Before(a.expiresAt.Add(-a.expirySkew))
}
func (a *AuthenticationService) IsExpired() bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.expiredLocked()
}
func (a *AuthenticationService) GetAccessToken(ctx context.Context) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	a.mu.Lock()
	if !a.expiredLocked() {
		token := a.accessToken
		a.mu.Unlock()
		return token, nil
	}
	if pending := a.refresh; pending != nil {
		a.mu.Unlock()
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-pending.done:
			return pending.token, pending.err
		}
	}
	pending := &tokenRefresh{done: make(chan struct{})}
	a.refresh = pending
	a.mu.Unlock()
	token, lifetime, err := a.requestToken(ctx)
	a.mu.Lock()
	if err == nil {
		a.accessToken = token
		a.expiresAt = time.Now().Add(lifetime)
		a.expirySkew = min(30*time.Second, lifetime/10)
	}
	pending.token, pending.err = token, err
	a.refresh = nil
	close(pending.done)
	a.mu.Unlock()
	return token, err
}
func (a *AuthenticationService) requestToken(ctx context.Context) (string, time.Duration, error) {
	body, err := json.Marshal(map[string]string{"grant_type": "client_credentials", "client_id": a.cfg.ClientId, "client_secret": a.cfg.ClientSecret, "audience": a.cfg.Audience})
	if err != nil {
		return "", 0, err
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, a.cfg.AuthenticationUrl, bytes.NewReader(body))
	if err != nil {
		return "", 0, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := a.httpClient.Do(request)
	if err != nil {
		return "", 0, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("OAuth request failed with HTTP %d", response.StatusCode)
	}
	var token tokenResponse
	if err = json.NewDecoder(io.LimitReader(response.Body, 64<<10)).Decode(&token); err != nil {
		return "", 0, fmt.Errorf("invalid OAuth response: %w", err)
	}
	if token.AccessToken == "" || token.ExpiresIn <= 0 || int64(token.ExpiresIn) > math.MaxInt64/int64(time.Second) || (token.TokenType != "" && !strings.EqualFold(token.TokenType, "Bearer")) {
		return "", 0, fmt.Errorf("OAuth response has an invalid token, expiry or token type")
	}
	return token.AccessToken, time.Duration(token.ExpiresIn) * time.Second, nil
}
