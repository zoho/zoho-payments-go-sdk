package internal

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/zoho/zoho-payments-go-sdk/model"
	znet "github.com/zoho/zoho-payments-go-sdk/net"
)

func DoList[T any](caller *Caller, path string, query url.Values, key string) (*model.ListResponse[T], error) {
	response, err := caller.Do(znet.GET, path, query, nil)
	if err != nil {
		return nil, err
	}
	rawItems, pageContext, err := ListEnvelope([]byte(response.Body()), key)
	if err != nil {
		return nil, err
	}
	data := make([]T, 0, len(rawItems))
	for _, rawItem := range rawItems {
		var item T
		if err := json.Unmarshal(rawItem, &item); err != nil {
			return nil, fmt.Errorf("zohopayments: failed to decode list item: %w", err)
		}
		data = append(data, item)
	}
	return &model.ListResponse[T]{Data: data, PageContext: pageContext}, nil
}
