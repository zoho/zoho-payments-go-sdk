package zohopayments

import (
	"github.com/zoho/zoho-payments-go-sdk/internal"
	znet "github.com/zoho/zoho-payments-go-sdk/net"
	"github.com/zoho/zoho-payments-go-sdk/services"
)

// Client is a Zoho Payments API client.
type Client struct {
	caller    *internal.Caller
	tokens    *internal.TokenManager
	edition   Edition
	transport znet.HTTPClient

	paymentLinks          *services.PaymentLinkService
	paymentSessions       *services.PaymentSessionService
	customers             *services.CustomerService
	payments              *services.PaymentService
	refunds               *services.RefundService
	paymentMethods        *services.PaymentMethodService
	paymentMethodSessions *services.PaymentMethodSessionService
	mandates              *services.MandateService
	collect               *services.CollectService
}

func (c *Client) Edition() Edition { return c.edition }

func (c *Client) PaymentLinks() *services.PaymentLinkService { return c.paymentLinks }

func (c *Client) PaymentSessions() *services.PaymentSessionService { return c.paymentSessions }

func (c *Client) Customers() *services.CustomerService { return c.customers }

func (c *Client) Payments() *services.PaymentService { return c.payments }

func (c *Client) Refunds() *services.RefundService { return c.refunds }

func (c *Client) PaymentMethods() *services.PaymentMethodService { return c.paymentMethods }

func (c *Client) PaymentMethodSessions() *services.PaymentMethodSessionService {
	return c.paymentMethodSessions
}

func (c *Client) Mandates() *services.MandateService { return c.mandates }

func (c *Client) Collect() *services.CollectService { return c.collect }

func (c *Client) UpdateToken(newAccessToken string) {
	c.tokens.Update(newAccessToken)
}

func (c *Client) Close() error {
	if closer, ok := c.transport.(interface{ Close() error }); ok {
		return closer.Close()
	}
	return nil
}
