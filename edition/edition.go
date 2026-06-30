// Package edition defines the Zoho Payments deployment editions and their base URLs.
package edition

import (
	"fmt"
	"strings"
)

// Edition identifies the Zoho Payments data center / environment a client targets.
type Edition int

const (
	IN Edition = iota + 1

	INSandbox

	US
)

type info struct {
	name        string
	baseURL     string
	accountsURL string
}

var data = map[Edition]info{
	IN:        {"IN", "https://payments.zoho.in/api/v1", "https://accounts.zoho.in"},
	INSandbox: {"IN_SANDBOX", "https://paymentssandbox.zoho.in/api/v1", "https://accounts.zoho.in"},
	US:        {"US", "https://payments.zoho.com/api/v1", "https://accounts.zoho.com"},
}

func (e Edition) BaseURL() string { return data[e].baseURL }

func (e Edition) AccountsURL() string { return data[e].accountsURL }

func (e Edition) String() string {
	if i, ok := data[e]; ok {
		return i.name
	}
	return fmt.Sprintf("Edition(%d)", int(e))
}

func (e Edition) IsUS() bool { return e == US }

func (e Edition) IsIN() bool { return e == IN || e == INSandbox }

func (e Edition) Valid() bool {
	_, ok := data[e]
	return ok
}

func Parse(name string) (Edition, error) {
	switch strings.ToUpper(strings.TrimSpace(name)) {
	case "IN":
		return IN, nil
	case "IN_SANDBOX":
		return INSandbox, nil
	case "US":
		return US, nil
	default:
		return 0, fmt.Errorf("zohopayments: unknown edition %q", name)
	}
}
