package model

// Transfer is a single split-settlement transfer with its reversals and source payment.
type Transfer struct {
	TransferID                 *string                 `json:"transfer_id,omitempty"`
	PaymentID                  *string                 `json:"payment_id,omitempty"`
	TotalAmount                *Decimal                `json:"total_amount,omitempty"`
	NetAmount                  *Decimal                `json:"net_amount,omitempty"`
	FeeAmount                  *Decimal                `json:"fee_amount,omitempty"`
	FeeRate                    *Decimal                `json:"fee_rate,omitempty"`
	FeeTaxAmount               *Decimal                `json:"fee_tax_amount,omitempty"`
	FeeTaxRate                 *Decimal                `json:"fee_tax_rate,omitempty"`
	ConnectedAccountID         *string                 `json:"connected_account_id,omitempty"`
	ConnectedAccountName       *string                 `json:"connected_account_name,omitempty"`
	Description                *string                 `json:"description,omitempty"`
	Currency                   *string                 `json:"currency,omitempty"`
	Status                     *string                 `json:"status,omitempty"`
	FailureCode                *string                 `json:"failure_code,omitempty"`
	ReversedAmount             *Decimal                `json:"reversed_amount,omitempty"`
	AvailableAmountForReversal *Decimal                `json:"available_amount_for_reversal,omitempty"`
	CreatedTime                *string                 `json:"created_time,omitempty"`
	Reversals                  []TransferReversalEntry `json:"reversals,omitempty"`
	PaymentDetails             *TransferPaymentDetails `json:"payment_details,omitempty"`
}

// TransferReversalEntry is a reversal recorded against a transfer.
type TransferReversalEntry struct {
	ReversalID  *string  `json:"reversal_id,omitempty"`
	TotalAmount *Decimal `json:"total_amount,omitempty"`
	NetAmount   *Decimal `json:"net_amount,omitempty"`
	FeeAmount   *Decimal `json:"fee_amount,omitempty"`
	Description *string  `json:"description,omitempty"`
	Status      *string  `json:"status,omitempty"`
	CreatedTime *string  `json:"created_time,omitempty"`
}

// TransferPaymentDetails is the payment a transfer was split from.
type TransferPaymentDetails struct {
	Amount               *Decimal `json:"amount,omitempty"`
	AmountFormatted      *string  `json:"amount_formatted,omitempty"`
	Currency             *string  `json:"currency,omitempty"`
	PaymentDateFormatted *string  `json:"payment_date_formatted,omitempty"`
	StatusFormatted      *string  `json:"status_formatted,omitempty"`
	PaymentDate          *string  `json:"payment_date,omitempty"`
	Status               *string  `json:"status,omitempty"`
}
