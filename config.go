package zohopayments

import "time"

// Config holds the settings used to construct a [Client].
type Config struct {
	AccountID      string
	Edition        Edition
	OAuthToken     string
	ConnectTimeout time.Duration
	RequestTimeout time.Duration
	HTTPClient     HTTPClient
	DefaultHeaders map[string]string
}
