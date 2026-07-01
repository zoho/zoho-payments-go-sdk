package model

type AddressDetail struct {
	Name         *string `json:"name,omitempty"`
	AddressID    *string `json:"address_id,omitempty"`
	AddressLine1 *string `json:"address_line1,omitempty"`
	AddressLine2 *string `json:"address_line2,omitempty"`
	City         *string `json:"city,omitempty"`
	State        *string `json:"state,omitempty"`
	PostalCode   *string `json:"postal_code,omitempty"`
	Country      *string `json:"country,omitempty"`
}
