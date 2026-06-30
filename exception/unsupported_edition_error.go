package exception

import (
	"fmt"

	"github.com/zoho/zoho-payments-go-sdk/edition"
)

// UnsupportedEditionError indicates an operation that is not available for the client's configured edition.
type UnsupportedEditionError struct {
	*ZohoPaymentsError
	Operation string
	Edition   edition.Edition
}

func NewUnsupportedEditionError(operation string, ed edition.Edition) *UnsupportedEditionError {
	msg := fmt.Sprintf("zohopayments: operation %q is not supported for edition %s", operation, ed)
	return &UnsupportedEditionError{
		ZohoPaymentsError: NewZohoPaymentsError(msg, nil),
		Operation:         operation,
		Edition:           ed,
	}
}

func (e *UnsupportedEditionError) Unwrap() error { return e.ZohoPaymentsError }
