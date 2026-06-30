# Zoho Payments Go SDK

Official Go SDK for the Zoho Payments API — supports IN, IN Sandbox, and US editions.

API reference:
- **India:** https://www.zoho.com/in/payments/api/v1/introduction/
- **United States:** https://www.zoho.com/us/payments/api/v1/introduction/

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

## Requirements

- **Go 1.21+**
- No third-party dependencies (standard library only).

## Installation

```sh
go get github.com/zoho/zoho-payments-go-sdk
```

```go
import zohopayments "github.com/zoho/zoho-payments-go-sdk"
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	zohopayments "github.com/zoho/zoho-payments-go-sdk"
)

func main() {
	// 1. Build the client.
	client, err := zohopayments.NewClient(&zohopayments.Config{
		AccountID:  "23137556",
		Edition:    zohopayments.EditionIN,
		OAuthToken: "1000.xxxx.yyyy",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 2. Use a service.
	link, err := client.PaymentLinks().Create(&zohopayments.PaymentLinkCreateParams{
		Amount:      500.00,
		Currency:    "INR",
		Description: "Order #1234",
		Email:       zohopayments.String("customer@example.com"),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created:", *link.PaymentLinkID)
}
```

Optional fields use pointers; the helpers `zohopayments.String`, `Int`, `Int64`,
`Float64`, and `Bool` return a pointer to a value.

## Editions

| Edition | Payments API Base URL | OAuth Accounts URL |
|---------|-----------------------|--------------------|
| `EditionIN` | `https://payments.zoho.in/api/v1` | `https://accounts.zoho.in` |
| `EditionINSandbox` | `https://paymentssandbox.zoho.in/api/v1` | `https://accounts.zoho.in` |
| `EditionUS` | `https://payments.zoho.com/api/v1` | `https://accounts.zoho.com` |

Helper methods: `edition.IsIN()` returns `true` for both `EditionIN` and
`EditionINSandbox`; `edition.IsUS()` returns `true` for `EditionUS`.
`ParseEdition("IN")` parses an edition name (case-insensitive).

## Authentication

### Access token only

```go
client, err := zohopayments.NewClient(&zohopayments.Config{
	AccountID:  "23137556",
	Edition:    zohopayments.EditionIN,
	OAuthToken: "1000.access_token_here",
})
```

### Token refresh

The SDK does **not** auto-refresh tokens. Use `GenerateAccessToken` when the
access token expires, then push the new token into the client:

```go
fresh, err := zohopayments.GenerateAccessToken(
	"refresh_token",
	"client_id",
	"client_secret",
	"redirect_uri",
	zohopayments.EditionIN,
)
if err != nil {
	log.Fatal(err)
}

// Persist the new token in your storage layer, then update the running
// client (thread-safe, no rebuild needed).
client.UpdateToken(fresh.AccessToken)
```

`OAuthToken` exposes `AccessToken` and `ExpiresIn` (token lifetime in seconds,
as returned by IAM).

## Client configuration

```go
client, err := zohopayments.NewClient(&zohopayments.Config{
	AccountID:      "23137556",                       // Required
	Edition:        zohopayments.EditionIN,           // Required
	OAuthToken:     "1000.xxxx.yyyy",                 // Required
	ConnectTimeout: 15 * time.Second,                 // Default: 30s
	RequestTimeout: 45 * time.Second,                 // Default: 60s
	DefaultHeaders: map[string]string{"X-Custom-Header": "value"},
})
```

Reserved headers (`Authorization`, `User-Agent`, `Accept`, `Content-Type`,
`Content-Length`, `Host`) are managed by the SDK and cannot be overridden via
`DefaultHeaders`.

## Services

| Accessor | Description | Editions |
|----------|-------------|----------|
| `client.PaymentLinks()` | Payment link CRUD | All |
| `client.PaymentSessions()` | Payment sessions | All |
| `client.Customers()` | Customers | All (List/Delete: US only) |
| `client.Payments()` | Payments | All (Create: US only) |
| `client.Refunds()` | Refunds | All |
| `client.PaymentMethods()` | Saved payment methods | US only |
| `client.PaymentMethodSessions()` | Payment-method collection sessions | US only |
| `client.Mandates()` | Recurring mandates | IN only |
| `client.Collect()` | Virtual accounts (Collect) | IN only |

Calling an edition-gated operation on the wrong edition returns an
`*UnsupportedEditionError`.

## Examples

### Payment link

```go
link, err := client.PaymentLinks().Create(&zohopayments.PaymentLinkCreateParams{
	Amount:      500.00,
	Currency:    "INR",
	Description: "Order #1234",
	Email:       zohopayments.String("customer@example.com"),
	NotifyCustomer: &zohopayments.NotifyCustomerParams{
		Email: zohopayments.Bool(true),
		SMS:   zohopayments.Bool(false),
	},
})

fetched, err := client.PaymentLinks().Get(*link.PaymentLinkID)
cancelled, err := client.PaymentLinks().Cancel(*link.PaymentLinkID)
```

### Customer

```go
customer, err := client.Customers().Create(&zohopayments.CustomerCreateParams{
	Name:     "Jane Doe",
	Email:    "jane@example.com",
	MetaData: []zohopayments.MetaDataParams{{Key: "source", Value: "web"}},
})
```

### Refund

```go
refund, err := client.Refunds().Create("pay_abc", &zohopayments.RefundCreateParams{
	Amount: 100.00,
	Reason: "requested_by_customer",
	Type:   "full",
})
```

### Mandate (IN)

```go
enrollment, err := client.Mandates().CreateEnrollmentSession(
	&zohopayments.MandateEnrollmentSessionCreateParams{
		Amount:      0.00,
		Currency:    "INR",
		CustomerID:  "cust_abc",
		Description: "SIP enrollment",
		MandateDetails: &zohopayments.MandateDetailsParams{
			PaymentMethodType: "upi",
			Frequency:         "monthly",
			Description:       "Monthly SIP",
			AmountRule:        "variable",
			MaxAmount:         zohopayments.Float64(5000.00),
		},
	},
)
```

## Error handling

Every API error is one of the typed errors below; match them with `errors.As`.
The specific types embed `*APIError`, so matching against `*APIError` catches
any of them.

| Error | HTTP |
|-------|------|
| `*AuthenticationError` | 401 |
| `*PermissionError` | 403 |
| `*ResourceNotFoundError` | 404 |
| `*InvalidRequestError` | 400, 422 |
| `*RateLimitError` | 429 |
| `*APIError` | Any other non-2xx |
| `*ConnectionError` | Network / I/O failure |
| `*ValidationError` | Client-side parameter validation (request not sent) |
| `*UnsupportedEditionError` | Operation not available for the client's edition |

```go
_, err := client.Payments().Get("pay_missing")

var authErr *zohopayments.AuthenticationError
var invalidErr *zohopayments.InvalidRequestError
var apiErr *zohopayments.APIError
switch {
case errors.As(err, &authErr):
	// refresh the token and retry
case errors.As(err, &invalidErr):
	fmt.Println(invalidErr.Code, invalidErr.Message)
case errors.As(err, &apiErr):
	fmt.Println("other API error:", apiErr.HTTPStatusCode)
}
```

## Custom HTTP transport

`DefaultHTTPClient` wraps `net/http`. To plug in your own retries, proxy, or
instrumentation, implement the `HTTPClient` interface:

```go
type MyTransport struct{}

func (MyTransport) Execute(req *zohopayments.Request) (*zohopayments.Response, error) {
	// ... send req.Method() / req.URL() / req.Headers() / req.Body()
	return zohopayments.NewResponse(200, nil, `{...}`), nil
}

client, err := zohopayments.NewClient(&zohopayments.Config{
	AccountID:  "...",
	Edition:    zohopayments.EditionIN,
	OAuthToken: "...",
	HTTPClient: MyTransport{},
})
```

When you inject a custom transport you cannot also set `ConnectTimeout` — the
transport manages its own connection lifecycle. If your transport implements
`io.Closer`, `client.Close()` will call it.

## Thread safety

- `Client` is safe to share across goroutines.
- `client.UpdateToken()` is atomic; in-flight requests on other goroutines see
  either the old or the new token.
- Services are constructed eagerly and are stateless beyond the shared client.

## License

Apache 2.0
