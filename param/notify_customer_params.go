package param

// NotifyCustomerParams toggles email/SMS notifications for a payment link.
type NotifyCustomerParams struct {
	Email *bool
	SMS   *bool
}

func (n *NotifyCustomerParams) toMap() map[string]any {
	body := map[string]any{}
	putBool(body, "email", n.Email)
	putBool(body, "sms", n.SMS)
	return body
}
