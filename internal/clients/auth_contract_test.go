package clients

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/i3dnet/nakama-i3d/config"
	"github.com/i3dnet/nakama-i3d/internal/openapi"
	"github.com/i3dnet/nakama-i3d/internal/tests"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func oauthConfig(base, auth string) *config.Config {
	return &config.Config{OneApi: config.OneApi{BaseUrl: base, AuthenticationUrl: auth, UseBearerAuth: true, ClientId: "client", ClientSecret: "test-only", Audience: "audience"}, Retry: config.Retry{Attempts: 3}}
}
func TestOAuthUsesConfiguredEndpoint(t *testing.T) {
	var calls atomic.Int32
	authServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		if r.URL.Path != "/custom/token" {
			t.Errorf("wrong token path: %s", r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "test-token", "expires_in": 3600, "token_type": "Bearer"})
	}))
	defer authServer.Close()
	apiServer := httptest.NewServer(http.NotFoundHandler())
	defer apiServer.Close()
	auth := NewAuthentication(oauthConfig(apiServer.URL, authServer.URL+"/custom/token"))
	token, err := auth.GetAccessToken(context.Background())
	require.NoError(t, err)
	require.Equal(t, "test-token", token)
	require.EqualValues(t, 1, calls.Load())
}
func TestOAuthFailureDoesNotPanic(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusUnauthorized) }))
	defer server.Close()
	cfg := oauthConfig(server.URL, server.URL+"/token")
	client := NewOneApiClient(cfg, NewAuthentication(cfg), tests.NewMockLogger())
	require.NotPanics(t, func() { got, err := client.GetClient(context.Background()); require.Error(t, err); require.Nil(t, got) })
}
func TestConcurrentProviderCallsInitializeSafely(t *testing.T) {
	client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
		writeInstances(w, []openapi.ApplicationInstance{validProviderInstance()})
	})
	start := make(chan struct{})
	results := make(chan error, 20)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			_, err := client.GetApplicationInstance(context.Background(), "instance-1")
			results <- err
		}()
	}
	close(start)
	wg.Wait()
	close(results)
	for err := range results {
		require.NoError(t, err)
	}
}
func TestAllocationFailureIsNotBlindlyRetried(t *testing.T) {
	var calls atomic.Int32
	client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.WriteHeader(http.StatusServiceUnavailable)
	})
	client.cfg.Attempts = 3
	_, err := client.AllocateApplicationInstance(context.Background(), nil, "")
	require.Error(t, err)
	require.EqualValues(t, 1, calls.Load())
}
func TestReadRetriesTransientFailure(t *testing.T) {
	var calls atomic.Int32
	client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) == 1 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		writeInstances(w, []openapi.ApplicationInstance{validProviderInstance()})
	})
	client.cfg.Attempts = 3
	_, err := client.GetApplicationInstance(context.Background(), "instance-1")
	require.NoError(t, err)
	require.EqualValues(t, 2, calls.Load())
}
func TestOAuthRejectsInvalidTokenResponses(t *testing.T) {
	for _, body := range []string{`{"access_token":"","expires_in":3600}`, `{"access_token":"token","expires_in":0}`, `{"access_token":"token","expires_in":-1}`} {
		t.Run(body, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte(body)) }))
			defer server.Close()
			auth := NewAuthentication(oauthConfig(server.URL, server.URL+"/token"))
			_, err := auth.GetAccessToken(context.Background())
			require.Error(t, err)
		})
	}
}
func TestShortLivedTokenIsCached(t *testing.T) {
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		_, _ = w.Write([]byte(`{"access_token":"short","expires_in":1}`))
	}))
	defer server.Close()
	auth := NewAuthentication(oauthConfig(server.URL, server.URL+"/token"))
	_, err := auth.GetAccessToken(context.Background())
	require.NoError(t, err)
	_, err = auth.GetAccessToken(context.Background())
	require.NoError(t, err)
	require.EqualValues(t, 1, calls.Load())
}
func TestOAuthSharesRefreshAndCancelsWaiter(t *testing.T) {
	entered, release := make(chan struct{}), make(chan struct{})
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) == 1 {
			close(entered)
		}
		<-release
		_, _ = w.Write([]byte(`{"access_token":"shared","expires_in":3600}`))
	}))
	defer server.Close()
	auth := NewAuthentication(oauthConfig(server.URL, server.URL))
	result := make(chan error, 1)
	go func() { _, err := auth.GetAccessToken(context.Background()); result <- err }()
	<-entered
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err := auth.GetAccessToken(ctx)
	require.ErrorIs(t, err, context.DeadlineExceeded)
	var wg sync.WaitGroup
	errs := make(chan error, 20)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := auth.GetAccessToken(context.Background())
			errs <- err
			_ = auth.IsExpired()
		}()
	}
	close(release)
	require.NoError(t, <-result)
	wg.Wait()
	close(errs)
	for err := range errs {
		require.NoError(t, err)
	}
	require.EqualValues(t, 1, calls.Load())
}
func TestOAuthRefreshChangesRequestAuthorization(t *testing.T) {
	var tokens atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/token" {
			n := tokens.Add(1)
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": fmt.Sprintf("token-%d", n), "expires_in": 3600})
			return
		}
		if got := r.Header.Get("Authorization"); got != fmt.Sprintf("Bearer token-%d", tokens.Load()) {
			t.Errorf("stale Authorization: %q", got)
		}
		writeInstances(w, []openapi.ApplicationInstance{validProviderInstance()})
	}))
	defer server.Close()
	cfg := oauthConfig(server.URL, server.URL+"/token")
	auth := NewAuthentication(cfg)
	client := NewOneApiClient(cfg, auth, tests.NewMockLogger())
	_, err := client.GetApplicationInstance(context.Background(), "instance-1")
	require.NoError(t, err)
	auth.mu.Lock()
	auth.expiresAt = time.Now().Add(-time.Second)
	auth.mu.Unlock()
	_, err = client.GetApplicationInstance(context.Background(), "instance-1")
	require.NoError(t, err)
	require.EqualValues(t, 2, tokens.Load())
}
func TestProviderTimeoutAndNonRetryableReads(t *testing.T) {
	t.Run("timeout", func(t *testing.T) {
		var calls atomic.Int32
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { calls.Add(1); <-r.Context().Done() }))
		defer server.Close()
		cfg := &config.Config{OneApi: config.OneApi{BaseUrl: server.URL}, ProviderTimeout: 20 * time.Millisecond, Retry: config.Retry{Attempts: 3}}
		client := NewOneApiClient(cfg, nil, tests.NewMockLogger())
		_, err := client.GetApplicationInstance(context.Background(), "instance-1")
		require.Error(t, err)
		require.EqualValues(t, 1, calls.Load())
	})
	for _, status := range []int{401, 404} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var calls atomic.Int32
			client := contractClient(t, func(w http.ResponseWriter, r *http.Request) { calls.Add(1); w.WriteHeader(status) })
			_, err := client.GetApplicationInstance(context.Background(), "instance-1")
			require.Error(t, err)
			require.EqualValues(t, 1, calls.Load())
		})
	}
}
func TestRetryCancellationDuringBackoff(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	retry := NewRetryExecutor(tests.NewMockLogger(), &config.Config{Retry: config.Retry{Delay: time.Hour, MaxDelay: time.Hour}}, func(error) bool { return true })
	var calls int
	err := retry.Run(ctx, 3, func() error { calls++; cancel(); return errors.New("transient") })
	require.ErrorIs(t, err, context.Canceled)
	require.Equal(t, 1, calls)
	require.Error(t, retry.Run(context.Background(), 0, func() error { t.Fatal("zero attempts invoked request"); return nil }))
}
func TestOAuthRequestHonorsCancellation(t *testing.T) {
	entered := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
		close(entered)
		<-r.Context().Done()
	}))
	defer server.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	auth := NewAuthentication(oauthConfig(server.URL, server.URL))
	result := make(chan error, 1)
	go func() { _, err := auth.GetAccessToken(ctx); result <- err }()
	<-entered
	cancel()
	select {
	case err := <-result:
		require.ErrorIs(t, err, context.Canceled)
	case <-time.After(time.Second):
		t.Fatal("OAuth request ignored cancellation")
	}
}

func TestReadRetriesTruncatedSuccessfulResponse(t *testing.T) {
	var calls atomic.Int32
	client := contractClient(t, func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) == 1 {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", "100")
			_, _ = w.Write([]byte("["))
			return
		}
		writeInstances(w, []openapi.ApplicationInstance{validProviderInstance()})
	})
	client.cfg.Attempts = 3
	got, err := client.GetApplicationInstance(context.Background(), "instance-1")
	require.NoError(t, err)
	require.NotNil(t, got)
	require.EqualValues(t, 2, calls.Load())
}
