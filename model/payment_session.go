package model

type PaymentSession struct {
	PaymentsSessionID *string                 `json:"payments_session_id,omitempty"`
	AccessKey         *string                 `json:"access_key,omitempty"`
	Currency          *string                 `json:"currency,omitempty"`
	Amount            *Decimal                `json:"amount,omitempty"`
	Status            *string                 `json:"status,omitempty"`
	CreatedTime       *int64                  `json:"created_time,omitempty"`
	ExpiryTime        *int64                  `json:"expiry_time,omitempty"`
	Payments          []PaymentSessionPayment `json:"payments,omitempty"`
	MetaData          []MetaData              `json:"meta_data,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	InvoiceNumber     *string                 `json:"invoice_number,omitempty"`
	ReferenceNumber   *string                 `json:"reference_number,omitempty"`
	MaxRetryCount     *int                    `json:"max_retry_count,omitempty"`
	Configurations    *Configurations         `json:"configurations,omitempty"`
}
