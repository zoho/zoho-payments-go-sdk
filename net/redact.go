package net

import (
	"fmt"
	"strings"

	"github.com/zoho/zoho-payments-go-sdk/redact"
)

func (r Request) String() string   { return redactRequest(r) }
func (r Request) GoString() string { return redactRequest(r) }

func (b RequestBuilder) String() string {
	return fmt.Sprintf("RequestBuilder{Method: %s, URL: %s, Headers: %v, Body: %s}",
		b.method, b.url, maskAuth(b.headers), b.body)
}
func (b RequestBuilder) GoString() string { return b.String() }

func redactRequest(request Request) string {
	return fmt.Sprintf("Request{Method: %s, URL: %s, Headers: %v, Body: %s}",
		request.method, request.url, maskAuth(request.headers), request.body)
}

func maskAuth(headers map[string][]string) map[string][]string {
	result := make(map[string][]string, len(headers))
	for name, values := range headers {
		if strings.EqualFold(name, "Authorization") {
			result[name] = []string{redact.Mask}
		} else {
			result[name] = values
		}
	}
	return result
}
