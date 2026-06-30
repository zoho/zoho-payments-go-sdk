package exception

// ResourceNotFoundError is returned for HTTP 404 responses.
type ResourceNotFoundError struct{ *APIError }

func (e *ResourceNotFoundError) Unwrap() error { return e.APIError }
