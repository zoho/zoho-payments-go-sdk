package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/zoho/zoho-payments-go-sdk/model"
)

func MarshalBody(value any) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(value); err != nil {
		return nil, fmt.Errorf("zohopayments: failed to encode request body: %w", err)
	}
	return bytes.TrimRight(buffer.Bytes(), "\n"), nil
}

func DecodeEnvelope(body []byte, result any, keys ...string) error {
	if len(bytes.TrimSpace(body)) == 0 {
		return nil
	}
	var root map[string]json.RawMessage
	if err := json.Unmarshal(body, &root); err != nil {
		return fmt.Errorf("zohopayments: failed to decode response body: %w", err)
	}
	for _, key := range keys {
		if raw, ok := root[key]; ok {
			if err := json.Unmarshal(raw, result); err != nil {
				return fmt.Errorf("zohopayments: failed to decode %q: %w", key, err)
			}
			return nil
		}
	}
	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("zohopayments: failed to decode response body: %w", err)
	}
	return nil
}

func ListEnvelope(body []byte, key string) ([]json.RawMessage, model.PageContext, error) {
	pageContext := model.PageContext{}
	if len(bytes.TrimSpace(body)) == 0 {
		return nil, pageContext, nil
	}
	var root map[string]json.RawMessage
	if err := json.Unmarshal(body, &root); err != nil {
		return nil, pageContext, fmt.Errorf("zohopayments: failed to decode response body: %w", err)
	}
	var items []json.RawMessage
	if raw, ok := root[key]; ok {
		if err := json.Unmarshal(raw, &items); err != nil {
			return nil, pageContext, fmt.Errorf("zohopayments: failed to decode %q: %w", key, err)
		}
	}
	if raw, ok := root["page_context"]; ok {
		_ = json.Unmarshal(raw, &pageContext)
	}
	return items, pageContext, nil
}

func stringField(data map[string]any, key string) string {
	value, ok := data[key]
	if !ok || value == nil {
		return ""
	}
	switch typed := value.(type) {
	case string:
		return typed
	case float64:
		if typed == float64(int64(typed)) {
			return strconv.FormatInt(int64(typed), 10)
		}
		return strconv.FormatFloat(typed, 'f', -1, 64)
	default:
		return fmt.Sprint(typed)
	}
}

func toInt64(value any) int64 {
	switch number := value.(type) {
	case float64:
		return int64(number)
	case json.Number:
		parsed, _ := number.Int64()
		return parsed
	case string:
		parsed, _ := strconv.ParseInt(number, 10, 64)
		return parsed
	default:
		return 0
	}
}
