//go:build !i3d_smoke

package main

import (
	"context"
	"github.com/heroiclabs/nakama-common/runtime"
)

func registerSmokeChecks(runtime.Initializer) error                         { return nil }
func smokeBeforeNotify(context.Context, runtime.NakamaModule, string) error { return nil }
