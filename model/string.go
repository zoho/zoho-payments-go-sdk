package model

import (
	"bytes"
	"encoding/json"
)

type String string

func (s *String) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		return nil
	}
	if data[0] == '{' || data[0] == '[' {
		return nil
	}
	if data[0] == '"' {
		var unquoted string
		if err := json.Unmarshal(data, &unquoted); err != nil {
			return err
		}
		*s = String(unquoted)
		return nil
	}
	*s = String(data)
	return nil
}

func (s String) String() string { return string(s) }
