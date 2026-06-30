package param

import "github.com/zoho/zoho-payments-go-sdk/exception"

// MandateDetailsParams describes the mandate rules for an enrollment.
type MandateDetailsParams struct {
	PaymentMethodType string
	Frequency         string
	Description       string
	AmountRule        string
	MaxAmount         *float64
	StartDate         *string
	EndDate           *string
	DebitDay          *int
	DebitRule         *string
}

func (p *MandateDetailsParams) validate() error {
	if err := Require("payment_method_type", p.PaymentMethodType); err != nil {
		return err
	}
	if err := Require("frequency", p.Frequency); err != nil {
		return err
	}
	if err := Require("description", p.Description); err != nil {
		return err
	}
	if err := Require("amount_rule", p.AmountRule); err != nil {
		return err
	}
	if p.AmountRule == "variable" && p.MaxAmount == nil {
		return exception.NewValidationError("max_amount", "is required when amount_rule is \"variable\"")
	}
	return validateMaxLen("description", p.Description, maxDescriptionLength)
}

func (p *MandateDetailsParams) toMap() map[string]any {
	body := map[string]any{
		"payment_method_type": p.PaymentMethodType,
		"frequency":           p.Frequency,
		"description":         p.Description,
		"amount_rule":         p.AmountRule,
	}
	putFloat(body, "max_amount", p.MaxAmount)
	putStr(body, "start_date", p.StartDate)
	putStr(body, "end_date", p.EndDate)
	putInt(body, "debit_day", p.DebitDay)
	putStr(body, "debit_rule", p.DebitRule)
	return body
}
