package model

type PaymentSummary struct {
	PaymentID       *string               `json:"payment_id,omitempty"`
	Amount          *Decimal              `json:"amount,omitempty"`
	Currency        *string               `json:"currency,omitempty"`
	ReceiptEmail    *string               `json:"receipt_email,omitempty"`
	ReferenceNumber *string               `json:"reference_number,omitempty"`
	AmountCaptured  *Decimal              `json:"amount_captured,omitempty"`
	AmountRefunded  *Decimal              `json:"amount_refunded,omitempty"`
	FeeAmount       *Decimal              `json:"fee_amount,omitempty"`
	NetAmount       *Decimal              `json:"net_amount,omitempty"`
	Status          *string               `json:"status,omitempty"`
	Date            *int64                `json:"date,omitempty"`
	PaymentMethod   *PaymentMethodSummary `json:"payment_method,omitempty"`
}
