package net

import "time"

// Request is an immutable HTTP request handed to an HTTPClient.
type Request struct {
	method  Method
	url     string
	headers map[string][]string
	body    string
	timeout time.Duration
}

func (r *Request) Method() Method { return r.method }

func (r *Request) URL() string { return r.url }

func (r *Request) Headers() map[string][]string {
	result := make(map[string][]string, len(r.headers))
	for name, headerValues := range r.headers {
		copied := make([]string, len(headerValues))
		copy(copied, headerValues)
		result[name] = copied
	}
	return result
}

func (r *Request) Body() string { return r.body }

func (r *Request) Timeout() time.Duration { return r.timeout }
