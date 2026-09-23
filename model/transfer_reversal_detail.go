package model

// TransferReversalDetail is a single transfer reversal.
type TransferReversalDetail struct {
	TransferReversalID   *string  `json:"transfer_reversal_id,omitempty"`
	TransferID           *string  `json:"transfer_id,omitempty"`
	PaymentID            *string  `json:"payment_id,omitempty"`
	RefundID             *string  `json:"refund_id,omitempty"`
	TotalAmount          *Decimal `json:"total_amount,omitempty"`
	ConnectedAccountID   *string  `json:"connected_account_id,omitempty"`
	ConnectedAccountName *string  `json:"connected_account_name,omitempty"`
	Status               *string  `json:"status,omitempty"`
	FailureCode          *string  `json:"failure_code,omitempty"`
	Description          *string  `json:"description,omitempty"`
	Currency             *string  `json:"currency,omitempty"`
	CreatedTime          *string  `json:"created_time,omitempty"`
}
