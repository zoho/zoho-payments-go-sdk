// Package exception defines the Zoho Payments SDK error hierarchy.
package exception

// ZohoPaymentsError is the base type embedded by every error the SDK returns.
type ZohoPaymentsError struct {
	Message string
	cause   error
}

func NewZohoPaymentsError(message string, cause error) *ZohoPaymentsError {
	return &ZohoPaymentsError{Message: message, cause: cause}
}

func (e *ZohoPaymentsError) Error() string { return e.Message }

func (e *ZohoPaymentsError) Unwrap() error { return e.cause }
