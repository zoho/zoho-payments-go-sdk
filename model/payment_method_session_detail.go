package model

type PaymentMethodSessionDetail struct {
	PaymentMethodID *string `json:"payment_method_id,omitempty"`
	Status          *string `json:"status,omitempty"`
	CreatedTime     *int64  `json:"created_time,omitempty"`
	Type            *string `json:"type,omitempty"`
}
