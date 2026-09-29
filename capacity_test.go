package fleetmanager

import (
	"encoding/json"
	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestCapacitySurvivesJSONStorage(t *testing.T) {
	data, err := json.Marshal(&runtime.InstanceInfo{Metadata: map[string]any{MaxPlayers: 8}})
	require.NoError(t, err)
	var loaded runtime.InstanceInfo
	require.NoError(t, json.Unmarshal(data, &loaded))
	capacity, err := getMaxPlayers(&loaded)
	require.NoError(t, err)
	require.Equal(t, 8, capacity)
}
func TestCapacityRejectsInvalidValues(t *testing.T) {
	for _, value := range []any{0, -1, 1.5, math.NaN(), math.Inf(1), float64(math.MaxInt), json.Number("2.5"), "8", nil} {
		_, err := getMaxPlayers(&runtime.InstanceInfo{Metadata: map[string]any{MaxPlayers: value}})
		require.Error(t, err, "value %v", value)
	}
}
