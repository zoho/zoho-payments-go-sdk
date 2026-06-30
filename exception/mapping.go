package exception

func FromResponse(status int, code, message string) error {
	base := NewAPIError(status, code, message)
	switch status {
	case 400, 422:
		return &InvalidRequestError{base}
	case 401:
		return &AuthenticationError{base}
	case 403:
		return &PermissionError{base}
	case 404:
		return &ResourceNotFoundError{base}
	case 429:
		return &RateLimitError{base}
	default:
		return base
	}
}
