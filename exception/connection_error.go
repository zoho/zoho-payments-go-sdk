package exception

import "fmt"

// ConnectionError wraps a network or I/O failure that occurred before a complete API response was received.
type ConnectionError struct {
	*ZohoPaymentsError
}

func NewConnectionError(message string, cause error) *ConnectionError {
	return &ConnectionError{NewZohoPaymentsError(message, cause)}
}

func (e *ConnectionError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("zohopayments: connection error: %s: %v", e.Message, e.cause)
	}
	return "zohopayments: connection error: " + e.Message
}

func (e *ConnectionError) Unwrap() error { return e.ZohoPaymentsError }
