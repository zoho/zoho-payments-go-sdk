package net

import (
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

// DefaultHTTPClient is the SDK's net/http-backed HTTPClient implementation.
type DefaultHTTPClient struct {
	client         *http.Client
	defaultTimeout time.Duration
}

func NewDefaultHTTPClient(connectTimeout, requestTimeout time.Duration) *DefaultHTTPClient {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   connectTimeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   connectTimeout,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return &DefaultHTTPClient{
		client:         &http.Client{Transport: transport},
		defaultTimeout: requestTimeout,
	}
}

func (c *DefaultHTTPClient) Execute(request *Request) (*Response, error) {
	var bodyReader io.Reader
	if request.Body() != "" {
		bodyReader = strings.NewReader(request.Body())
	}

	httpRequest, err := http.NewRequest(string(request.Method()), request.URL(), bodyReader)
	if err != nil {
		return nil, err
	}

	for name, values := range request.Headers() {
		for _, value := range values {
			httpRequest.Header.Add(name, value)
		}
	}

	timeout := request.Timeout()
	if timeout <= 0 {
		timeout = c.defaultTimeout
	}

	client := c.client
	if timeout > 0 {
		clone := *c.client
		clone.Timeout = timeout
		client = &clone
	}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	responseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	headers := make(map[string][]string, len(httpResponse.Header))
	for name, headerValues := range httpResponse.Header {
		copied := make([]string, len(headerValues))
		copy(copied, headerValues)
		headers[name] = copied
	}

	return NewResponse(httpResponse.StatusCode, headers, string(responseBody)), nil
}

func (c *DefaultHTTPClient) Close() error {
	if transport, ok := c.client.Transport.(*http.Transport); ok {
		transport.CloseIdleConnections()
	}
	return nil
}
