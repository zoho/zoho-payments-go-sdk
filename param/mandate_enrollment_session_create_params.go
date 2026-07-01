package param

import "github.com/zoho/zoho-payments-go-sdk/exception"

// MandateEnrollmentSessionCreateParams are the parameters for creating a mandate enrollment session.
type MandateEnrollmentSessionCreateParams struct {
	Amount         float64
	Currency       string
	CustomerID     string
	Description    string
	MandateDetails *MandateDetailsParams
	InvoiceNumber  *string
	MaxRetryCount  *int
	MetaData       []MetaDataParams
	Configurations *MandateConfigurationsParams
}

func (p *MandateEnrollmentSessionCreateParams) Validate() error {
	if err := Require("currency", p.Currency); err != nil {
		return err
	}
	if err := Require("customer_id", p.CustomerID); err != nil {
		return err
	}
	if err := Require("description", p.Description); err != nil {
		return err
	}
	if p.MandateDetails == nil {
		return exception.NewValidationError("mandate_details", "is required")
	}
	if err := p.MandateDetails.validate(); err != nil {
		return err
	}
	if p.MaxRetryCount != nil && (*p.MaxRetryCount < 1 || *p.MaxRetryCount > 3) {
		return exception.NewValidationError("max_retry_count", "must be between 1 and 3")
	}
	if err := validateMaxLen("description", p.Description, maxDescriptionLength); err != nil {
		return err
	}
	if err := validateMaxLenPtr("invoice_number", p.InvoiceNumber, maxInvoiceNumberLength); err != nil {
		return err
	}
	if p.Configurations != nil {
		if err := p.Configurations.validate(); err != nil {
			return err
		}
	}
	return validateMetaData(p.MetaData)
}

func (p *MandateEnrollmentSessionCreateParams) ToBody() map[string]any {
	body := map[string]any{
		"type":            "mandate_enrollment",
		"amount":          p.Amount,
		"currency":        p.Currency,
		"customer_id":     p.CustomerID,
		"description":     p.Description,
		"mandate_details": p.MandateDetails.toMap(),
	}
	putStr(body, "invoice_number", p.InvoiceNumber)
	putInt(body, "max_retry_count", p.MaxRetryCount)
	if metaData := metaDataToList(p.MetaData); metaData != nil {
		body["meta_data"] = metaData
	}
	if p.Configurations != nil {
		body["configurations"] = p.Configurations.toMap()
	}
	return body
}
