package model

type VirtualAccountPayment struct {
	PaymentID                  *string                      `json:"payment_id,omitempty"`
	CustomerID                 *string                      `json:"customer_id,omitempty"`
	VirtualAccountID           *string                      `json:"virtual_account_id,omitempty"`
	CustomerName               *string                      `json:"customer_name,omitempty"`
	CustomerEmail              *string                      `json:"customer_email,omitempty"`
	Amount                     *Decimal                     `json:"amount,omitempty"`
	ReceiptEmail               *string                      `json:"receipt_email,omitempty"`
	DialingCode                *string                      `json:"dialing_code,omitempty"`
	Phone                      *string                      `json:"phone,omitempty"`
	ReferenceNumber            *string                      `json:"reference_number,omitempty"`
	TransactionReferenceNumber *string                      `json:"transaction_reference_number,omitempty"`
	PaymentType                *string                      `json:"payment_type,omitempty"`
	Currency                   *string                      `json:"currency,omitempty"`
	Balance                    *Decimal                     `json:"balance,omitempty"`
	AmountCaptured             *Decimal                     `json:"amount_captured,omitempty"`
	AmountRefunded             *Decimal                     `json:"amount_refunded,omitempty"`
	FeeAmount                  *Decimal                     `json:"fee_amount,omitempty"`
	Status                     *string                      `json:"status,omitempty"`
	TransactionType            *string                      `json:"transaction_type,omitempty"`
	FraudAlert                 *string                      `json:"fraud_alert,omitempty"`
	FailureReason              *string                      `json:"failure_reason,omitempty"`
	FailureCategory            *string                      `json:"failure_category,omitempty"`
	NextAction                 *string                      `json:"next_action,omitempty"`
	Tip                        *string                      `json:"tip,omitempty"`
	Date                       *int64                       `json:"date,omitempty"`
	PaymentMethod              *VirtualAccountPaymentMethod `json:"payment_method,omitempty"`
}
