package param

import "github.com/zoho/zoho-payments-go-sdk/exception"

// PaymentSessionCreateParams are the parameters for creating a payment session.
type PaymentSessionCreateParams struct {
	Amount          float64
	Currency        string
	Description     string
	ExpiresIn       *int
	MetaData        []MetaDataParams
	InvoiceNumber   *string
	ReferenceNumber *string
	MaxRetryCount   *int
	Configurations  *ConfigurationsParams
}

func (p *PaymentSessionCreateParams) Validate() error {
	if err := Require("currency", p.Currency); err != nil {
		return err
	}
	if err := Require("description", p.Description); err != nil {
		return err
	}
	if p.ExpiresIn != nil && (*p.ExpiresIn < 300 || *p.ExpiresIn > 900) {
		return exception.NewValidationError("expires_in", "must be between 300 and 900 seconds")
	}
	if p.MaxRetryCount != nil && (*p.MaxRetryCount < 1 || *p.MaxRetryCount > 5) {
		return exception.NewValidationError("max_retry_count", "must be between 1 and 5")
	}
	if err := validateMaxLen("description", p.Description, maxDescriptionLength); err != nil {
		return err
	}
	if err := validateMaxLenPtr("invoice_number", p.InvoiceNumber, maxInvoiceNumberLength); err != nil {
		return err
	}
	if err := validateMaxLenPtr("reference_number", p.ReferenceNumber, maxReferenceLength); err != nil {
		return err
	}
	if p.Configurations != nil {
		if err := p.Configurations.validate(); err != nil {
			return err
		}
	}
	return validateMetaData(p.MetaData)
}

func (p *PaymentSessionCreateParams) ToBody() map[string]any {
	body := map[string]any{
		"amount":      p.Amount,
		"currency":    p.Currency,
		"description": p.Description,
	}
	putInt(body, "expires_in", p.ExpiresIn)
	putStr(body, "invoice_number", p.InvoiceNumber)
	putStr(body, "reference_number", p.ReferenceNumber)
	putInt(body, "max_retry_count", p.MaxRetryCount)
	if metaData := metaDataToList(p.MetaData); metaData != nil {
		body["meta_data"] = metaData
	}
	if p.Configurations != nil {
		body["configurations"] = p.Configurations.toMap()
	}
	return body
}
