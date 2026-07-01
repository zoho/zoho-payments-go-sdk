package param

// PostalAddressParams is a postal address (Country is required).
type PostalAddressParams struct {
	Country      string
	Name         *string
	AddressLine1 *string
	AddressLine2 *string
	City         *string
	State        *string
	PostalCode   *string
}

func (p *PostalAddressParams) validate() error {
	return Require("country", p.Country)
}

func (p *PostalAddressParams) toMap() map[string]any {
	body := map[string]any{"country": p.Country}
	putStr(body, "name", p.Name)
	putStr(body, "address_line1", p.AddressLine1)
	putStr(body, "address_line2", p.AddressLine2)
	putStr(body, "city", p.City)
	putStr(body, "state", p.State)
	putStr(body, "postal_code", p.PostalCode)
	return body
}
