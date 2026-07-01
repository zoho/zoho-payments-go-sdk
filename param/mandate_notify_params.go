package param

// MandateNotifyParams are the parameters for sending a mandate debit notification.
type MandateNotifyParams struct {
	MandateID     string
	Amount        float64
	ExecutionDate string
	Description   string
	InvoiceNumber string
}

func (p *MandateNotifyParams) Validate() error {
	if err := Require("mandate_id", p.MandateID); err != nil {
		return err
	}
	if err := Require("execution_date", p.ExecutionDate); err != nil {
		return err
	}
	if err := Require("description", p.Description); err != nil {
		return err
	}
	if err := Require("invoice_number", p.InvoiceNumber); err != nil {
		return err
	}
	if err := validateMaxLen("description", p.Description, maxDescriptionLength); err != nil {
		return err
	}
	return validateMaxLen("invoice_number", p.InvoiceNumber, maxInvoiceNumberLength)
}

func (p *MandateNotifyParams) ToBody() map[string]any {
	return map[string]any{
		"mandate_id":     p.MandateID,
		"amount":         p.Amount,
		"execution_date": p.ExecutionDate,
		"description":    p.Description,
		"invoice_number": p.InvoiceNumber,
	}
}
