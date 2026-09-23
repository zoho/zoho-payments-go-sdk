package model

// ConnectedAccountSummary is a connected account as returned by the list endpoint.
type ConnectedAccountSummary struct {
	ConnectedAccountID *string `json:"connected_account_id,omitempty"`
	AccountName        *string `json:"account_name,omitempty"`
	EmailID            *string `json:"email_id,omitempty"`
	UnderWritingStatus *string `json:"under_writing_status,omitempty"`
	TransferStatus     *string `json:"transfer_status,omitempty"`
	CreatedBy          *string `json:"created_by,omitempty"`
	LastModifiedBy     *string `json:"last_modified_by,omitempty"`
	CreatedTime        *string `json:"created_time,omitempty"`
	LastModifiedTime   *string `json:"last_modified_time,omitempty"`
}
