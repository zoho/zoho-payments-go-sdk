package exception

// InvalidRequestError is returned for HTTP 400 and 422 responses (malformed request or a validation error reported by the API).
type InvalidRequestError struct{ *APIError }

func (e *InvalidRequestError) Unwrap() error { return e.APIError }
