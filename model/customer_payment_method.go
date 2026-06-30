package model

type CustomerPaymentMethod struct {
	PaymentMethodID *string           `json:"payment_method_id,omitempty"`
	Type            *string           `json:"type,omitempty"`
	Brand           *string           `json:"brand,omitempty"`
	LastFourDigits  *string           `json:"last_four_digits,omitempty"`
	ExpiryMonth     *string           `json:"expiry_month,omitempty"`
	ExpiryYear      *string           `json:"expiry_year,omitempty"`
	Card            *CustomerCard     `json:"card,omitempty"`
	AchDebit        *CustomerAchDebit `json:"ach_debit,omitempty"`
}
