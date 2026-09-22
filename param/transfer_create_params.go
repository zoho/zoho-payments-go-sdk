package param

import (
	"fmt"

	"github.com/zoho/zoho-payments-go-sdk/exception"
)

// TransferCreateParams are the parameters for splitting a payment across connected accounts.
type TransferCreateParams struct {
	PaymentID     string
	TransferSplit []TransferSplitParams
}

// TransferSplitParams is a single split entry within a transfer request.
type TransferSplitParams struct {
	ConnectedAccountID string
	Amount             string
	Description        *string
}

func (p *TransferCreateParams) Validate() error {
	if err := Require("payment_id", p.PaymentID); err != nil {
		return err
	}
	if len(p.TransferSplit) == 0 {
		return exception.NewValidationError("transfer_split", "is required")
	}
	for index, split := range p.TransferSplit {
		if err := split.validate(index); err != nil {
			return err
		}
	}
	return nil
}

func (p *TransferCreateParams) ToBody() map[string]any {
	splits := make([]map[string]any, 0, len(p.TransferSplit))
	for _, split := range p.TransferSplit {
		splits = append(splits, split.toMap())
	}
	return map[string]any{
		"payment_id":     p.PaymentID,
		"transfer_split": splits,
	}
}

func (s TransferSplitParams) validate(index int) error {
	prefix := fmt.Sprintf("transfer_split[%d].", index)
	if err := Require(prefix+"connected_account_id", s.ConnectedAccountID); err != nil {
		return err
	}
	if err := Require(prefix+"amount", s.Amount); err != nil {
		return err
	}
	return validateMaxLenPtr(prefix+"description", s.Description, maxDescriptionLength)
}

func (s TransferSplitParams) toMap() map[string]any {
	body := map[string]any{
		"connected_account_id": s.ConnectedAccountID,
		"amount":               s.Amount,
	}
	putStr(body, "description", s.Description)
	return body
}
