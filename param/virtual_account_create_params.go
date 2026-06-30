package param

// VirtualAccountCreateParams are the parameters for creating a virtual account.
type VirtualAccountCreateParams struct {
	Description     string
	CustomerID      *string
	MinimumAmount   *float64
	MaximumAmount   *float64
	ExpiresAt       *string
	ReferenceNumber *string
	MetaData        []MetaDataParams
}

func (p *VirtualAccountCreateParams) Validate() error {
	if err := Require("description", p.Description); err != nil {
		return err
	}
	if err := validateMaxLen("description", p.Description, maxDescriptionLength); err != nil {
		return err
	}
	if err := validateMaxLenPtr("reference_number", p.ReferenceNumber, maxReferenceLength); err != nil {
		return err
	}
	return validateMetaData(p.MetaData)
}

func (p *VirtualAccountCreateParams) ToBody() map[string]any {
	body := map[string]any{"description": p.Description}
	putStr(body, "customer_id", p.CustomerID)
	putFloat(body, "minimum_amount", p.MinimumAmount)
	putFloat(body, "maximum_amount", p.MaximumAmount)
	putStr(body, "expires_at", p.ExpiresAt)
	putStr(body, "reference_number", p.ReferenceNumber)
	if metaData := metaDataToList(p.MetaData); metaData != nil {
		body["meta_data"] = metaData
	}
	return body
}
