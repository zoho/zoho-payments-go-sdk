package model

import (
	"bytes"
	"encoding/json"
)

// coerceStrings rewrites a JSON object so that the named keys, where they hold a
// bare number or boolean, are quoted into JSON strings.
//
// It exists because the API is not consistent about the form of its date/time
// values: /payouts returns initiated_time as "1789631992831" while /payouts/{id}
// returns it as 1789631992. Go's encoding/json rejects a number into a string
// field, so the value is normalised before decoding. This mirrors the PHP SDK's
// Coerce::optStr, which casts any scalar to a string, and the lenient Long and
// Integer adapters the Java SDK registers with Gson.
//
// Values are quoted from their literal source bytes, so nothing is reformatted.
// Objects, arrays, nulls and existing strings are left untouched.
func coerceStrings(data []byte, keys ...string) ([]byte, error) {
	var object map[string]json.RawMessage
	if err := json.Unmarshal(data, &object); err != nil || object == nil {
		// Leave malformed or null input to the caller's decode, which reports it.
		return data, nil
	}
	patched := false
	for _, key := range keys {
		raw, ok := object[key]
		if !ok {
			continue
		}
		value := bytes.TrimSpace(raw)
		if len(value) == 0 || value[0] == '"' || value[0] == '{' || value[0] == '[' ||
			bytes.Equal(value, []byte("null")) {
			continue
		}
		quoted, err := json.Marshal(string(value))
		if err != nil {
			return nil, err
		}
		object[key] = quoted
		patched = true
	}
	if !patched {
		return data, nil
	}
	return json.Marshal(object)
}
