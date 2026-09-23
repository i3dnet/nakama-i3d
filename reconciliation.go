package fleetmanager

import (
	"context"
	"errors"
	"fmt"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/internal/clients"
	"github.com/i3dnet/nakama-i3d/internal/storage"
	"time"
)

func (fm *I3dFleetManager) providerScope() string {
	return NewFilterBuilder().Add(applicationId, fm.cfg.ApplicationId).Add(FleetId, fm.cfg.FleetId).Query()
}
func (fm *I3dFleetManager) runReconciliation() {
	defer fm.operations.Done()
	ticker := time.NewTicker(fm.cfg.ReconcileInterval)
	defer ticker.Stop()
	var missing map[string]string
	for {
		if fm.ctx.Err() != nil {
			return
		}
		timeout := fm.cfg.ReconcileTimeout
		if timeout <= 0 {
			timeout = 30 * time.Second
		}
		ctx, cancel := context.WithTimeout(fm.ctx, timeout)
		next, err := fm.reconcileOnce(ctx, missing, time.Now())
		cancel()
		if err != nil {
			missing = nil
			if fm.ctx.Err() == nil {
				fm.logger.WithField("error", err.Error()).Warn("fleet reconciliation failed")
			}
		} else {
			missing = next
		}
		select {
		case <-fm.ctx.Done():
			return
		case <-ticker.C:
		}
	}
}
func (fm *I3dFleetManager) reconcileOnce(ctx context.Context, previous map[string]string, now time.Time) (missing map[string]string, err error) {
	start := time.Now()
	defer func() { fm.recordOperation("reconciliation", start, err) }()
	return fm.reconcileSnapshot(ctx, previous, now)
}
func (fm *I3dFleetManager) reconcileSnapshot(ctx context.Context, previous map[string]string, now time.Time) (map[string]string, error) {
	// Captured first: new allocations and concurrent writes cannot be overwritten.
	snapshot, err := fm.storage.SnapshotGameSessions(ctx)
	if err != nil {
		return nil, err
	}
	provider := map[string]*runtime.InstanceInfo{}
	cursor := ""
	seen := map[string]bool{}
	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		page, err := fm.client.ListApplicationInstances(ctx, fm.providerScope(), 100, cursor)
		if err != nil {
			return nil, err
		}
		if page == nil {
			return nil, fmt.Errorf("provider returned a nil page")
		}
		for _, instance := range page.Instances {
			if instance == nil || instance.Id == "" {
				return nil, fmt.Errorf("provider returned an invalid instance")
			}
			if _, duplicate := provider[instance.Id]; duplicate {
				return nil, fmt.Errorf("provider repeated instance across pages")
			}
			provider[instance.Id] = instance
		}
		if page.NextCursor == "" {
			break
		}
		if seen[page.NextCursor] {
			return nil, fmt.Errorf("provider cursor repeated")
		}
		seen[page.NextCursor] = true
		cursor = page.NextCursor
	}
	byID := map[string]*api.StorageObject{}
	missing := map[string]string{}
	createdDuringScan := map[string]bool{}
	skew := fm.cfg.ReconcileClockSkew
	if skew <= 0 {
		skew = 5 * time.Second
	}
	cutoff := now.Add(-skew)
	for _, obj := range snapshot {
		byID[obj.Key] = obj
		scoped, allocatedAt, err := storage.SnapshotInScope(obj, fm.cfg.ApplicationId, fm.cfg.FleetId)
		if err != nil {
			return nil, err
		}
		// Storage pagination is not a transaction-wide snapshot. An allocation
		// may be written after this pass starts but before its page is read.
		// Include the configured bound on inter-node clock skew and precision.
		if (!allocatedAt.IsZero() && !allocatedAt.Before(cutoff)) ||
			(obj.UpdateTime != nil && !obj.UpdateTime.AsTime().Before(cutoff)) {
			createdDuringScan[obj.Key] = true
			continue
		}
		incoming := provider[obj.Key]
		if incoming != nil && incoming.Status == clients.ApplicationInstanceStatus[5] {
			continue
		}
		if !scoped || allocatedAt.IsZero() || now.Sub(allocatedAt) < fm.cfg.ReconcileGracePeriod {
			continue
		}
		missing[obj.Key] = obj.Version
		if previous[obj.Key] != obj.Version {
			continue
		}
		if err := fm.storage.ReconcileGameSession(ctx, obj, nil, "", ""); err != nil && !errors.Is(err, runtime.ErrStorageRejectedVersion) {
			return nil, err
		}
	}
	for id, incoming := range provider {
		if createdDuringScan[id] {
			continue
		}
		if incoming.Status != clients.ApplicationInstanceStatus[5] {
			continue
		}
		if err := fm.storage.ReconcileGameSession(ctx, byID[id], incoming, fm.cfg.ApplicationId, fm.cfg.FleetId); err != nil && !errors.Is(err, runtime.ErrStorageRejectedVersion) {
			return nil, err
		}
	}
	return missing, nil
}
