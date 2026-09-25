package param

import "github.com/zoho/zoho-payments-go-sdk/exception"

// ConnectedAccountCreateParams are the parameters for onboarding a connected account.
type ConnectedAccountCreateParams struct {
	AccountName                 string
	EmailID                     string
	Pan                         string
	Mcc                         string
	BusinessDescription         string
	ConnectedAccountBankAccount *ConnectedAccountBankAccountParams
}

// ConnectedAccountBankAccountParams is the bank account of the connected account being created.
type ConnectedAccountBankAccountParams struct {
	RoutingNumber string
	AccountNumber string
}

func (p *ConnectedAccountCreateParams) Validate() error {
	if err := Require("account_name", p.AccountName); err != nil {
		return err
	}
	if err := Require("email_id", p.EmailID); err != nil {
		return err
	}
	if err := Require("pan", p.Pan); err != nil {
		return err
	}
	if err := Require("mcc", p.Mcc); err != nil {
		return err
	}
	if err := Require("business_description", p.BusinessDescription); err != nil {
		return err
	}
	if p.ConnectedAccountBankAccount == nil {
		return exception.NewValidationError("connected_account_bank_account", "is required")
	}
	return p.ConnectedAccountBankAccount.validate()
}

func (p *ConnectedAccountCreateParams) ToBody() map[string]any {
	return map[string]any{
		"account_name":                   p.AccountName,
		"email_id":                       p.EmailID,
		"pan":                            p.Pan,
		"mcc":                            p.Mcc,
		"business_description":           p.BusinessDescription,
		"connected_account_bank_account": p.ConnectedAccountBankAccount.toMap(),
	}
}

func (b *ConnectedAccountBankAccountParams) validate() error {
	if err := Require("routing_number", b.RoutingNumber); err != nil {
		return err
	}
	return Require("account_number", b.AccountNumber)
}

func (b *ConnectedAccountBankAccountParams) toMap() map[string]any {
	return map[string]any{
		"routing_number": b.RoutingNumber,
		"account_number": b.AccountNumber,
	}
}
