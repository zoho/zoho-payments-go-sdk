package param

// PaymentMethodSessionCreateParams are the parameters for creating a payment method session (US edition).
type PaymentMethodSessionCreateParams struct {
	CustomerID  string
	Description *string
}

func (p *PaymentMethodSessionCreateParams) Validate() error {
	if err := Require("customer_id", p.CustomerID); err != nil {
		return err
	}
	return validateMaxLenPtr("description", p.Description, maxDescriptionLength)
}

func (p *PaymentMethodSessionCreateParams) ToBody() map[string]any {
	body := map[string]any{"customer_id": p.CustomerID}
	putStr(body, "description", p.Description)
	return body
}
