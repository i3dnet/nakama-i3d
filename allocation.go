package fleetmanager

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

type allocationOutcome struct {
	instance *runtime.InstanceInfo
	err      error
}

// allocationResult serializes completion against abandonment. A result already
// published wins over a concurrently expired context; an abandoned result remains
// owned by its worker, which must reclaim any known allocation.
type allocationResult struct {
	mu               sync.Mutex
	done             chan struct{}
	ready, abandoned bool
	outcome          allocationOutcome
}

func (r *allocationResult) complete(outcome allocationOutcome) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.abandoned {
		return false
	}
	r.outcome, r.ready = outcome, true
	close(r.done)
	return true
}

func (r *allocationResult) wait(ctx context.Context) allocationOutcome {
	select {
	case <-r.done:
	case <-ctx.Done():
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.ready {
		return r.outcome
	}
	r.abandoned = true
	return allocationOutcome{err: ctx.Err()}
}

func (fm *I3dFleetManager) runAllocationStage(parent context.Context, timeout time.Duration, work func(context.Context) allocationOutcome) allocationOutcome {
	ctx, cancel := context.WithTimeout(context.WithoutCancel(parent), timeout)
	defer cancel()
	stopShutdown := context.AfterFunc(fm.ctx, cancel)
	defer stopShutdown()
	if fm.ctx.Err() != nil {
		cancel()
	}
	result := &allocationResult{done: make(chan struct{})}
	// The coordinator is still tracked, so shutdown cannot finish before this Add.
	fm.operations.Add(1)
	go func() {
		defer fm.operations.Done()
		outcome := work(ctx)
		if !result.complete(outcome) {
			fm.reclaimAllocation(parent, outcome.instance)
		}
	}()
	return result.wait(ctx)
}

func (fm *I3dFleetManager) allocationFinalizeTimeout() time.Duration {
	if fm.cfg.AllocationFinalizeTimeout > 0 {
		return fm.cfg.AllocationFinalizeTimeout
	}
	return 30 * time.Second
}

// Cleanup uses a fresh bound even after allocation/shutdown cancellation. Never
// retry an ambiguous restart, and never guess an ID from an ambiguous allocation.
// A storage write may have committed before failing; reconciliation conditionally
// removes stale cache state rather than deleting a potentially newer generation.
func (fm *I3dFleetManager) reclaimAllocation(parent context.Context, instance *runtime.InstanceInfo) {
	if instance == nil || strings.TrimSpace(instance.Id) == "" {
		return
	}
	ctx, cancel := context.WithTimeout(context.WithoutCancel(parent), fm.allocationFinalizeTimeout())
	defer cancel()
	if err := fm.client.RestartApplicationInstance(ctx, instance.Id); err != nil {
		fm.logger.WithField("instance_id", instance.Id).Error("failed to reclaim allocation: %v", err)
	}
}
