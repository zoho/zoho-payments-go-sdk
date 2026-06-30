package model

type MandatePaymentMethod struct {
	Type *string     `json:"type,omitempty"`
	Upi  *MandateUpi `json:"upi,omitempty"`
}
