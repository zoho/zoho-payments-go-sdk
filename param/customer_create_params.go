package param

// CustomerCreateParams are the parameters for creating a customer.
type CustomerCreateParams struct {
	Name             string
	Email            string
	Phone            *string
	PhoneCountryCode *string
	MetaData         []MetaDataParams
}

func (p *CustomerCreateParams) Validate() error {
	if err := Require("name", p.Name); err != nil {
		return err
	}
	if err := Require("email", p.Email); err != nil {
		return err
	}
	return validateMetaData(p.MetaData)
}

func (p *CustomerCreateParams) ToBody() map[string]any {
	body := map[string]any{
		"name":  p.Name,
		"email": p.Email,
	}
	putStr(body, "phone", p.Phone)
	putStr(body, "phone_country_code", p.PhoneCountryCode)
	if metaData := metaDataToList(p.MetaData); metaData != nil {
		body["meta_data"] = metaData
	}
	return body
}
