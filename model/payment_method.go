package model

type PaymentMethod struct {
	PaymentMethodID  *string          `json:"payment_method_id,omitempty"`
	CustomerID       *string          `json:"customer_id,omitempty"`
	CustomerName     *string          `json:"customer_name,omitempty"`
	CustomerEmail    *string          `json:"customer_email,omitempty"`
	Type             *string          `json:"type,omitempty"`
	Status           *string          `json:"status,omitempty"`
	Currency         *string          `json:"currency,omitempty"`
	Source           *string          `json:"source,omitempty"`
	CreatedTime      *int64           `json:"created_time,omitempty"`
	LastModifiedTime *int64           `json:"last_modified_time,omitempty"`
	Card             *SavedCardDetail `json:"card,omitempty"`
	AchDebit         *AchDebitDetail  `json:"ach_debit,omitempty"`
	BillingAddress   *AddressDetail   `json:"billing_address,omitempty"`
}
