package fleetmanager

import (
	"encoding/json"
	"fmt"
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

	var result T
	if err := json.Unmarshal([]byte(payload), result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal updateInstanceInfo request: %w", err)
	}

	return &result, nil
}
