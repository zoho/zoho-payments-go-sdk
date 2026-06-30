package model

type PaymentMethodSession struct {
	PaymentMethodSessionID *string                     `json:"payment_method_session_id,omitempty"`
	CustomerID             *string                     `json:"customer_id,omitempty"`
	Description            *string                     `json:"description,omitempty"`
	CreatedTime            *int64                      `json:"created_time,omitempty"`
	Status                 *string                     `json:"status,omitempty"`
	PaymentMethod          *PaymentMethodSessionDetail `json:"payment_method,omitempty"`
}
