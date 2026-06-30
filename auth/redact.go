package auth

import (
	"fmt"

	"github.com/zoho/zoho-payments-go-sdk/redact"
)

func (t OAuthToken) String() string {
	return fmt.Sprintf("OAuthToken{AccessToken: %s, ExpiresIn: %d}", redact.Token(t.AccessToken), t.ExpiresIn)
}

func (t OAuthToken) GoString() string { return t.String() }
