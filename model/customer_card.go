package model

type CustomerCard struct {
	CardHolderName *string `json:"card_holder_name,omitempty"`
	LastFourDigits *string `json:"last_four_digits,omitempty"`
	ExpiryMonth    *string `json:"expiry_month,omitempty"`
	ExpiryYear     *string `json:"expiry_year,omitempty"`
}
