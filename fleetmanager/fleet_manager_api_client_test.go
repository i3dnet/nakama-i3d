package fleetmanager

import (
	"reflect"
	"sort"
	"testing"
)

// sortKeyValues sorts the key values by key
func sortKeyValues(kvs []KeyValue) {
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Key < kvs[j].Key
	})
}

// TestConvertMetadata tests the convertMetadata function to ensure that the metadata is converted to the correct format
func TestConvertMetadata(t *testing.T) {
	tests := []struct {
		name     string
		input    []KeyValue
		expected map[string]any
	}{
		{
			name:     "empty slice",
			input:    []KeyValue{},
			expected: map[string]any{},
		},
		{
			name: "single key-value pair",
			input: []KeyValue{
				{Key: "players", Value: "7"},
			},
			expected: map[string]any{
				"players": "7",
			},
		},
		{
			name: "multiple key-value pairs",
			input: []KeyValue{
				{Key: "players", Value: "7"},
				{Key: "duration", Value: "200"},
				{Key: "level", Value: "expert"},
			},
			expected: map[string]any{
				"players":  "7",
				"duration": "200",
				"level":    "expert",
			},
		},
		{
			name: "duplicate keys (last value wins)",
			input: []KeyValue{
				{Key: "score", Value: "100"},
				{Key: "score", Value: "200"},
			},
			expected: map[string]any{
				"score": "200",
			},
		},
		{
			name: "special characters in keys and values",
			input: []KeyValue{
				{Key: "special@key", Value: "value!123"},
				{Key: "emoji", Value: "🎮"},
			},
			expected: map[string]any{
				"special@key": "value!123",
				"emoji":       "🎮",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertMetadata(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("convertMetadata() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestConvertFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    OriginalMetadata
		expected KeyValueMetadata
	}{
		{
			name: "Basic conversion",
			input: OriginalMetadata{
				Metadata: []map[string]any{
					{
						"players":  "7",
						"duration": "200",
					},
				},
			},
			expected: KeyValueMetadata{
				Metadata: []KeyValue{
					{Key: "players", Value: "7"},
					{Key: "duration", Value: "200"},
				},
			},
		},
		{
			name: "Multiple metadata maps",
			input: OriginalMetadata{
				Metadata: []map[string]any{
					{
						"key1": "value1",
					},
					{
						"key2": "value2",
					},
				},
			},
			expected: KeyValueMetadata{
				Metadata: []KeyValue{
					{Key: "key1", Value: "value1"},
					{Key: "key2", Value: "value2"},
				},
			},
		},
		{
			name: "Different value types",
			input: OriginalMetadata{
				Metadata: []map[string]any{
					{
						"string": "text",
						"number": 42,
						"bool":   true,
					},
				},
			},
			expected: KeyValueMetadata{
				Metadata: []KeyValue{
					{Key: "string", Value: "text"},
					{Key: "number", Value: "42"},
					{Key: "bool", Value: "true"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertFormat(tt.input)

			// Sort both expected and result slices before comparison
			sortKeyValues(result.Metadata)
			sortKeyValues(tt.expected.Metadata)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("convertFormat() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestGetSpecificMetadata tests that we can extract a particular piece of Metadata from the available metadata
func TestGetSpecificMetadata(t *testing.T) {
	tests := []struct {
		name          string
		metadata      OriginalMetadata
		key           string
		expectedValue string
		expectedFound bool
	}{
		{
			name: "existing key",
			metadata: OriginalMetadata{
				Metadata: []map[string]any{
					{
						"players":  "7",
						"duration": "200",
					},
				},
			},
			key:           "players",
			expectedValue: "7",
			expectedFound: true,
		},
		{
			name: "non-existent key",
			metadata: OriginalMetadata{
				Metadata: []map[string]any{
					{
						"players": "7",
					},
				},
			},
			key:           "notfound",
			expectedValue: "",
			expectedFound: false,
		},
		{
			name: "empty metadata",
			metadata: OriginalMetadata{
				Metadata: []map[string]any{},
			},
			key:           "players",
			expectedValue: "",
			expectedFound: false,
		},
		{
			name: "non-string value",
			metadata: OriginalMetadata{
				Metadata: []map[string]any{
					{
						"count": 42,
					},
				},
			},
			key:           "count",
			expectedValue: "42",
			expectedFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, found := getSpecificMetadata(tt.metadata, tt.key)

			if found != tt.expectedFound {
				t.Errorf("found = %v, expected %v", found, tt.expectedFound)
			}

			if value != tt.expectedValue {
				t.Errorf("value = %v, expected %v", value, tt.expectedValue)
			}
		})
	}
}
