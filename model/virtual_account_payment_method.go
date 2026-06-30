package model

type VirtualAccountPaymentMethod struct {
	PaymentMethodID *string `json:"payment_method_id,omitempty"`
	Type            *string `json:"type,omitempty"`
}
