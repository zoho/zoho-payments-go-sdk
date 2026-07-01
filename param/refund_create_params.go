package param

// RefundCreateParams are the parameters for creating a refund.
type RefundCreateParams struct {
	Amount      float64
	Reason      string
	Type        string
	Description *string
	MetaData    []MetaDataParams
}

func (p *RefundCreateParams) Validate() error {
	if err := Require("reason", p.Reason); err != nil {
		return err
	}
	if err := Require("type", p.Type); err != nil {
		return err
	}
	if err := validateMaxLenPtr("description", p.Description, maxDescriptionLength); err != nil {
		return err
	}
	return validateMetaData(p.MetaData)
}

func (p *RefundCreateParams) ToBody() map[string]any {
	body := map[string]any{
		"amount": p.Amount,
		"reason": p.Reason,
		"type":   p.Type,
	}
	putStr(body, "description", p.Description)
	if metaData := metaDataToList(p.MetaData); metaData != nil {
		body["meta_data"] = metaData
	}
	return body
}
