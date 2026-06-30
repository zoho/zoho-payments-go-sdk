package model

type CustomerSummary struct {
	CustomerID       *string `json:"customer_id,omitempty"`
	CustomerName     *string `json:"customer_name,omitempty"`
	CustomerEmail    *string `json:"customer_email,omitempty"`
	CustomerPhone    *string `json:"customer_phone,omitempty"`
	CustomerStatus   *string `json:"customer_status,omitempty"`
	CreatedTime      *int64  `json:"created_time,omitempty"`
	LastModifiedTime *int64  `json:"last_modified_time,omitempty"`
}
