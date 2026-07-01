package param

import "github.com/zoho/zoho-payments-go-sdk/exception"

// PaymentLinkUpdateParams are the parameters for updating a payment link.
type PaymentLinkUpdateParams struct {
	Description      *string
	Email            *string
	Phone            *string
	PhoneCountryCode *string
	ExpiresAt        *string
	ReferenceID      *string
	ReturnURL        *string
	NotifyCustomer   *NotifyCustomerParams
	Configurations   *PaymentLinkConfigurationsParams
}

func (p *PaymentLinkUpdateParams) hasAny() bool {
	return p.Description != nil || p.Email != nil || p.Phone != nil ||
		p.PhoneCountryCode != nil || p.ExpiresAt != nil || p.ReferenceID != nil ||
		p.ReturnURL != nil || p.NotifyCustomer != nil || p.Configurations != nil
}

func (p *PaymentLinkUpdateParams) Validate() error {
	if !p.hasAny() {
		return exception.NewValidationError("", "at least one field must be provided")
	}
	if err := validateMaxLenPtr("description", p.Description, maxDescriptionLength); err != nil {
		return err
	}
	return validateMaxLenPtr("reference_id", p.ReferenceID, maxReferenceLength)
}

func (p *PaymentLinkUpdateParams) ToBody() map[string]any {
	body := map[string]any{}
	putStr(body, "description", p.Description)
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
