package fleetmanager

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// UpdateInstanceInfoRequest is the data transfer object for the UpdateInstanceInfo RPC
type UpdateInstanceInfoRequest struct {
	Id          string         `json:"id"`
	PlayerCount int            `json:"player_count"`
	Metadata    map[string]any `json:"metadata"`
}

type DeleteInstanceInfoRequest struct {
	Id string `json:"id"`
}

func FromPayloadToRequest[T any](payload string) (*T, error) {

	result := new(T)
	if err := json.Unmarshal([]byte(payload), result); err != nil {
		typeName := reflect.TypeOf(result).Elem().Name()
		if typeName == "" {
			typeName = reflect.TypeOf(result).Name()
		}

		return nil, fmt.Errorf("failed to unmarshal %s request: %w", typeName, err)
	}

	return result, nil
}
