package exception

// AuthenticationError is returned for HTTP 401 responses (invalid or expired OAuth token).
type AuthenticationError struct{ *APIError }

func (e *AuthenticationError) Unwrap() error { return e.APIError }
