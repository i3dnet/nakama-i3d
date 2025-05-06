package fleetmanager

import (
	"fmt"
	"github.com/heroiclabs/nakama-common/runtime"
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

	if p, ok := instance.Metadata[MaxPlayers]; ok {
		if result, ok := p.(int); ok {
			return result, nil
		}
	}

	return 0, fmt.Errorf("maxPlayers not found in instance")
}
