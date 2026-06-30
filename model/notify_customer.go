package model

type NotifyCustomer struct {
	Email *bool `json:"email,omitempty"`
	SMS   *bool `json:"sms,omitempty"`
}
