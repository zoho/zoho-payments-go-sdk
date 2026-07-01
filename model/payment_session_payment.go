package model

type PaymentSessionPayment struct {
	PaymentID   *string `json:"payment_id,omitempty"`
	Status      *string `json:"status,omitempty"`
	CreatedTime *int64  `json:"created_time,omitempty"`
}
