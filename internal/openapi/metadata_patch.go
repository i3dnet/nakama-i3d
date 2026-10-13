package openapi

// MetadataPatchCollection models the writable subset used by the adapter for
// allocation and instance updates. The response models include read-only fields
// and cannot represent explicit null metadata values. Keep this request model in
// sync with MetadataPatchCollection/MetadataChange in api/openapi.yaml.
type MetadataPatchCollection struct {
	Metadata []MetadataChange `json:"metadata"`
}

// MetadataChange merges a key into existing metadata. A nil Value explicitly
// deletes the key; an empty string is a value and must not be omitted.
type MetadataChange struct {
	Key   string  `json:"key"`
	Value *string `json:"value"`
}
