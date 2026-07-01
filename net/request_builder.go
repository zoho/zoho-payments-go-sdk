package net

import "time"

// RequestBuilder builds an immutable Request.
type RequestBuilder struct {
	method  Method
	url     string
	headers map[string][]string
	body    string
	timeout time.Duration
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{headers: make(map[string][]string)}
}

func (b *RequestBuilder) Method(method Method) *RequestBuilder {
	b.method = method
	return b
}

func (b *RequestBuilder) URL(url string) *RequestBuilder {
	b.url = url
	return b
}

func (b *RequestBuilder) Header(name, value string) *RequestBuilder {
	b.headers[name] = append(b.headers[name], value)
	return b
}

func (b *RequestBuilder) SetHeader(name, value string) *RequestBuilder {
	b.headers[name] = []string{value}
	return b
}

func (b *RequestBuilder) Headers(headers map[string]string) *RequestBuilder {
	for name, value := range headers {
		b.headers[name] = append(b.headers[name], value)
	}
	return b
}

func (b *RequestBuilder) Body(body string) *RequestBuilder {
	b.body = body
	return b
}

func (b *RequestBuilder) Timeout(timeout time.Duration) *RequestBuilder {
	b.timeout = timeout
	return b
}

func (b *RequestBuilder) Build() *Request {
	headers := make(map[string][]string, len(b.headers))
	for name, headerValues := range b.headers {
		copied := make([]string, len(headerValues))
		copy(copied, headerValues)
		headers[name] = copied
	}
	return &Request{
		method:  b.method,
		url:     b.url,
		headers: headers,
		body:    b.body,
		timeout: b.timeout,
	}
}
