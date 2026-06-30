package model

type MandateNotification struct {
	MandateID             *string               `json:"mandate_id,omitempty"`
	MandateNotificationID *string               `json:"mandate_notification_id,omitempty"`
	CustomerID            *string               `json:"customer_id,omitempty"`
	MandateAmount         *Decimal              `json:"mandate_amount,omitempty"`
	Currency              *string               `json:"currency,omitempty"`
	AmountRule            *string               `json:"amount_rule,omitempty"`
	NotificationAmount    *Decimal              `json:"notification_amount,omitempty"`
	NotificationStatus    *string               `json:"notification_status,omitempty"`
	Description           *string               `json:"description,omitempty"`
	InvoiceNumber         *string               `json:"invoice_number,omitempty"`
	Amount                *Decimal              `json:"amount,omitempty"`
	NotificationDate      *int64                `json:"notification_date,omitempty"`
	ExecutionDate         *int64                `json:"execution_date,omitempty"`
	PaymentMethod         *MandatePaymentMethod `json:"payment_method,omitempty"`
}
