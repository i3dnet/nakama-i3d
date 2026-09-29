package fleetmanager

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
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

	if strings.TrimSpace(payload) == "null" {
		return nil, fmt.Errorf("request must be an object")
	}
	result := new(T)
	decoder := json.NewDecoder(strings.NewReader(payload))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(result); err != nil {
		typeName := reflect.TypeOf(result).Elem().Name()
		if typeName == "" {
			typeName = reflect.TypeOf(result).Name()
		}

		return nil, fmt.Errorf("failed to unmarshal %s request: %w", typeName, err)
	}

	if err := decoder.Decode(new(any)); err != io.EOF {
		return nil, fmt.Errorf("request must contain exactly one JSON object")
	}
	return result, nil
}
