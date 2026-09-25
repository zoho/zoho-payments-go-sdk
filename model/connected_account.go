package model

// ConnectedAccount is a split-settlement connected account.
type ConnectedAccount struct {
	ConnectedAccountID                 *string                       `json:"connected_account_id,omitempty"`
	EmailID                            *string                       `json:"email_id,omitempty"`
	AccountName                        *string                       `json:"account_name,omitempty"`
	Pan                                *string                       `json:"pan,omitempty"`
	Mcc                                *string                       `json:"mcc,omitempty"`
	BusinessDescription                *string                       `json:"business_description,omitempty"`
	UnderWritingStatus                 *string                       `json:"under_writing_status,omitempty"`
	TransferStatus                     *string                       `json:"transfer_status,omitempty"`
	PayoutStatus                       *string                       `json:"payout_status,omitempty"`
	PayoutDelayDays                    *int                          `json:"payout_delay_days,omitempty"`
	PayoutStatementDescriptor          *string                       `json:"payout_statement_descriptor,omitempty"`
	StatementDescriptorRestrictedChars *string                       `json:"statement_descriptor_restricted_chars,omitempty"`
	CreatedBy                          *string                       `json:"created_by,omitempty"`
	LastModifiedBy                     *string                       `json:"last_modified_by,omitempty"`
	CreatedTime                        *string                       `json:"created_time,omitempty"`
	LastModifiedTime                   *string                       `json:"last_modified_time,omitempty"`
	ConnectedAccountBankAccounts       []ConnectedAccountBankAccount `json:"connected_account_bank_accounts,omitempty"`
}

// ConnectedAccountBankAccount is a bank account linked to a connected account.
type ConnectedAccountBankAccount struct {
	ConnectedAccountBankAccountID *string `json:"connected_account_bank_account_id,omitempty"`
	Currency                      *string `json:"currency,omitempty"`
	RoutingNumber                 *string `json:"routing_number,omitempty"`
	LastFourDigits                *string `json:"last_four_digits,omitempty"`
	CreatedBy                     *string `json:"created_by,omitempty"`
	LastModifiedBy                *string `json:"last_modified_by,omitempty"`
	CreatedTime                   *string `json:"created_time,omitempty"`
	LastModifiedTime              *string `json:"last_modified_time,omitempty"`
}
