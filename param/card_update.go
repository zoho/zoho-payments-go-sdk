package param

// CardUpdate carries updatable card fields for a saved payment method.
type CardUpdate struct {
	ExpiryMonth *string
	ExpiryYear  *string
}

func (c *CardUpdate) toMap() map[string]any {
	body := map[string]any{}
	putStr(body, "expiry_month", c.ExpiryMonth)
	putStr(body, "expiry_year", c.ExpiryYear)
	return body
}
