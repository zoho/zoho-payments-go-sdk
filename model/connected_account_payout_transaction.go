package model

// ConnectedAccountPayoutTransaction is a single transaction settled within a connected account payout.
type ConnectedAccountPayoutTransaction struct {
	PayoutID              *string  `json:"payout_id,omitempty"`
	MerchantTransactionID *string  `json:"merchant_transaction_id,omitempty"`
	TransactionID         *string  `json:"transaction_id,omitempty"`
	ParentTransactionID   *string  `json:"parent_transaction_id,omitempty"`
	TransactionType       *string  `json:"transaction_type,omitempty"`
	TransactionTime       *string  `json:"transaction_time,omitempty"`
	NetAmount             *Decimal `json:"net_amount,omitempty"`
	Amount                *Decimal `json:"amount,omitempty"`
	Currency              *string  `json:"currency,omitempty"`
	CustomerID            *string  `json:"customer_id,omitempty"`
	CustomerName          *string  `json:"customer_name,omitempty"`
	Description           *string  `json:"description,omitempty"`
}
