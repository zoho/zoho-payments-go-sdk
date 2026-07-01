package param

import "github.com/zoho/zoho-payments-go-sdk/exception"

// VirtualAccountUpdateParams are the parameters for updating a virtual account.
type VirtualAccountUpdateParams struct {
	Description     *string
	MinimumAmount   *float64
	MaximumAmount   *float64
	ExpiresAt       *string
	ReferenceNumber *string
	MetaData        []MetaDataParams
}

func (p *VirtualAccountUpdateParams) hasAny() bool {
	return p.Description != nil || p.MinimumAmount != nil || p.MaximumAmount != nil ||
		p.ExpiresAt != nil || p.ReferenceNumber != nil || len(p.MetaData) > 0
}

func (p *VirtualAccountUpdateParams) Validate() error {
	if !p.hasAny() {
		return exception.NewValidationError("", "at least one field must be provided")
	}
	if err := validateMaxLenPtr("description", p.Description, maxDescriptionLength); err != nil {
		return err
	}
	if err := validateMaxLenPtr("reference_number", p.ReferenceNumber, maxReferenceLength); err != nil {
		return err
	}
	return validateMetaData(p.MetaData)
}

func (p *VirtualAccountUpdateParams) ToBody() map[string]any {
	body := map[string]any{}
	putStr(body, "description", p.Description)
	putFloat(body, "minimum_amount", p.MinimumAmount)
	putFloat(body, "maximum_amount", p.MaximumAmount)
	putStr(body, "expires_at", p.ExpiresAt)
	putStr(body, "reference_number", p.ReferenceNumber)
	if metaData := metaDataToList(p.MetaData); metaData != nil {
		body["meta_data"] = metaData
	}
	return body
}
