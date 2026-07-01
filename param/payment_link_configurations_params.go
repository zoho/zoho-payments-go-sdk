package param

// PaymentLinkConfigurationsParams configures a payment link.
type PaymentLinkConfigurationsParams struct {
	AllowedPaymentMethods []string
}

func (p *PaymentLinkConfigurationsParams) toMap() map[string]any {
	body := map[string]any{}
	if len(p.AllowedPaymentMethods) > 0 {
		body["allowed_payment_methods"] = p.AllowedPaymentMethods
	}
	return body
}
