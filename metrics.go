package fleetmanager

import (
	"context"
	"errors"
	"time"
)

// Only fixed operations and outcomes reach metrics; identifiers belong in logs.
func (fm *I3dFleetManager) recordOperation(operation string, start time.Time, err error) {
	if fm.nk == nil {
		return
	}
	result := "success"
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		result = "timeout"
	case errors.Is(err, context.Canceled):
		result = "canceled"
	case err != nil:
		result = "error"
	}
	tags := map[string]string{"result": result}
	fm.nk.MetricsCounterAdd("i3d_"+operation+"_total", tags, 1)
	fm.nk.MetricsTimerRecord("i3d_"+operation+"_duration", tags, time.Since(start))
}
