package param

// PaymentLinkCreateParams are the parameters for creating a payment link.
type PaymentLinkCreateParams struct {
	Amount           float64
	Currency         string
	Description      string
	Email            *string
	Phone            *string
	PhoneCountryCode *string
	ExpiresAt        *string
	ReferenceID      *string
	ReturnURL        *string
	NotifyCustomer   *NotifyCustomerParams
	Configurations   *PaymentLinkConfigurationsParams
}

func (p *PaymentLinkCreateParams) Validate() error {
	if err := Require("currency", p.Currency); err != nil {
		return err
	}
	if err := Require("description", p.Description); err != nil {
		return err
	}
	if err := validateMaxLen("description", p.Description, maxDescriptionLength); err != nil {
		return err
	}
	return validateMaxLenPtr("reference_id", p.ReferenceID, maxReferenceLength)
}

func (p *PaymentLinkCreateParams) ToBody() map[string]any {
	body := map[string]any{
		"amount":      p.Amount,
		"currency":    p.Currency,
		"description": p.Description,
	}
	putStr(body, "email", p.Email)
	putStr(body, "phone", p.Phone)
	putStr(body, "phone_country_code", p.PhoneCountryCode)
	putStr(body, "expires_at", p.ExpiresAt)
	putStr(body, "reference_id", p.ReferenceID)
	putStr(body, "return_url", p.ReturnURL)
	if p.NotifyCustomer != nil {
		body["notify_customer"] = p.NotifyCustomer.toMap()
	}
	if p.Configurations != nil {
		body["configurations"] = p.Configurations.toMap()
	}
	return body
}
