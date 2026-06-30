package model

type Upi struct {
	UpiID       *string `json:"upi_id,omitempty"`
	Channel     *string `json:"channel,omitempty"`
	AccountType *string `json:"account_type,omitempty"`
}
