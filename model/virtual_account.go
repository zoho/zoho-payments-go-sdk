package model

type VirtualAccount struct {
	VirtualAccountID *string    `json:"virtual_account_id,omitempty"`
	AccountNumber    *string    `json:"account_number,omitempty"`
	IfscCode         *string    `json:"ifsc_code,omitempty"`
	BeneficiaryName  *string    `json:"beneficiary_name,omitempty"`
	Description      *string    `json:"description,omitempty"`
	CustomerID       *string    `json:"customer_id,omitempty"`
	ReferenceNumber  *string    `json:"reference_number,omitempty"`
	Status           *string    `json:"status,omitempty"`
	ExpiresAt        *string    `json:"expires_at,omitempty"`
	CreatedTime      *int64     `json:"created_time,omitempty"`
	LastModifiedTime *int64     `json:"last_modified_time,omitempty"`
	MetaData         []MetaData `json:"meta_data,omitempty"`
	MinimumAmount    *Decimal   `json:"minimum_amount,omitempty"`
	MaximumAmount    *Decimal   `json:"maximum_amount,omitempty"`
	AmountPaid       *Decimal   `json:"amount_paid,omitempty"`
}
