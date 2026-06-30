package model

type Configurations struct {
	AllowedPaymentMethods []string            `json:"allowed_payment_methods,omitempty"`
	HostedPageParameters  *HostedPageResponse `json:"hosted_page_parameters,omitempty"`
}
