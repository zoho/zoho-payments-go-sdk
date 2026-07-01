package model

type PaymentLinkPayment struct {
	PaymentID *string  `json:"payment_id,omitempty"`
	Amount    *Decimal `json:"amount,omitempty"`
	Type      *string  `json:"type,omitempty"`
	Status    *string  `json:"status,omitempty"`
	Date      *int64   `json:"date,omitempty"`
}
