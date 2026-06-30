package model

type BankTransfer struct {
	VirtualAccountNumber *string `json:"virtual_account_number,omitempty"`
	Mode                 *string `json:"mode,omitempty"`
	PayerName            *string `json:"payer_name,omitempty"`
	PayerAccountNo       *string `json:"payer_account_no,omitempty"`
	PayerIfscCode        *string `json:"payer_ifsc_code,omitempty"`
}
