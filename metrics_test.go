package fleetmanager

import (
	"context"
	"errors"
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/i3dnet/nakama-i3d/internal/clients"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"sync"
	"testing"
	"time"
)

type metricEvent struct {
	name string
	tags map[string]string
}
type metricRecorder struct {
	runtime.NakamaModule
	mu     sync.Mutex
	events []metricEvent
}

func (m *metricRecorder) MetricsCounterAdd(name string, tags map[string]string, delta int64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.events = append(m.events, metricEvent{name, tags})
}
func (m *metricRecorder) MetricsTimerRecord(name string, tags map[string]string, value time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.events = append(m.events, metricEvent{name, tags})
}
func (m *metricRecorder) requireResult(t *testing.T, operation, result string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	require.Len(t, m.events, 2)
	require.Equal(t, "i3d_"+operation+"_total", m.events[0].name)
	require.Equal(t, "i3d_"+operation+"_duration", m.events[1].name)
	for _, event := range m.events {
		require.Equal(t, map[string]string{"result": result}, event.tags)
	}
}
func TestAllocationMetricsRecordTerminalOutcomes(t *testing.T) {
	for _, failed := range []bool{false, true} {
		t.Run(fmt.Sprint(failed), func(t *testing.T) {
			fm, client, cache, _, _ := createFixture(t)
			metrics := &metricRecorder{}
			fm.nk = metrics
			result := "success"
			if failed {
				result = "error"
				client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("failed"))
			} else {
				client.EXPECT().AllocateApplicationInstance(gomock.Any(), gomock.Any(), gomock.Any()).Return(&runtime.InstanceInfo{Id: "private-instance"}, nil)
				cache.EXPECT().CreateGameSession(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			}
			done := make(chan createResult, 1)
			_, err := fm.Create(context.Background(), 2, []string{"private-user"}, nil, nil, resultCallback(done))
			require.NoError(t, err)
			got := awaitCreate(t, done)
			if failed {
				require.Error(t, got.err)
			} else {
				require.NoError(t, got.err)
			}
			metrics.requireResult(t, "allocation", result)
		})
	}
}
func TestReconciliationMetricsRecordOutcomes(t *testing.T) {
	for _, failed := range []bool{false, true} {
		t.Run(fmt.Sprint(failed), func(t *testing.T) {
			fm, _, client := sessionFixture(t, 0, 2)
			metrics := &metricRecorder{}
			fm.nk = metrics
			result := "success"
			if failed {
				result = "error"
				client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").Return(nil, errors.New("outage"))
			} else {
				client.EXPECT().ListApplicationInstances(gomock.Any(), gomock.Any(), 100, "").Return(&clients.ApplicationInstanceListResponse{}, nil)
			}
			_, _ = fm.reconcileOnce(context.Background(), nil, time.Now())
			metrics.requireResult(t, "reconciliation", result)
		})
	}
}
