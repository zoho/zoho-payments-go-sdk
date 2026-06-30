package zohopayments

import (
	"fmt"

	"github.com/zoho/zoho-payments-go-sdk/redact"
)

func (c Config) String() string {
	return fmt.Sprintf("Config{AccountID: %s, Edition: %s, OAuthToken: %s, ConnectTimeout: %s, RequestTimeout: %s, DefaultHeaders: %v}",
		c.AccountID, c.Edition, redact.Token(c.OAuthToken), c.ConnectTimeout, c.RequestTimeout, c.DefaultHeaders)
}
func (c Config) GoString() string { return c.String() }

func (c *Client) String() string   { return "Client{Edition: " + c.edition.String() + "}" }
func (c *Client) GoString() string { return c.String() }
