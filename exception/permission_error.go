package exception

// PermissionError is returned for HTTP 403 responses (insufficient permissions).
type PermissionError struct{ *APIError }

func (e *PermissionError) Unwrap() error { return e.APIError }
