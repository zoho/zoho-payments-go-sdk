package zohopayments

import (
	"fmt"
	"strings"

	"github.com/zoho/zoho-payments-go-sdk/internal"
	znet "github.com/zoho/zoho-payments-go-sdk/net"
	"github.com/zoho/zoho-payments-go-sdk/services"
)

var reservedHeaders = map[string]bool{
	"authorization":  true,
	"user-agent":     true,
	"accept":         true,
	"content-type":   true,
	"content-length": true,
	"host":           true,
}

func NewClient(cfg *Config) (*Client, error) {
	if cfg == nil {
		return nil, newValidationError("", "config must not be nil")
	}
	if cfg.AccountID == "" {
		return nil, newValidationError("AccountID", "is required")
	}
	if !cfg.Edition.Valid() {
		return nil, newValidationError("Edition", "a valid edition is required")
	}
	if cfg.OAuthToken == "" {
		return nil, newValidationError("OAuthToken", "is required")
	}
	if cfg.HTTPClient != nil && cfg.ConnectTimeout != 0 {
		return nil, newValidationError("ConnectTimeout", "cannot be set together with a custom HTTPClient")
	}
	for name := range cfg.DefaultHeaders {
		if reservedHeaders[strings.ToLower(name)] {
			return nil, newValidationError("DefaultHeaders",
				fmt.Sprintf("header %q is managed by the SDK and cannot be overridden", name))
		}
	}

	connectTimeout := cfg.ConnectTimeout
	if connectTimeout <= 0 {
		connectTimeout = internal.DefaultConnectTimeout
	}
	requestTimeout := cfg.RequestTimeout
	if requestTimeout <= 0 {
		requestTimeout = internal.DefaultRequestTimeout
	}

	transport := cfg.HTTPClient
	if transport == nil {
		transport = znet.NewDefaultHTTPClient(connectTimeout, requestTimeout)
	}

	headers := make(map[string]string, len(cfg.DefaultHeaders))
	for name, value := range cfg.DefaultHeaders {
		headers[name] = value
	}

	tokens := internal.NewTokenManager(cfg.OAuthToken)
	caller := internal.NewCaller(transport, tokens, cfg.Edition, cfg.AccountID, requestTimeout, headers)

	client := &Client{
		caller:    caller,
		tokens:    tokens,
		edition:   cfg.Edition,
		transport: transport,
	}
	client.paymentLinks = services.NewPaymentLinkService(caller, cfg.Edition)
	client.paymentSessions = services.NewPaymentSessionService(caller, cfg.Edition)
	client.customers = services.NewCustomerService(caller, cfg.Edition)
	client.payments = services.NewPaymentService(caller, cfg.Edition)
	client.refunds = services.NewRefundService(caller, cfg.Edition)
	client.paymentMethods = services.NewPaymentMethodService(caller, cfg.Edition)
	client.paymentMethodSessions = services.NewPaymentMethodSessionService(caller, cfg.Edition)
	client.mandates = services.NewMandateService(caller, cfg.Edition)
	client.collect = services.NewCollectService(caller, cfg.Edition)

	return client, nil
}

func GenerateAccessToken(refreshToken, clientID, clientSecret, redirectURI string, edition Edition) (*OAuthToken, error) {
	return internal.GenerateAccessToken(refreshToken, clientID, clientSecret, redirectURI, edition)
}
