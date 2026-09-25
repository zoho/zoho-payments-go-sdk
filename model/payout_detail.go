package model

// PayoutDetail is a single payout with its transaction summary, comments and destination account.
type PayoutDetail struct {
	PayoutID              *string                   `json:"payout_id,omitempty"`
	Amount                *Decimal                  `json:"amount,omitempty"`
	ProcessingFee         *Decimal                  `json:"processing_fee,omitempty"`
	Currency              *string                   `json:"currency,omitempty"`
	Status                *string                   `json:"status,omitempty"`
	FailureCode           *string                   `json:"failure_code,omitempty"`
	FailureMessage        *string                   `json:"failure_message,omitempty"`
	StatementDescriptor   *string                   `json:"statement_descriptor,omitempty"`
	PayoutMethod          *string                   `json:"payout_method,omitempty"`
	InitiatedTime         *String                   `json:"initiated_time,omitempty"`
	ArrivalDate           *String                   `json:"arrival_date,omitempty"`
	ProcessedDate         *String                   `json:"processed_date,omitempty"`
	Type                  *string                   `json:"type,omitempty"`
	Fee                   *Decimal                  `json:"fee,omitempty"`
	FeeRate               *Decimal                  `json:"fee_rate,omitempty"`
	FixedFee              *Decimal                  `json:"fixed_fee,omitempty"`
	TaxAmount             *Decimal                  `json:"tax_amount,omitempty"`
	TaxRate               *Decimal                  `json:"tax_rate,omitempty"`
	PayoutBankReferenceID *string                   `json:"payout_bank_reference_id,omitempty"`
	Comments              []PayoutComment           `json:"comments,omitempty"`
	TransactionSummary    *PayoutTransactionSummary `json:"transaction_summary,omitempty"`
	AccountDetails        *PayoutAccountDetails     `json:"account_details,omitempty"`
}

// PayoutTransactionSummary breaks a payout down by transaction type.
type PayoutTransactionSummary struct {
	Charge      *PayoutTypeSummary `json:"charge,omitempty"`
	Refund      *PayoutTypeSummary `json:"refund,omitempty"`
	Adjustment  *PayoutTypeSummary `json:"adjustment,omitempty"`
	TotalAmount *Decimal           `json:"total_amount,omitempty"`
}

// PayoutTypeSummary is the aggregate for one transaction type within a payout.
type PayoutTypeSummary struct {
	TransactionType          *string  `json:"transaction_type,omitempty"`
	TransactionTypeFormatted *string  `json:"transaction_type_formatted,omitempty"`
	Count                    *int     `json:"count,omitempty"`
	NetAmount                *Decimal `json:"net_amount,omitempty"`
	NetAmountFormatted       *string  `json:"net_amount_formatted,omitempty"`
	Fee                      *Decimal `json:"fee,omitempty"`
	FeeFormatted             *string  `json:"fee_formatted,omitempty"`
	Amount                   *Decimal `json:"amount,omitempty"`
	AmountFormatted          *string  `json:"amount_formatted,omitempty"`
	Tax                      *Decimal `json:"tax,omitempty"`
	TaxFormatted             *string  `json:"tax_formatted,omitempty"`
}

// PayoutAccountDetails is the destination account of a payout.
type PayoutAccountDetails struct {
	BankName                    *string `json:"bank_name,omitempty"`
	AccountHolder               *string `json:"account_holder,omitempty"`
	AccountNumberLastFourDigits *string `json:"account_number_last_four_digits,omitempty"`
	RoutingNumber               *string `json:"routing_number,omitempty"`
	Type                        *string `json:"type,omitempty"`
	Country                     *string `json:"country,omitempty"`
	Currency                    *string `json:"currency,omitempty"`
}

// PayoutComment is an audit entry recorded against a payout.
type PayoutComment struct {
	CommentID     *string  `json:"comment_id,omitempty"`
	Amount        *Decimal `json:"amount,omitempty"`
	OperationType *string  `json:"operation_type,omitempty"`
	ActionType    *string  `json:"action_type,omitempty"`
	Type          *string  `json:"type,omitempty"`
	Description   *string  `json:"description,omitempty"`
	CreatedBy     *string  `json:"created_by,omitempty"`
	CreatedTime   *String  `json:"created_time,omitempty"`
}
