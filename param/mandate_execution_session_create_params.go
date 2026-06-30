package param

import "github.com/zoho/zoho-payments-go-sdk/exception"

// MandateExecutionSessionCreateParams are the parameters for creating a mandate execution session.
type MandateExecutionSessionCreateParams struct {
	Amount        float64
	Currency      string
	CustomerID    string
	Description   string
	InvoiceNumber string
	MaxRetryCount *int
	MetaData      []MetaDataParams
}

func (p *MandateExecutionSessionCreateParams) Validate() error {
	if err := Require("currency", p.Currency); err != nil {
		return err
	}
	if err := Require("customer_id", p.CustomerID); err != nil {
		return err
	}
	if err := Require("description", p.Description); err != nil {
		return err
	}
	if err := Require("invoice_number", p.InvoiceNumber); err != nil {
		return err
	}
	if p.MaxRetryCount != nil && (*p.MaxRetryCount < 1 || *p.MaxRetryCount > 3) {
		return exception.NewValidationError("max_retry_count", "must be between 1 and 3")
	}
	if err := validateMaxLen("description", p.Description, maxDescriptionLength); err != nil {
		return err
	}
	return validateMaxLen("invoice_number", p.InvoiceNumber, maxInvoiceNumberLength)
}

func (p *MandateExecutionSessionCreateParams) ToBody() map[string]any {
	body := map[string]any{
		"type":           "mandate_execution",
		"amount":         p.Amount,
		"currency":       p.Currency,
		"customer_id":    p.CustomerID,
		"description":    p.Description,
		"invoice_number": p.InvoiceNumber,
	}
	putInt(body, "max_retry_count", p.MaxRetryCount)
	if metaData := metaDataToList(p.MetaData); metaData != nil {
		body["meta_data"] = metaData
	}
	return body
}
