package clients

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/require"
)

type oauthRefreshTransport func(*http.Request) (*http.Response, error)

func (f oauthRefreshTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

type oauthRefreshResult struct {
	token string
	err   error
}

func TestOAuthWaitersRetryCanceledLeader(t *testing.T) {
	for _, leaderErr := range []error{context.Canceled, context.DeadlineExceeded} {
		t.Run(leaderErr.Error(), func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				leaderCtx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				releaseRetry := make(chan struct{})
				var calls atomic.Int32
				auth := NewAuthentication(oauthConfig("http://oauth.test", "http://oauth.test/token"))
				auth.httpClient = &http.Client{Transport: oauthRefreshTransport(func(r *http.Request) (*http.Response, error) {
					if calls.Add(1) == 1 {
						<-r.Context().Done()
						return nil, r.Context().Err()
					}
					<-releaseRetry
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(`{"access_token":"replacement","expires_in":3600}`)),
					}, nil
				})}
				leaderResult := make(chan error, 1)
				go func() { _, err := auth.GetAccessToken(leaderCtx); leaderResult <- err }()
				synctest.Wait()

				const waiters = 20
				results := make(chan oauthRefreshResult, waiters)
				for range waiters {
					go func() {
						token, err := auth.GetAccessToken(context.Background())
						results <- oauthRefreshResult{token, err}
					}()
				}
				// Every waiter has joined the first refresh before its caller ends.
				synctest.Wait()
				if leaderErr == context.Canceled {
					cancel()
				}
				// For the deadline case, the bubble advances its clock while waiting.
				err := <-leaderResult
				synctest.Wait()
				close(releaseRetry)
				require.ErrorIs(t, err, leaderErr)
				for range waiters {
					result := <-results
					require.NoError(t, result.err)
					require.Equal(t, "replacement", result.token)
				}
				require.EqualValues(t, 2, calls.Load(), "waiters must share one replacement refresh")
			})
		})
	}
}

func TestOAuthWaitersShareProviderFailure(t *testing.T) {
	for _, providerErr := range []error{errors.New("provider unavailable"), context.DeadlineExceeded} {
		t.Run(providerErr.Error(), func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				release := make(chan struct{})
				var calls atomic.Int32
				auth := NewAuthentication(oauthConfig("http://oauth.test", "http://oauth.test/token"))
				auth.httpClient = &http.Client{Transport: oauthRefreshTransport(func(*http.Request) (*http.Response, error) {
					calls.Add(1)
					<-release
					return nil, providerErr
				})}
				const callers = 10
				results := make(chan error, callers)
				for range callers {
					go func() { _, err := auth.GetAccessToken(context.Background()); results <- err }()
				}
				synctest.Wait()
				close(release)
				for range callers {
					require.ErrorIs(t, <-results, providerErr)
				}
				require.EqualValues(t, 1, calls.Load(), "provider failure must stay shared while the leader context is live")
			})
		})
	}
}

func TestOAuthCanceledWaiterDoesNotAcceptSharedSuccess(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		waiterCtx, cancel := context.WithCancel(context.Background())
		defer cancel()
		var calls atomic.Int32
		auth := NewAuthentication(oauthConfig("http://oauth.test", "http://oauth.test/token"))
		auth.httpClient = &http.Client{Transport: oauthRefreshTransport(func(*http.Request) (*http.Response, error) {
			calls.Add(1)
			<-release
			cancel()
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader(`{"access_token":"shared","expires_in":3600}`)),
			}, nil
		})}
		leaderResult := make(chan oauthRefreshResult, 1)
		go func() {
			token, err := auth.GetAccessToken(context.Background())
			leaderResult <- oauthRefreshResult{token, err}
		}()
		synctest.Wait()

		waiterResult := make(chan error, 1)
		go func() { _, err := auth.GetAccessToken(waiterCtx); waiterResult <- err }()
		synctest.Wait()
		// Cancel the waiter as the leader completes a successful refresh.
		close(release)
		require.ErrorIs(t, <-waiterResult, context.Canceled)
		result := <-leaderResult
		require.NoError(t, result.err)
		require.Equal(t, "shared", result.token)
		require.EqualValues(t, 1, calls.Load())
	})
}
