package main

import (
	"context"
	"github.com/heroiclabs/nakama-common/runtime"
	"testing"
	"time"
)

type testLogger struct{ runtime.Logger }

func (testLogger) Debug(string, ...interface{}) {}
func (testLogger) Error(string, ...interface{}) {}

type testEntry struct{ runtime.MatchmakerEntry }

func (testEntry) GetProperties() map[string]interface{} { return map[string]interface{}{} }
func (testEntry) GetPresence() runtime.Presence         { return testPresence{} }

type testPresence struct{ runtime.Presence }

func (testPresence) GetUserId() string { return "player" }

type testFleet struct {
	runtime.FleetManager
	callback runtime.FmCreateCallbackFn
}

func (f *testFleet) Create(_ context.Context, _ int, _ []string, _ []runtime.FleetUserLatencies, _ map[string]any, cb runtime.FmCreateCallbackFn) (map[string]string, error) {
	f.callback = cb
	return nil, nil
}

type testNakama struct {
	runtime.NakamaModule
	fleet *testFleet
	send  func(context.Context, []*runtime.NotificationSend) error
}

func (n *testNakama) GetFleetManager() runtime.FleetManager { return n.fleet }
func (n *testNakama) NotificationsSend(ctx context.Context, notifications []*runtime.NotificationSend) error {
	return n.send(ctx, notifications)
}
func TestNotificationSurvivesMatchmakerHookCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sent := false
	nk := &testNakama{fleet: &testFleet{}, send: func(ctx context.Context, n []*runtime.NotificationSend) error {
		sent = true
		if ctx.Err() != nil {
			t.Errorf("notification context canceled: %v", ctx.Err())
		}
		deadline, ok := ctx.Deadline()
		if !ok || time.Until(deadline) > 11*time.Second {
			t.Error("notification must have a bounded deadline")
		}
		if len(n) != 1 || n[0].UserID != "player" || n[0].Content["Port"] != 7777 {
			t.Errorf("unexpected notification: %+v", n)
		}
		return nil
	}}
	if _, err := MatchmakerMatched(ctx, testLogger{}, nil, nk, []runtime.MatchmakerEntry{testEntry{}}); err != nil {
		t.Fatal(err)
	}
	cancel()
	nk.fleet.callback(runtime.CreateSuccess, &runtime.InstanceInfo{ConnectionInfo: &runtime.ConnectionInfo{IpAddress: "127.0.0.1", Port: 7777}}, []*runtime.SessionInfo{{UserId: "player"}}, nil, nil)
	if !sent {
		t.Error("notification not sent")
	}
}

func TestInvalidRuntimeConfigurationDoesNotFallBackToProcessCredentials(t *testing.T) {
	t.Setenv("PROJECT_ROOT", t.TempDir())
	t.Setenv("I3D_APPLICATION_ID", "process-app")
	t.Setenv("I3D_ACCESS_TOKEN", "test-only")
	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_ENV, map[string]string{"I3D_USE_BEARER_AUTH": "invalid"})
	// Valid process credentials must not cause construction/registration after
	// explicit runtime input fails validation.
	err := InitModule(ctx, testLogger{}, nil, nil, nil)
	runtimeErr, ok := err.(*runtime.Error)
	if !ok || runtimeErr.Code != 3 {
		t.Fatalf("expected invalid runtime configuration, got %v", err)
	}
}

func TestMalformedRuntimeEnvironmentDoesNotFallBack(t *testing.T) {
	t.Setenv("PROJECT_ROOT", t.TempDir())
	t.Setenv("I3D_APPLICATION_ID", "process-app")
	t.Setenv("I3D_ACCESS_TOKEN", "test-only")
	ctx := context.WithValue(context.Background(), runtime.RUNTIME_CTX_ENV, "malformed")
	err := InitModule(ctx, testLogger{}, nil, nil, nil)
	if runtimeErr, ok := err.(*runtime.Error); !ok || runtimeErr.Code != 3 {
		t.Fatalf("expected invalid runtime environment, got %v", err)
	}
}
