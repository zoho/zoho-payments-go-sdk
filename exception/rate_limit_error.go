package exception

// RateLimitError is returned for HTTP 429 responses (too many requests).
type RateLimitError struct{ *APIError }

func (e *RateLimitError) Unwrap() error { return e.APIError }
