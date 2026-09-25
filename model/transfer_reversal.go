package model

// TransferReversal is a transfer reversal as returned by the reversal list endpoint.
type TransferReversal struct {
	TransferReversalID   *string  `json:"transfer_reversal_id,omitempty"`
	Currency             *string  `json:"currency,omitempty"`
	TotalAmount          *Decimal `json:"total_amount,omitempty"`
	ConnectedAccountID   *string  `json:"connected_account_id,omitempty"`
	ConnectedAccountName *string  `json:"connected_account_name,omitempty"`
	Status               *string  `json:"status,omitempty"`
	FailureCode          *string  `json:"failure_code,omitempty"`
	CreatedTime          *string  `json:"created_time,omitempty"`
	TransferID           *string  `json:"transfer_id,omitempty"`
}
