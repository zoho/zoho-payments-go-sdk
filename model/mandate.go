package model

type Mandate struct {
	MandateID     *string               `json:"mandate_id,omitempty"`
	CustomerID    *string               `json:"customer_id,omitempty"`
	CustomerName  *string               `json:"customer_name,omitempty"`
	CustomerEmail *string               `json:"customer_email,omitempty"`
	CustomerPhone *string               `json:"customer_phone,omitempty"`
	Amount        *Decimal              `json:"amount,omitempty"`
	Currency      *string               `json:"currency,omitempty"`
	AmountRule    *string               `json:"amount_rule,omitempty"`
	Frequency     *string               `json:"frequency,omitempty"`
	Status        *string               `json:"status,omitempty"`
	Description   *string               `json:"description,omitempty"`
	DebitDay      *int                  `json:"debit_day,omitempty"`
	DebitRule     *string               `json:"debit_rule,omitempty"`
	StartDate     *int64                `json:"start_date,omitempty"`
	EndDate       *int64                `json:"end_date,omitempty"`
	PaymentMethod *MandatePaymentMethod `json:"payment_method,omitempty"`
}
