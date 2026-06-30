package param

// ConfigurationsParams configures allowed payment methods and the hosted page for a payment session.
type ConfigurationsParams struct {
	AllowedPaymentMethods []string
	HostedPageParameters  *HostedPageParams
}

func (c *ConfigurationsParams) validate() error {
	if c.HostedPageParameters != nil {
		return c.HostedPageParameters.validate()
	}
	return nil
}

func (c *ConfigurationsParams) toMap() map[string]any {
	body := map[string]any{}
	if len(c.AllowedPaymentMethods) > 0 {
		body["allowed_payment_methods"] = c.AllowedPaymentMethods
	}
	if c.HostedPageParameters != nil {
		body["hosted_page_parameters"] = c.HostedPageParameters.toMap()
	}
	return body
}
