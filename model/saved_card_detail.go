package model

type SavedCardDetail struct {
	CardHolderName *string     `json:"card_holder_name,omitempty"`
	LastFourDigits *string     `json:"last_four_digits,omitempty"`
	ExpiryMonth    *string     `json:"expiry_month,omitempty"`
	ExpiryYear     *string     `json:"expiry_year,omitempty"`
	Brand          *string     `json:"brand,omitempty"`
	Funding        *string     `json:"funding,omitempty"`
	Country        *string     `json:"country,omitempty"`
	CardChecks     *CardChecks `json:"card_checks,omitempty"`
}
