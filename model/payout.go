package model

// Payout is a payout as returned by the payout list endpoint.
type Payout struct {
	PayoutID              *string                   `json:"payout_id,omitempty"`
	Amount                *Decimal                  `json:"amount,omitempty"`
	Currency              *string                   `json:"currency,omitempty"`
	Status                *string                   `json:"status,omitempty"`
	FailureCode           *string                   `json:"failure_code,omitempty"`
	FailureMessage        *string                   `json:"failure_message,omitempty"`
	StatementDescriptor   *string                   `json:"statement_descriptor,omitempty"`
	PayoutMethod          *string                   `json:"payout_method,omitempty"`
	InitiatedTime         *string                   `json:"initiated_time,omitempty"`
	ArrivalDate           *string                   `json:"arrival_date,omitempty"`
	ProcessedDate         *string                   `json:"processed_date,omitempty"`
	Type                  *string                   `json:"type,omitempty"`
	PayoutBankReferenceID *string                   `json:"payout_bank_reference_id,omitempty"`
	BankAccountDetails    *PayoutBankAccountDetails `json:"bank_account_details,omitempty"`
}

// PayoutBankAccountDetails is the bank account a payout is sent to.
type PayoutBankAccountDetails struct {
	BankName                    *string `json:"bank_name,omitempty"`
	AccountNumberLastFourDigits *string `json:"account_number_last_four_digits,omitempty"`
	AccountHolderName           *string `json:"account_holder_name,omitempty"`
	RoutingNumber               *string `json:"routing_number,omitempty"`
}
