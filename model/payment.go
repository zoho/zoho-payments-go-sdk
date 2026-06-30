package model

type Payment struct {
	PaymentID                  *string              `json:"payment_id,omitempty"`
	Phone                      *string              `json:"phone,omitempty"`
	Amount                     *Decimal             `json:"amount,omitempty"`
	Currency                   *string              `json:"currency,omitempty"`
	PaymentsSessionID          *string              `json:"payments_session_id,omitempty"`
	ReceiptEmail               *string              `json:"receipt_email,omitempty"`
	ReferenceNumber            *string              `json:"reference_number,omitempty"`
	TransactionReferenceNumber *string              `json:"transaction_reference_number,omitempty"`
	InvoiceNumber              *string              `json:"invoice_number,omitempty"`
	AmountCaptured             *Decimal             `json:"amount_captured,omitempty"`
	AmountRefunded             *Decimal             `json:"amount_refunded,omitempty"`
	FeeAmount                  *Decimal             `json:"fee_amount,omitempty"`
	NetTaxAmount               *Decimal             `json:"net_tax_amount,omitempty"`
	TotalFeeAmount             *Decimal             `json:"total_fee_amount,omitempty"`
	NetAmount                  *Decimal             `json:"net_amount,omitempty"`
	Status                     *string              `json:"status,omitempty"`
	ExchangeRate               *float64             `json:"exchange_rate,omitempty"`
	StatementDescriptor        *string              `json:"statement_descriptor,omitempty"`
	Description                *string              `json:"description,omitempty"`
	Date                       *int64               `json:"date,omitempty"`
	PaymentMethod              *PaymentMethodDetail `json:"payment_method,omitempty"`
	MetaData                   []MetaData           `json:"meta_data,omitempty"`
}
