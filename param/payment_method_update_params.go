package param

// PaymentMethodUpdateParams are the parameters for updating a saved payment method (US edition).
type PaymentMethodUpdateParams struct {
	Type           string
	Card           *CardUpdate
	AchDebit       *AchDebitUpdate
	BillingAddress *PostalAddressParams
}

func (p *PaymentMethodUpdateParams) Validate() error {
	if err := Require("type", p.Type); err != nil {
		return err
	}
	if p.BillingAddress != nil {
		return p.BillingAddress.validate()
	}
	return nil
}

func (p *PaymentMethodUpdateParams) ToBody() map[string]any {
	body := map[string]any{"type": p.Type}
	if p.Card != nil {
		body["card"] = p.Card.toMap()
	}
	if p.AchDebit != nil {
		body["ach_debit"] = p.AchDebit.toMap()
	}
	if p.BillingAddress != nil {
		body["billing_address"] = p.BillingAddress.toMap()
	}
	return body
}
