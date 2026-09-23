package model

// PayoutTransaction is a single transaction settled within a payout.
type PayoutTransaction struct {
	PayoutID            *string  `json:"payout_id,omitempty"`
	TransactionID       *string  `json:"transaction_id,omitempty"`
	ParentTransactionID *string  `json:"parent_transaction_id,omitempty"`
	NetAmount           *Decimal `json:"net_amount,omitempty"`
	Amount              *Decimal `json:"amount,omitempty"`
	Fee                 *Decimal `json:"fee,omitempty"`
	Tax                 *Decimal `json:"tax,omitempty"`
	TransactionType     *string  `json:"transaction_type,omitempty"`
	TransactionTime     *string  `json:"transaction_time,omitempty"`
	Currency            *string  `json:"currency,omitempty"`
	CustomerID          *string  `json:"customer_id,omitempty"`
	CustomerName        *string  `json:"customer_name,omitempty"`
	Description         *string  `json:"description,omitempty"`
}
