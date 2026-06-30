package param

// AchDebitUpdate carries updatable ACH debit fields for a saved payment method.
type AchDebitUpdate struct {
	AccountHolderType *string
}

func (a *AchDebitUpdate) toMap() map[string]any {
	body := map[string]any{}
	putStr(body, "account_holder_type", a.AccountHolderType)
	return body
}
