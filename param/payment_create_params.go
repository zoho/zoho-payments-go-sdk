package param

// PaymentCreateParams are the parameters for creating a payment with a saved payment method (US edition).
type PaymentCreateParams struct {
	CustomerID          string
	PaymentMethodID     string
	Amount              float64
	Currency            string
	CustomerOnSession   *bool
	BrowserInfo         *BrowserInfo
	StatementDescriptor *string
	Description         *string
	ShippingAddress     *PostalAddressParams
	MetaData            []MetaDataParams
}

func (p *PaymentCreateParams) Validate() error {
	if err := Require("customer_id", p.CustomerID); err != nil {
		return err
	}
	if err := Require("payment_method_id", p.PaymentMethodID); err != nil {
		return err
	}
	if err := Require("currency", p.Currency); err != nil {
		return err
	}
	if err := validateMaxLenPtr("description", p.Description, maxDescriptionLength); err != nil {
		return err
	}
	if p.ShippingAddress != nil {
		if err := p.ShippingAddress.validate(); err != nil {
			return err
		}
	}
	return validateMetaData(p.MetaData)
}

func (p *PaymentCreateParams) ToBody() map[string]any {
	body := map[string]any{
		"customer_id":       p.CustomerID,
		"payment_method_id": p.PaymentMethodID,
		"amount":            p.Amount,
		"currency":          p.Currency,
	}
	putBool(body, "customer_on_session", p.CustomerOnSession)
	if p.BrowserInfo != nil {
		body["browser_info"] = p.BrowserInfo.toMap()
	}
	putStr(body, "statement_descriptor", p.StatementDescriptor)
	putStr(body, "description", p.Description)
	if p.ShippingAddress != nil {
		body["shipping_address"] = p.ShippingAddress.toMap()
	}
	if metaData := metaDataToList(p.MetaData); metaData != nil {
		body["meta_data"] = metaData
	}
	return body
}
