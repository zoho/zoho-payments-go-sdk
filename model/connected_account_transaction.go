package model

// ConnectedAccountTransaction is a transaction on a connected account's balance.
type ConnectedAccountTransaction struct {
	TransactionID       *string  `json:"transaction_id,omitempty"`
	ParentTransactionID *string  `json:"parent_transaction_id,omitempty"`
	TransactionType     *string  `json:"transaction_type,omitempty"`
	TransactionTime     *string  `json:"transaction_time,omitempty"`
	NetAmount           *Decimal `json:"net_amount,omitempty"`
	Amount              *Decimal `json:"amount,omitempty"`
	Currency            *string  `json:"currency,omitempty"`
	Status              *string  `json:"status,omitempty"`
	AvailableOn         *string  `json:"available_on,omitempty"`
	SettledTime         *string  `json:"settled_time,omitempty"`
	PaymentMethod       *string  `json:"payment_method,omitempty"`
	CardBrand           *string  `json:"card_brand,omitempty"`
	CardType            *string  `json:"card_type,omitempty"`
	CustomerID          *string  `json:"customer_id,omitempty"`
	CustomerName        *string  `json:"customer_name,omitempty"`
	Description         *string  `json:"description,omitempty"`
	RefundReason        *string  `json:"refund_reason,omitempty"`
}
