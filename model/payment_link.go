package model

type PaymentLink struct {
	PaymentLinkID    *string              `json:"payment_link_id,omitempty"`
	URL              *string              `json:"url,omitempty"`
	ExpiresAt        *string              `json:"expires_at,omitempty"`
	Amount           *Decimal             `json:"amount,omitempty"`
	AmountPaid       *Decimal             `json:"amount_paid,omitempty"`
	Currency         *string              `json:"currency,omitempty"`
	Status           *string              `json:"status,omitempty"`
	Email            *string              `json:"email,omitempty"`
	ReferenceID      *string              `json:"reference_id,omitempty"`
	Description      *string              `json:"description,omitempty"`
	ReturnURL        *string              `json:"return_url,omitempty"`
	Phone            *string              `json:"phone,omitempty"`
	PhoneCountryCode *string              `json:"phone_country_code,omitempty"`
	CreatedTime      *int64               `json:"created_time,omitempty"`
	CreatedByID      *string              `json:"created_by_id,omitempty"`
	CreatedBy        *string              `json:"created_by,omitempty"`
	LastModifiedByID *string              `json:"last_modified_by_id,omitempty"`
	LastModified     *string              `json:"last_modified,omitempty"`
	Configurations   *Configurations      `json:"configurations,omitempty"`
	Payments         []PaymentLinkPayment `json:"payments,omitempty"`
}
