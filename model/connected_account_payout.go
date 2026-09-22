package model

import "encoding/json"

// ConnectedAccountPayout is a single payout made to a connected account.
type ConnectedAccountPayout struct {
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
	Comments              []ConnectedAccountPayoutComment           `json:"comments,omitempty"`
	TransactionSummary    *ConnectedAccountPayoutTransactionSummary `json:"transaction_summary,omitempty"`
	AccountDetails        *ConnectedAccountPayoutAccountDetails     `json:"account_details,omitempty"`
}

// ConnectedAccountPayoutComment is an audit entry recorded against a connected account payout.
type ConnectedAccountPayoutComment struct {
	CommentID     *string  `json:"comment_id,omitempty"`
	Amount        *Decimal `json:"amount,omitempty"`
	OperationType *string  `json:"operation_type,omitempty"`
	ActionType    *string  `json:"action_type,omitempty"`
	Type          *string  `json:"type,omitempty"`
	Description   *string  `json:"description,omitempty"`
	CreatedBy     *string  `json:"created_by,omitempty"`
	CreatedTime   *string  `json:"created_time,omitempty"`
}

// ConnectedAccountPayoutTransactionSummary breaks a connected account payout down by transaction type.
type ConnectedAccountPayoutTransactionSummary struct {
	Charge      *ConnectedAccountPayoutTransactionBreakdown `json:"charge,omitempty"`
	Refund      *ConnectedAccountPayoutTransactionBreakdown `json:"refund,omitempty"`
	Adjustment  *ConnectedAccountPayoutTransactionBreakdown `json:"adjustment,omitempty"`
	TotalAmount *Decimal                                    `json:"total_amount,omitempty"`
}

// ConnectedAccountPayoutTransactionBreakdown is the aggregate for one transaction type within a connected account payout.
type ConnectedAccountPayoutTransactionBreakdown struct {
	TransactionType *string  `json:"transaction_type,omitempty"`
	Count           *int     `json:"count,omitempty"`
	NetAmount       *Decimal `json:"net_amount,omitempty"`
	Fee             *Decimal `json:"fee,omitempty"`
	Amount          *Decimal `json:"amount,omitempty"`
	Tax             *Decimal `json:"tax,omitempty"`
}

// ConnectedAccountPayoutAccountDetails is the destination account of a connected account payout.
type ConnectedAccountPayoutAccountDetails struct {
	BankName                    *string `json:"bank_name,omitempty"`
	AccountNumberLastFourDigits *string `json:"account_number_last_four_digits,omitempty"`
	RoutingNumber               *string `json:"routing_number,omitempty"`
	Type                        *string `json:"type,omitempty"`
	Country                     *string `json:"country,omitempty"`
	Currency                    *string `json:"currency,omitempty"`
}

// UnmarshalJSON normalises the API's date/time values, which arrive as strings
// from the list endpoints and as numbers from the detail endpoints.
func (c *ConnectedAccountPayout) UnmarshalJSON(data []byte) error {
	type alias ConnectedAccountPayout
	patched, err := coerceStrings(data, "initiated_time", "processed_date", "arrival_date")
	if err != nil {
		return err
	}
	return json.Unmarshal(patched, (*alias)(c))
}

// UnmarshalJSON normalises the API's date/time values, which arrive as strings
// from the list endpoints and as numbers from the detail endpoints.
func (c *ConnectedAccountPayoutComment) UnmarshalJSON(data []byte) error {
	type alias ConnectedAccountPayoutComment
	patched, err := coerceStrings(data, "created_time")
	if err != nil {
		return err
	}
	return json.Unmarshal(patched, (*alias)(c))
}
