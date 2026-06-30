package model

type MandatePayment struct {
	PaymentsSessionID   *string                      `json:"payments_session_id,omitempty"`
	InvoiceNumber       *string                      `json:"invoice_number,omitempty"`
	CustomerID          *string                      `json:"customer_id,omitempty"`
	Amount              *Decimal                     `json:"amount,omitempty"`
	Currency            *string                      `json:"currency,omitempty"`
	Status              *string                      `json:"status,omitempty"`
	StatementDescriptor *string                      `json:"statement_descriptor,omitempty"`
	Description         *string                      `json:"description,omitempty"`
	ReferenceNumber     *string                      `json:"reference_number,omitempty"`
	Date                *int64                       `json:"date,omitempty"`
	PaymentMethod       *MandatePaymentPaymentMethod `json:"payment_method,omitempty"`
}
