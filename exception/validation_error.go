package exception

import "fmt"

// ValidationError indicates a client-side parameter validation failure; the request was never sent.
type ValidationError struct {
	*ZohoPaymentsError
	Field string
}

func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{ZohoPaymentsError: NewZohoPaymentsError(message, nil), Field: field}
}

func (e *ValidationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("zohopayments: invalid parameter %q: %s", e.Field, e.Message)
	}
	return "zohopayments: " + e.Message
}

func (e *ValidationError) Unwrap() error { return e.ZohoPaymentsError }
