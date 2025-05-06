package clients

import (
	"errors"
	"github.com/stretchr/testify/suite"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/config"
	"gitlab.com/i3Dnet/dev/game/projects/plugins/nakama/fleetmanager/internal/tests"
	"testing"
	"time"
)

type RetryExecutorTestSuite struct {
	suite.Suite
	logger *tests.MockLogger
	cfg    *config.Config
}

func (s *RetryExecutorTestSuite) SetupTest() {
	s.logger = tests.NewMockLogger()
	s.cfg = &config.Config{
		Retry: config.Retry{
			Delay:    5 * time.Millisecond,
			MaxDelay: 20 * time.Millisecond,
		},
	}
}

func (s *RetryExecutorTestSuite) TestSuccessOnFirstTry() {
	executor := NewRetryExecutor(s.logger, s.cfg, func(err error) bool { return true })

	var callCount int
	err := executor.Run(3, func() error {
		callCount++
		return nil
	})

	s.NoError(err)
	s.Equal(1, callCount)
}

func (s *RetryExecutorTestSuite) TestSuccessAfterRetry() {
	executor := NewRetryExecutor(s.logger, s.cfg, func(err error) bool { return true })

	var callCount int
	err := executor.Run(3, func() error {
		callCount++
		if callCount < 2 {
			return errors.New("temporary failure")
		}
		return nil
	})

	s.NoError(err)
	s.Equal(2, callCount)
}

func (s *RetryExecutorTestSuite) TestFailsAfterAllAttempts() {
	executor := NewRetryExecutor(s.logger, s.cfg, func(err error) bool { return true })

	err := executor.Run(3, func() error {
		return errors.New("always fails")
	})

	s.Error(err)
	s.EqualError(err, "always fails")
}

func (s *RetryExecutorTestSuite) TestShouldRetryFalse() {
	executor := NewRetryExecutor(s.logger, s.cfg, func(err error) bool { return false })

	var callCount int
	err := executor.Run(3, func() error {
		callCount++
		return errors.New("non-retriable")
	})

	s.Equal(1, callCount)
	s.EqualError(err, "non-retriable")
}

func TestRetryExecutorTestSuite(t *testing.T) {
	suite.Run(t, new(RetryExecutorTestSuite))
}
