package param

// TransferReversalCreateParams are the parameters for reversing a transfer.
type TransferReversalCreateParams struct {
	TransferID     string
	ReversalAmount string
	Description    *string
}

func (p *TransferReversalCreateParams) Validate() error {
	if err := Require("transfer_id", p.TransferID); err != nil {
		return err
	}
	if err := Require("reversal_amount", p.ReversalAmount); err != nil {
		return err
	}
	return validateMaxLenPtr("description", p.Description, maxDescriptionLength)
}

func (p *TransferReversalCreateParams) ToBody() map[string]any {
	body := map[string]any{
		"transfer_id":     p.TransferID,
		"reversal_amount": p.ReversalAmount,
	}
	putStr(body, "description", p.Description)
	return body
}
