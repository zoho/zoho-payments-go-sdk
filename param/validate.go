package param

import (
	"fmt"

	"github.com/zoho/zoho-payments-go-sdk/exception"
)

const (
	maxDescriptionLength   = 500
	maxInvoiceNumberLength = 50
	maxReferenceLength     = 50
	maxMetaDataEntries     = 5
	maxMetaDataKeyLength   = 20
	maxMetaDataValueLength = 500
)

func Require(field, value string) error {
	if value == "" {
		return exception.NewValidationError(field, "is required")
	}
	return nil
}

func validateMaxLen(field, value string, max int) error {
	if len(value) > max {
		return exception.NewValidationError(field, fmt.Sprintf("must be at most %d characters", max))
	}
	return nil
}

func validateMaxLenPtr(field string, value *string, max int) error {
	if value == nil {
		return nil
	}
	return validateMaxLen(field, *value, max)
}

func validateMetaData(metaData []MetaDataParams) error {
	if len(metaData) == 0 {
		return nil
	}
	if len(metaData) > maxMetaDataEntries {
		return exception.NewValidationError("meta_data", fmt.Sprintf("can have at most %d entries", maxMetaDataEntries))
	}
	for _, entry := range metaData {
		if entry.Key == "" {
			return exception.NewValidationError("meta_data", "key must not be empty")
		}
		if len(entry.Key) > maxMetaDataKeyLength {
			return exception.NewValidationError("meta_data", fmt.Sprintf("key must be at most %d characters", maxMetaDataKeyLength))
		}
		if len(entry.Value) > maxMetaDataValueLength {
			return exception.NewValidationError("meta_data", fmt.Sprintf("value must be at most %d characters", maxMetaDataValueLength))
		}
	}
	return nil
}
