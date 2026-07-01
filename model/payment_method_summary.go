package model

type PaymentMethodSummary struct {
	PaymentMethodID *string `json:"payment_method_id,omitempty"`
	Type            *string `json:"type,omitempty"`
}
