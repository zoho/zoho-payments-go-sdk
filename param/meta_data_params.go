package param

// MetaDataParams is a custom metadata key-value pair attached to a request.
type MetaDataParams struct {
	Key   string
	Value string
}

func (m MetaDataParams) toMap() map[string]any {
	return map[string]any{"key": m.Key, "value": m.Value}
}
