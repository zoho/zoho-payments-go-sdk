package model

type CustomerAchDebit struct {
	AccountHolderName *string `json:"account_holder_name,omitempty"`
	LastFourDigits    *string `json:"last_four_digits,omitempty"`
	AccountHolderType *string `json:"account_holder_type,omitempty"`
	AccountType       *string `json:"account_type,omitempty"`
	BankName          *string `json:"bank_name,omitempty"`
	RoutingNumber     *string `json:"routing_number,omitempty"`
}
