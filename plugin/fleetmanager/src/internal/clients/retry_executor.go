package clients

import (
	"github.com/heroiclabs/nakama-common/runtime"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/fleetmanager_config"
	"time"
)

type RetryExecutor struct {
	log          runtime.Logger
	shouldRetry  func(error) bool
	initialDelay time.Duration
	maxDelay     time.Duration
	factor       float64
}

func NewRetryExecutor(log runtime.Logger, cfg *fleetmanager_config.Config, shouldRetry func(error) bool) *RetryExecutor {
	return &RetryExecutor{
		log:          log,
		shouldRetry:  shouldRetry,
		initialDelay: cfg.Retry.Delay,
		maxDelay:     cfg.Retry.MaxDelay,
		factor:       2.0,
	}
}

func (r *RetryExecutor) Run(attempts int, fn func() error) error {
	delay := r.initialDelay
	for i := 0; i < attempts-1; i++ {
		err := fn()
		if err == nil {
			return nil
		}

		if !r.shouldRetry(err) {
			return err
		}

		r.log.WithField("error", err.Error()).Error("Retrying in %s (attempt %d/%d)", delay, i+1, attempts)
		time.Sleep(delay)

		delay = time.Duration(float64(delay) * r.factor)
		if delay > r.maxDelay {
			delay = r.maxDelay
		}
	}

	// final attempt
	return fn()
}
