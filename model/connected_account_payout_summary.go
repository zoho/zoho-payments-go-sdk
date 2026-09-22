package model

// ConnectedAccountPayoutSummary is a connected account payout as returned by the list endpoint.
type ConnectedAccountPayoutSummary struct {
	PayoutID              *string                                   `json:"payout_id,omitempty"`
	Amount                *Decimal                                  `json:"amount,omitempty"`
	Currency              *string                                   `json:"currency,omitempty"`
	Status                *string                                   `json:"status,omitempty"`
	FailureCode           *string                                   `json:"failure_code,omitempty"`
	FailureMessage        *string                                   `json:"failure_message,omitempty"`
	StatementDescriptor   *string                                   `json:"statement_descriptor,omitempty"`
	PayoutMethod          *string                                   `json:"payout_method,omitempty"`
	InitiatedTime         *string                                   `json:"initiated_time,omitempty"`
	ArrivalDate           *string                                   `json:"arrival_date,omitempty"`
	ProcessedDate         *string                                   `json:"processed_date,omitempty"`
	Type                  *string                                   `json:"type,omitempty"`
	PayoutBankReferenceID *string                                   `json:"payout_bank_reference_id,omitempty"`
	BankAccountDetails    *ConnectedAccountPayoutBankAccountDetails `json:"bank_account_details,omitempty"`
}

// ConnectedAccountPayoutBankAccountDetails is the bank account a connected account payout is sent to.
type ConnectedAccountPayoutBankAccountDetails struct {
	BankName                    *string `json:"bank_name,omitempty"`
	AccountNumberLastFourDigits *string `json:"account_number_last_four_digits,omitempty"`
}
