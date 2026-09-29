package clients

import (
	"context"
	"errors"
	"testing"
	"testing/synctest"
	"time"

	"github.com/i3dnet/nakama-i3d/config"
	"github.com/i3dnet/nakama-i3d/internal/tests"
	"github.com/stretchr/testify/require"
)

func TestZeroInitialRetryDelayGrowsWithinMaximum(t *testing.T) {
	for _, tc := range []struct {
		name     string
		maxDelay time.Duration
		calls    []time.Duration
	}{
		{"positive maximum", 4 * time.Millisecond, []time.Duration{0, 0, time.Millisecond, 3 * time.Millisecond, 7 * time.Millisecond, 11 * time.Millisecond}},
		{"submillisecond maximum", 500 * time.Microsecond, []time.Duration{0, 0, 500 * time.Microsecond, time.Millisecond, 1500 * time.Microsecond, 2 * time.Millisecond}},
		{"explicitly disabled backoff", 0, []time.Duration{0, 0, 0, 0, 0, 0}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				start := time.Now()
				var called []time.Duration
				failure := errors.New("temporary failure")
				retry := NewRetryExecutor(tests.NewMockLogger(), &config.Config{Retry: config.Retry{Delay: 0, MaxDelay: tc.maxDelay}}, func(error) bool { return true })
				err := retry.Run(context.Background(), len(tc.calls), func() error {
					called = append(called, time.Since(start))
					return failure
				})
				require.ErrorIs(t, err, failure)
				require.Equal(t, tc.calls, called)
			})
		})
	}
}
