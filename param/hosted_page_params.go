package param

// HostedPageParams configures a hosted payment page.
type HostedPageParams struct {
	Description      string
	SuccessURL       string
	FailureURL       string
	Name             *string
	Email            *string
	Phone            *string
	PhoneCountryCode *string
	UDF1             *string
	UDF2             *string
	UDF3             *string
	UDF4             *string
	UDF5             *string
}

func (h *HostedPageParams) validate() error {
	if err := Require("description", h.Description); err != nil {
		return err
	}
	if err := Require("success_url", h.SuccessURL); err != nil {
		return err
	}
	return Require("failure_url", h.FailureURL)
}

func (h *HostedPageParams) toMap() map[string]any {
	body := map[string]any{
		"description": h.Description,
		"success_url": h.SuccessURL,
		"failure_url": h.FailureURL,
	}
	putStr(body, "name", h.Name)
	putStr(body, "email", h.Email)
	putStr(body, "phone", h.Phone)
	putStr(body, "phone_country_code", h.PhoneCountryCode)
	putStr(body, "udf1", h.UDF1)
	putStr(body, "udf2", h.UDF2)
	putStr(body, "udf3", h.UDF3)
	putStr(body, "udf4", h.UDF4)
	putStr(body, "udf5", h.UDF5)
	return body
}
