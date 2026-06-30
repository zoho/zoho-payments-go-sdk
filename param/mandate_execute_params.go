package param

// MandateExecuteParams are the parameters for executing a mandate debit.
type MandateExecuteParams struct {
	CustomerID            string
	MandateID             string
	PaymentsSessionID     string
	InvoiceNumber         string
	Amount                float64
	MandateNotificationID *string
	ReceiptEmail          *string
	Phone                 *string
	PhoneCountryCode      *string
	Description           *string
	ReferenceNumber       *string
}

func (p *MandateExecuteParams) Validate() error {
	if err := Require("customer_id", p.CustomerID); err != nil {
		return err
	}
	if err := Require("mandate_id", p.MandateID); err != nil {
		return err
	}
	if err := Require("payments_session_id", p.PaymentsSessionID); err != nil {
		return err
	}
	if err := Require("invoice_number", p.InvoiceNumber); err != nil {
		return err
	}
	if err := validateMaxLenPtr("description", p.Description, maxDescriptionLength); err != nil {
		return err
	}
	if err := validateMaxLen("invoice_number", p.InvoiceNumber, maxInvoiceNumberLength); err != nil {
		return err
	}
	return validateMaxLenPtr("reference_number", p.ReferenceNumber, maxReferenceLength)
}

func (p *MandateExecuteParams) ToBody() map[string]any {
	body := map[string]any{
		"customer_id":         p.CustomerID,
		"mandate_id":          p.MandateID,
		"payments_session_id": p.PaymentsSessionID,
		"invoice_number":      p.InvoiceNumber,
		"amount":              p.Amount,
	}
	putStr(body, "mandate_notification_id", p.MandateNotificationID)
	putStr(body, "receipt_email", p.ReceiptEmail)
	putStr(body, "phone", p.Phone)
	putStr(body, "phone_country_code", p.PhoneCountryCode)
	putStr(body, "description", p.Description)
	putStr(body, "reference_number", p.ReferenceNumber)
	return body
}
