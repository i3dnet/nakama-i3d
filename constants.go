package fleetmanager

import (
	"encoding/json"
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
	"math"
	"strconv"
)

// Error codes for the Fleet Manager
var (
	ErrInvalidInput  = runtime.NewError("input is invalid", 3)       // INVALID_ARGUMENT
	ErrInternalError = runtime.NewError("internal server error", 13) // INTERNAL
)

// RPC IDs for the Fleet Manager
const (
	RpcIdUpdateInstanceInfo = "update_instance_info"
	RpcIdDeleteInstanceInfo = "delete_instance_info"
)

const MaxPlayers = "i3d_max_players"
const AllocationStatus = "5"

func getMaxPlayers(instance *runtime.InstanceInfo) (int, error) {
	if instance == nil {
		return 0, fmt.Errorf("maxPlayers not found in instance")
	}
	value := instance.Metadata[MaxPlayers]
	var capacity int
	switch number := value.(type) {
	case int:
		capacity = number
	case float64:
		if math.IsNaN(number) || math.IsInf(number, 0) || math.Trunc(number) != number || number >= math.Ldexp(1, strconv.IntSize-1) {
			return 0, fmt.Errorf("invalid maxPlayers")
		}
		capacity = int(number)
	case json.Number:
		n, err := number.Int64()
		if err != nil || int64(int(n)) != n {
			return 0, fmt.Errorf("invalid maxPlayers")
		}
		capacity = int(n)
	default:
		return 0, fmt.Errorf("maxPlayers not found in instance")
	}
	if capacity <= 0 {
		return 0, fmt.Errorf("maxPlayers must be positive")
	}
	return capacity, nil
}
