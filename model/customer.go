package model

type Customer struct {
	CustomerID       *string                 `json:"customer_id,omitempty"`
	Name             *string                 `json:"name,omitempty"`
	Email            *string                 `json:"email,omitempty"`
	Phone            *string                 `json:"phone,omitempty"`
	DialingCode      *string                 `json:"dialing_code,omitempty"`
	CreatedTime      *int64                  `json:"created_time,omitempty"`
	LastModifiedTime *int64                  `json:"last_modified_time,omitempty"`
	MetaData         []MetaData              `json:"meta_data,omitempty"`
	PaymentMethods   []CustomerPaymentMethod `json:"payment_methods,omitempty"`
}
