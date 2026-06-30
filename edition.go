package zohopayments

import "github.com/zoho/zoho-payments-go-sdk/edition"

// Edition identifies the Zoho Payments data center / environment a client targets.
type Edition = edition.Edition

const (
	EditionIN = edition.IN

	EditionINSandbox = edition.INSandbox

	EditionUS = edition.US
)

func ParseEdition(name string) (Edition, error) {
	return edition.Parse(name)
}
