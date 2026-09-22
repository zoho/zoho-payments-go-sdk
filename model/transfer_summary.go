package model

// TransferSummary is a transfer as returned by the transfer list endpoint.
type TransferSummary struct {
	TransferID           *string  `json:"transfer_id,omitempty"`
	PaymentID            *string  `json:"payment_id,omitempty"`
	TotalAmount          *Decimal `json:"total_amount,omitempty"`
	NetAmount            *Decimal `json:"net_amount,omitempty"`
	FeeAmount            *Decimal `json:"fee_amount,omitempty"`
	FeeRate              *Decimal `json:"fee_rate,omitempty"`
	FeeTaxAmount         *Decimal `json:"fee_tax_amount,omitempty"`
	ConnectedAccountID   *string  `json:"connected_account_id,omitempty"`
	ConnectedAccountName *string  `json:"connected_account_name,omitempty"`
	Description          *string  `json:"description,omitempty"`
	Currency             *string  `json:"currency,omitempty"`
	Status               *string  `json:"status,omitempty"`
	FailureCode          *string  `json:"failure_code,omitempty"`
	ReversedAmount       *Decimal `json:"reversed_amount,omitempty"`
	CreatedTime          *string  `json:"created_time,omitempty"`
}
