package model

type HostedPageResponse struct {
	Description      *string `json:"description,omitempty"`
	SuccessURL       *string `json:"success_url,omitempty"`
	FailureURL       *string `json:"failure_url,omitempty"`
	Name             *string `json:"name,omitempty"`
	Email            *string `json:"email,omitempty"`
	Phone            *string `json:"phone,omitempty"`
	PhoneCountryCode *string `json:"phone_country_code,omitempty"`
	UDF1             *string `json:"udf1,omitempty"`
	UDF2             *string `json:"udf2,omitempty"`
	UDF3             *string `json:"udf3,omitempty"`
	UDF4             *string `json:"udf4,omitempty"`
	UDF5             *string `json:"udf5,omitempty"`
}
