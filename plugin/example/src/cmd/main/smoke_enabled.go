//go:build i3d_smoke

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	"sync/atomic"
	"time"
)

const smokeCollection = "_i3D_instances"

func registerSmokeChecks(initializer runtime.Initializer) error {
	return initializer.RegisterRpc("i3d_smoke", smokeRPC)
}
func smokeBeforeNotify(ctx context.Context, nk runtime.NakamaModule, id string) error {
	objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{{Collection: smokeCollection, Key: id}})
	if err != nil {
		return err
	}
	if len(objects) != 1 {
		return fmt.Errorf("callback preceded storage")
	}
	var instance runtime.InstanceInfo
	if err := json.Unmarshal([]byte(objects[0].Value), &instance); err != nil {
		return err
	}
	if instance.Id != id || instance.Metadata["i3d_max_players"] != float64(2) {
		return fmt.Errorf("stored allocation is incomplete")
	}
	return nil
}
func smokeRPC(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	if ctx.Value(runtime.RUNTIME_CTX_USER_ID) != nil {
		return "", runtime.NewError("server only", 7)
	}
	var request struct {
		Op string `json:"op"`
		ID string `json:"id"`
	}
	if err := json.Unmarshal([]byte(payload), &request); err != nil {
		return "", err
	}
	switch request.Op {
	case "read":
		objects, err := nk.StorageRead(ctx, []*runtime.StorageRead{{Collection: smokeCollection, Key: request.ID}})
		if err != nil {
			return "", err
		}
		response := map[string]any{"exists": len(objects) > 0}
		if len(objects) > 0 {
			response["instance"] = json.RawMessage(objects[0].Value)
		}
		data, err := json.Marshal(response)
		return string(data), err
	case "create":
		result := make(chan map[string]any, 1)
		_, err := nk.GetFleetManager().Create(ctx, 2, nil, nil, map[string]any{"mode": "recovery"}, func(status runtime.FmCreateStatus, instance *runtime.InstanceInfo, _ []*runtime.SessionInfo, _ map[string]any, err error) {
			response := map[string]any{"success": err == nil && status == runtime.CreateSuccess}
			if instance != nil {
				response["id"] = instance.Id
			}
			result <- response
		})
		if err != nil {
			return "", err
		}
		select {
		case response := <-result:
			data, err := json.Marshal(response)
			return string(data), err
		case <-ctx.Done():
			return "", ctx.Err()
		}
	case "race":
		return smokeStorageRace(ctx, nk)
	default:
		return "", runtime.NewError("unknown smoke operation", 3)
	}
}
func smokeStorageRace(ctx context.Context, nk runtime.NakamaModule) (string, error) {
	id := "smoke-" + uuid.NewString()
	_, err := nk.StorageWrite(ctx, []*runtime.StorageWrite{{Collection: smokeCollection, Key: id, Value: `{"count":0}`, Version: "*"}})
	if err != nil {
		return "", err
	}
	defer nk.StorageDelete(context.WithoutCancel(ctx), []*runtime.StorageDelete{{Collection: smokeCollection, Key: id}})
	reads := []*runtime.StorageRead{{Collection: smokeCollection, Key: id}}
	barrier := make(chan struct{})
	results := make(chan error, 2)
	var mutations atomic.Int32
	for i := 0; i < 2; i++ {
		go func() {
			_, err := nk.StorageWriteRetry(ctx, reads, func(objects []*api.StorageObject) ([]*runtime.StorageWrite, error) {
				n := mutations.Add(1)
				if n == 2 {
					close(barrier)
				}
				if n <= 2 {
					select {
					case <-barrier:
					case <-ctx.Done():
						return nil, ctx.Err()
					}
				}
				var value struct {
					Count int `json:"count"`
				}
				if len(objects) != 1 {
					return nil, fmt.Errorf("missing test record")
				}
				if err := json.Unmarshal([]byte(objects[0].Value), &value); err != nil {
					return nil, err
				}
				value.Count++
				data, _ := json.Marshal(value)
				return []*runtime.StorageWrite{{Collection: smokeCollection, Key: id, Value: string(data), Version: objects[0].Version}}, nil
			}, 5)
			results <- err
		}()
	}
	for i := 0; i < 2; i++ {
		if err := <-results; err != nil {
			return "", err
		}
	}
	objects, err := nk.StorageRead(ctx, reads)
	if err != nil {
		return "", err
	}
	var result struct {
		Count int `json:"count"`
	}
	if len(objects) == 1 {
		_ = json.Unmarshal([]byte(objects[0].Value), &result)
	}
	if len(objects) != 1 || result.Count != 2 || mutations.Load() < 3 {
		return "", fmt.Errorf("native retry did not resolve conflict")
	}
	// Exercise native conditional-delete behavior and ensure its failed delete is safe.
	if err := nk.StorageDelete(ctx, []*runtime.StorageDelete{{Collection: smokeCollection, Key: id, Version: "stale-version"}}); err == nil {
		return "", fmt.Errorf("stale delete succeeded")
	}
	objects, err = nk.StorageRead(ctx, reads)
	if err != nil || len(objects) != 1 {
		return "", fmt.Errorf("stale delete removed record")
	}
	value, _ := json.Marshal(&runtime.InstanceInfo{Id: id, Status: "ALLOCATED", Metadata: map[string]any{"i3d_max_players": 1}})
	_, err = nk.StorageWrite(ctx, []*runtime.StorageWrite{{Collection: smokeCollection, Key: id, Value: string(value), Version: objects[0].Version}})
	if err != nil {
		return "", err
	}
	admissions := make(chan int, 2)
	for _, user := range []string{"first", "second"} {
		go func(user string) {
			joined, err := nk.GetFleetManager().Join(ctx, id, []string{user}, nil)
			if err != nil {
				admissions <- -1
				return
			}
			admissions <- len(joined.SessionInfo)
		}(user)
	}
	a, b := <-admissions, <-admissions
	if a < 0 || b < 0 || a+b != 1 {
		return "", fmt.Errorf("concurrent Join admissions: %d + %d", a, b)
	}
	// Verify the actual runtime's indexed fields and cursor ordering.
	for _, count := range []int{2, 0, 1} {
		key := fmt.Sprintf("%s-%d", id, count)
		value, _ := json.Marshal(&runtime.InstanceInfo{Id: key, Status: "ALLOCATED", CreateTime: time.Now().UTC(), PlayerCount: count, Metadata: map[string]any{"smoke": "sorting"}})
		_, err := nk.StorageWrite(ctx, []*runtime.StorageWrite{{Collection: smokeCollection, Key: key, Value: string(value), Version: "*"}})
		if err != nil {
			return "", err
		}
		defer nk.StorageDelete(context.WithoutCancel(ctx), []*runtime.StorageDelete{{Collection: smokeCollection, Key: key}})
	}
	cursor := ""
	for count := 0; count < 3; count++ {
		instances, next, err := nk.GetFleetManager().List(ctx, "+value.metadata.smoke:sorting", 1, cursor)
		if err != nil {
			return "", err
		}
		if len(instances) != 1 || instances[0].PlayerCount != count {
			return "", fmt.Errorf("storage index sort/page %d: %v", count, instances)
		}
		cursor = next
	}
	if cursor != "" {
		instances, next, err := nk.GetFleetManager().List(ctx, "+value.metadata.smoke:sorting", 1, cursor)
		if err != nil || len(instances) != 0 || next != "" {
			return "", fmt.Errorf("storage cursor did not finish")
		}
	}

	return `{"success":true}`, nil
}
