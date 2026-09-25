package model

// TransferCreateResponse is the per-split outcome of a transfer create request.
type TransferCreateResponse struct {
	Splits []TransferSplit `json:"splits,omitempty"`
	Status *string         `json:"status,omitempty"`
}

// TransferSplit is the outcome for one connected account within a transfer create request.
type TransferSplit struct {
	TransferID         *string  `json:"transfer_id,omitempty"`
	ConnectedAccountID *string  `json:"connected_account_id,omitempty"`
	Currency           *string  `json:"currency,omitempty"`
	TransferAmount     *Decimal `json:"transfer_amount,omitempty"`
	NetTransferAmount  *Decimal `json:"net_transfer_amount,omitempty"`
	FeeAmount          *Decimal `json:"fee_amount,omitempty"`
	FeeTaxAmount       *Decimal `json:"fee_tax_amount,omitempty"`
	Status             *string  `json:"status,omitempty"`
	ErrorCode          *string  `json:"error_code,omitempty"`
	Message            *string  `json:"message,omitempty"`
}
