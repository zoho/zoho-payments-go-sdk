package internal

import (
	"encoding/json"
	"net/url"
	"strings"
	"time"

	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/exception"
	znet "github.com/zoho/zoho-payments-go-sdk/net"
)

// Caller is the authenticated request dispatcher shared by all services.
type Caller struct {
	transport      znet.HTTPClient
	tokens         *TokenManager
	edition        edition.Edition
	accountID      string
	requestTimeout time.Duration
	defaultHeaders map[string]string
}

func NewCaller(transport znet.HTTPClient, tokens *TokenManager, ed edition.Edition,
	accountID string, requestTimeout time.Duration, defaultHeaders map[string]string) *Caller {
	return &Caller{
		transport:      transport,
		tokens:         tokens,
		edition:        ed,
		accountID:      accountID,
		requestTimeout: requestTimeout,
		defaultHeaders: defaultHeaders,
	}
}

func (c *Caller) buildURL(path string, query url.Values) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	values := url.Values{}
	for name, valueList := range query {
		for _, value := range valueList {
			values.Add(name, value)
		}
	}
	values.Set("account_id", c.accountID)
	return c.edition.BaseURL() + path + "?" + values.Encode()
}

func (c *Caller) Do(method znet.Method, path string, query url.Values, body []byte) (*znet.Response, error) {
	builder := znet.NewRequestBuilder().Method(method).URL(c.buildURL(path, query))

	for name, value := range c.defaultHeaders {
		builder.SetHeader(name, value)
	}
	builder.SetHeader("User-Agent", UserAgent())
	builder.SetHeader("Authorization", "Zoho-oauthtoken "+c.tokens.Get())
	builder.SetHeader("Accept", "application/json")
	if body != nil {
		builder.SetHeader("Content-Type", "application/json")
		builder.Body(string(body))
	}
	if c.requestTimeout > 0 {
		builder.Timeout(c.requestTimeout)
	}

	response, err := c.transport.Execute(builder.Build())
	if err != nil {
		return nil, exception.NewConnectionError("request failed", err)
	}
	if response.StatusCode() < 200 || response.StatusCode() >= 300 {
		return nil, c.toAPIError(response)
	}
	return response, nil
}

func (c *Caller) toAPIError(response *znet.Response) error {
	var code, message string
	if strings.TrimSpace(response.Body()) != "" {
		var parsed map[string]any
		if json.Unmarshal([]byte(response.Body()), &parsed) == nil {
			code = stringField(parsed, "code")
			message = stringField(parsed, "message")
		}
	}
	return exception.FromResponse(response.StatusCode(), code, message)
}

func (c *Caller) GetInto(path string, query url.Values, result any, keys ...string) error {
	response, err := c.Do(znet.GET, path, query, nil)
	if err != nil {
		return err
	}
	return DecodeEnvelope([]byte(response.Body()), result, keys...)
}

func (c *Caller) PostInto(path string, body []byte, result any, keys ...string) error {
	response, err := c.Do(znet.POST, path, nil, body)
	if err != nil {
		return err
	}
	return DecodeEnvelope([]byte(response.Body()), result, keys...)
}

func (c *Caller) PutInto(path string, body []byte, result any, keys ...string) error {
	response, err := c.Do(znet.PUT, path, nil, body)
	if err != nil {
		return err
	}
	return DecodeEnvelope([]byte(response.Body()), result, keys...)
}

func (c *Caller) DeleteResource(path string) error {
	_, err := c.Do(znet.DELETE, path, nil, nil)
	return err
}

func (c *Caller) Put(path string) error {
	_, err := c.Do(znet.PUT, path, nil, nil)
	return err
}
