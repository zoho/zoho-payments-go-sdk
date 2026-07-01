package net

// Response is an HTTP response produced by an HTTPClient.
type Response struct {
	statusCode int
	headers    map[string][]string
	body       string
}

func NewResponse(statusCode int, headers map[string][]string, body string) *Response {
	if headers == nil {
		headers = map[string][]string{}
	}
	return &Response{statusCode: statusCode, headers: headers, body: body}
}

func (r *Response) StatusCode() int { return r.statusCode }

func (r *Response) Headers() map[string][]string { return r.headers }

func (r *Response) Body() string { return r.body }
