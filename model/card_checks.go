package model

type CardChecks struct {
	AddressLineCheck *string `json:"address_line_check,omitempty"`
	PostalCodeCheck  *string `json:"postal_code_check,omitempty"`
	CvcCheck         *string `json:"cvc_check,omitempty"`
}
