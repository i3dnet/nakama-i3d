package fleetmanager

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
	"unsafe"
)

type mutableMetadataKey struct{ Value string }

func (k *mutableMetadataKey) MarshalText() ([]byte, error) { return []byte(k.Value), nil }

func TestMetadataCopyOwnsMutableTextMarshalerMapKeys(t *testing.T) {
	key := &mutableMetadataKey{Value: "original"}
	copied, err := cloneMetadata(map[string]any{"keyed": map[*mutableMetadataKey]string{key: "value"}})
	require.NoError(t, err)
	key.Value = "changed"
	keys := copied["keyed"].(map[*mutableMetadataKey]string)
	require.Len(t, keys, 1)
	for k := range keys {
		require.Equal(t, "original", k.Value)
	}
}

type privateMetadataFields struct{ Values []int }
type promotedMetadata struct{ privateMetadataFields }
type privateMutableMetadata struct{ values []int }

func (v privateMutableMetadata) MarshalJSON() ([]byte, error) { return json.Marshal(v.values) }

func TestCreateRejectsMetadataWithUncloneablePrivateState(t *testing.T) {
	for _, value := range []any{
		promotedMetadata{privateMetadataFields{Values: []int{1}}},
		privateMutableMetadata{values: []int{1}},
	} {
		fm, _, _, registry, _ := createFixture(t)
		_, err := fm.Create(context.Background(), 2, nil, nil, map[string]any{"value": value}, resultCallback(make(chan createResult, 1)))
		require.ErrorIs(t, err, ErrInvalidInput)
		registry.mu.Lock()
		require.Empty(t, registry.callbacks)
		registry.mu.Unlock()
	}
}

type namedMetadataPointer *int

func TestMetadataCopyPreservesNamedPointerType(t *testing.T) {
	value := 7
	copied, err := cloneMetadata(map[string]any{"pointer": namedMetadataPointer(&value)})
	require.NoError(t, err)
	require.IsType(t, namedMetadataPointer(nil), copied["pointer"])
	value = 9
	require.Equal(t, 7, *copied["pointer"].(namedMetadataPointer))
}

type channelMetadata chan int

func (channelMetadata) MarshalJSON() ([]byte, error) { return []byte("[]"), nil }

type functionMetadata func() int

func (functionMetadata) MarshalJSON() ([]byte, error) { return []byte("1"), nil }

type opaquePointerMetadata struct{ Pointer unsafe.Pointer }

func (opaquePointerMetadata) MarshalJSON() ([]byte, error) { return []byte("null"), nil }

func TestMetadataRejectsUncopyableReferencesEvenWithCustomJSON(t *testing.T) {
	value := 7
	for _, metadata := range []any{channelMetadata(make(chan int, 1)), functionMetadata(func() int { return value }), opaquePointerMetadata{Pointer: unsafe.Pointer(&value)}} {
		_, err := cloneMetadata(map[string]any{"value": metadata})
		require.Error(t, err)
	}
}
