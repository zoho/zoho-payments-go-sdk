package model

// TransferReversalCreateResponse is the outcome of a transfer reversal create request.
type TransferReversalCreateResponse struct {
	TransferReversalID     *string  `json:"transfer_reversal_id,omitempty"`
	TransferID             *string  `json:"transfer_id,omitempty"`
	TransferReversalAmount *Decimal `json:"transfer_reversal_amount,omitempty"`
	ConnectedAccountID     *string  `json:"connected_account_id,omitempty"`
	Currency               *string  `json:"currency,omitempty"`
	Status                 *string  `json:"status,omitempty"`
}
