package param

import (
	"net/url"
	"strconv"
)

func putStr(body map[string]any, key string, value *string) {
	if value != nil {
		body[key] = *value
	}
}

func putInt(body map[string]any, key string, value *int) {
	if value != nil {
		body[key] = *value
	}
}

func putFloat(body map[string]any, key string, value *float64) {
	if value != nil {
		body[key] = *value
	}
}

func putBool(body map[string]any, key string, value *bool) {
	if value != nil {
		body[key] = *value
	}
}

func metaDataToList(metaData []MetaDataParams) []map[string]any {
	if len(metaData) == 0 {
		return nil
	}
	result := make([]map[string]any, 0, len(metaData))
	for _, entry := range metaData {
		result = append(result, entry.toMap())
	}
	return result
}

func qSetStr(query url.Values, key string, value *string) {
	if value != nil {
		query.Set(key, *value)
	}
}

func qSetInt(query url.Values, key string, value *int) {
	if value != nil {
		query.Set(key, strconv.Itoa(*value))
	}
}
