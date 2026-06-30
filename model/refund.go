package model

type Refund struct {
	RefundID               *string    `json:"refund_id,omitempty"`
	PaymentID              *string    `json:"payment_id,omitempty"`
	ReferenceNumber        *string    `json:"reference_number,omitempty"`
	Amount                 *Decimal   `json:"amount,omitempty"`
	DefaultCurrencyAmount  *Decimal   `json:"default_currency_amount,omitempty"`
	Type                   *string    `json:"type,omitempty"`
	Reason                 *string    `json:"reason,omitempty"`
	Description            *string    `json:"description,omitempty"`
	Status                 *string    `json:"status,omitempty"`
	NetworkReferenceNumber *string    `json:"network_reference_number,omitempty"`
	FailureReason          *string    `json:"failure_reason,omitempty"`
	Date                   *int64     `json:"date,omitempty"`
	MetaData               []MetaData `json:"meta_data,omitempty"`
}
