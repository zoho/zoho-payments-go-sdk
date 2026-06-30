package exception

import "fmt"

// APIError is the base error for any non-2xx API response.
type APIError struct {
	*ZohoPaymentsError
	HTTPStatusCode int
	Code           string
}

func NewAPIError(status int, code, message string) *APIError {
	return &APIError{
		ZohoPaymentsError: NewZohoPaymentsError(message, nil),
		HTTPStatusCode:    status,
		Code:              code,
	}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("zohopayments: API error (HTTP %d): code=%s, message=%s",
		e.HTTPStatusCode, e.Code, e.Message)
}

func (e *APIError) Unwrap() error { return e.ZohoPaymentsError }
