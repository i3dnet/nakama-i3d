package clients

import (
	"context"
	"errors"
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/config"
	"io"
	"net"
	"time"
)

type RetryExecutor struct {
	log          runtime.Logger
	shouldRetry  func(error) bool
	initialDelay time.Duration
	maxDelay     time.Duration
}

func NewRetryExecutor(log runtime.Logger, cfg *config.Config, shouldRetry func(error) bool) *RetryExecutor {
	return &RetryExecutor{log: log, shouldRetry: shouldRetry, initialDelay: cfg.Delay, maxDelay: cfg.MaxDelay}
}
func (r *RetryExecutor) Run(ctx context.Context, attempts int, fn func() error) error {
	if attempts < 1 {
		return fmt.Errorf("retry attempts must be positive")
	}
	delay := r.initialDelay
	for attempt := 0; attempt < attempts; attempt++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		err := fn()
		if err == nil {
			return nil
		}
		if attempt == attempts-1 || r.shouldRetry == nil || !r.shouldRetry(err) {
			return err
		}
		r.log.Warn("retrying provider read (attempt %d/%d)", attempt+1, attempts)
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}
		if delay > r.maxDelay/2 {
			delay = r.maxDelay
		} else {
			delay *= 2
		}
	}
	return nil
}

type providerHTTPError struct {
	code int
	err  error
}

func (e *providerHTTPError) Error() string { return e.err.Error() }
func (e *providerHTTPError) Unwrap() error { return e.err }
func retryableRead(err error) bool {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	var response *providerHTTPError
	if errors.As(err, &response) {
		switch response.code {
		case 408, 429, 500, 502, 503, 504:
			return true
		}
		return false
	}
	var network net.Error
	return errors.As(err, &network) || errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF)
}
