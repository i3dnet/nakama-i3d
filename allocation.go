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

// runAllocationStage abandons work after timeout. work may run workGrace longer:
// a non-idempotent call canceled at abandonment can succeed remotely without
// returning the confirmed result its worker needs to reclaim. Shutdown cancels both.
func (fm *I3dFleetManager) runAllocationStage(parent context.Context, timeout, workGrace time.Duration, work func(context.Context) allocationOutcome) allocationOutcome {
	workCtx, cancelWork := context.WithTimeout(context.WithoutCancel(parent), timeout+workGrace)
	stopShutdown := context.AfterFunc(fm.ctx, cancelWork)
	if fm.ctx.Err() != nil {
		cancelWork()
	}
	ctx, cancel := context.WithTimeout(workCtx, timeout)
	defer cancel()
	result := &allocationResult{done: make(chan struct{})}
	// The coordinator is still tracked, so shutdown cannot finish before this Add.
	fm.operations.Add(1)
	go func() {
		defer fm.operations.Done()
		// Canceling workCtx also cancels ctx, so it must follow result.complete.
		defer cancelWork()
		defer stopShutdown()
		outcome := work(workCtx)
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
