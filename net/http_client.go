package net

// HTTPClient is the pluggable transport contract.
type HTTPClient interface {
	Execute(request *Request) (*Response, error)
}
