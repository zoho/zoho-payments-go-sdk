package param

// MandateConfigurationsParams configures the hosted page for a mandate session.
type MandateConfigurationsParams struct {
	HostedPageParameters *HostedPageParams
}

func (p *MandateConfigurationsParams) validate() error {
	if p.HostedPageParameters != nil {
		return p.HostedPageParameters.validate()
	}
	return nil
}

func (p *MandateConfigurationsParams) toMap() map[string]any {
	body := map[string]any{}
	if p.HostedPageParameters != nil {
		body["hosted_page_parameters"] = p.HostedPageParameters.toMap()
	}
	return body
}
