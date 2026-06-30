package zohopayments

import (
	"github.com/zoho/zoho-payments-go-sdk/auth"
	"github.com/zoho/zoho-payments-go-sdk/exception"
	"github.com/zoho/zoho-payments-go-sdk/model"
	znet "github.com/zoho/zoho-payments-go-sdk/net"
	"github.com/zoho/zoho-payments-go-sdk/param"
	"github.com/zoho/zoho-payments-go-sdk/services"
)

type OAuthToken = auth.OAuthToken

type (
	HTTPClient     = znet.HTTPClient
	Request        = znet.Request
	RequestBuilder = znet.RequestBuilder
	Response       = znet.Response
	Method         = znet.Method
)

const (
	MethodGet    = znet.GET
	MethodPost   = znet.POST
	MethodPut    = znet.PUT
	MethodDelete = znet.DELETE
)

func NewResponse(statusCode int, headers map[string][]string, body string) *Response {
	return znet.NewResponse(statusCode, headers, body)
}

func NewRequestBuilder() *RequestBuilder { return znet.NewRequestBuilder() }

type (
	ZohoPaymentsError       = exception.ZohoPaymentsError
	APIError                = exception.APIError
	AuthenticationError     = exception.AuthenticationError
	PermissionError         = exception.PermissionError
	ResourceNotFoundError   = exception.ResourceNotFoundError
	InvalidRequestError     = exception.InvalidRequestError
	RateLimitError          = exception.RateLimitError
	ConnectionError         = exception.ConnectionError
	ValidationError         = exception.ValidationError
	UnsupportedEditionError = exception.UnsupportedEditionError
)

func newValidationError(field, message string) error {
	return exception.NewValidationError(field, message)
}

type (
	PaymentLinkService          = services.PaymentLinkService
	PaymentSessionService       = services.PaymentSessionService
	PaymentService              = services.PaymentService
	RefundService               = services.RefundService
	CustomerService             = services.CustomerService
	PaymentMethodService        = services.PaymentMethodService
	PaymentMethodSessionService = services.PaymentMethodSessionService
	MandateService              = services.MandateService
	CollectService              = services.CollectService
)

type (
	Decimal                     = model.Decimal
	PageContext                 = model.PageContext
	MetaData                    = model.MetaData
	NotifyCustomer              = model.NotifyCustomer
	HostedPageResponse          = model.HostedPageResponse
	Configurations              = model.Configurations
	PaymentLink                 = model.PaymentLink
	PaymentLinkPayment          = model.PaymentLinkPayment
	PaymentSession              = model.PaymentSession
	PaymentSessionPayment       = model.PaymentSessionPayment
	Payment                     = model.Payment
	PaymentSummary              = model.PaymentSummary
	PaymentMethodSummary        = model.PaymentMethodSummary
	Refund                      = model.Refund
	Customer                    = model.Customer
	CustomerSummary             = model.CustomerSummary
	CustomerCard                = model.CustomerCard
	CustomerAchDebit            = model.CustomerAchDebit
	CustomerPaymentMethod       = model.CustomerPaymentMethod
	PaymentMethod               = model.PaymentMethod
	PaymentMethodDetail         = model.PaymentMethodDetail
	PaymentMethodSession        = model.PaymentMethodSession
	PaymentMethodSessionDetail  = model.PaymentMethodSessionDetail
	CardChecks                  = model.CardChecks
	SavedCardDetail             = model.SavedCardDetail
	CardDetail                  = model.CardDetail
	AchDebitDetail              = model.AchDebitDetail
	AddressDetail               = model.AddressDetail
	Upi                         = model.Upi
	NetBanking                  = model.NetBanking
	BankTransfer                = model.BankTransfer
	Mandate                     = model.Mandate
	MandateUpi                  = model.MandateUpi
	MandatePaymentMethod        = model.MandatePaymentMethod
	MandateNotification         = model.MandateNotification
	MandatePayment              = model.MandatePayment
	VirtualAccount              = model.VirtualAccount
	VirtualAccountPayment       = model.VirtualAccountPayment
	VirtualAccountPaymentMethod = model.VirtualAccountPaymentMethod
)

type (
	MetaDataParams                       = param.MetaDataParams
	NotifyCustomerParams                 = param.NotifyCustomerParams
	PostalAddressParams                  = param.PostalAddressParams
	HostedPageParams                     = param.HostedPageParams
	ConfigurationsParams                 = param.ConfigurationsParams
	PaymentLinkConfigurationsParams      = param.PaymentLinkConfigurationsParams
	BrowserInfo                          = param.BrowserInfo
	CardUpdate                           = param.CardUpdate
	AchDebitUpdate                       = param.AchDebitUpdate
	MandateDetailsParams                 = param.MandateDetailsParams
	MandateConfigurationsParams          = param.MandateConfigurationsParams
	PaymentLinkCreateParams              = param.PaymentLinkCreateParams
	PaymentLinkUpdateParams              = param.PaymentLinkUpdateParams
	PaymentSessionCreateParams           = param.PaymentSessionCreateParams
	PaymentCreateParams                  = param.PaymentCreateParams
	PaymentListParams                    = param.PaymentListParams
	RefundCreateParams                   = param.RefundCreateParams
	CustomerCreateParams                 = param.CustomerCreateParams
	CustomerListParams                   = param.CustomerListParams
	PaymentMethodUpdateParams            = param.PaymentMethodUpdateParams
	PaymentMethodSessionCreateParams     = param.PaymentMethodSessionCreateParams
	MandateEnrollmentSessionCreateParams = param.MandateEnrollmentSessionCreateParams
	MandateExecutionSessionCreateParams  = param.MandateExecutionSessionCreateParams
	MandateNotifyParams                  = param.MandateNotifyParams
	MandateExecuteParams                 = param.MandateExecuteParams
	VirtualAccountCreateParams           = param.VirtualAccountCreateParams
	VirtualAccountUpdateParams           = param.VirtualAccountUpdateParams
	VirtualAccountPaymentListParams      = param.VirtualAccountPaymentListParams
)

func String(v string) *string { return param.String(v) }

func Int(v int) *int { return param.Int(v) }

func Int64(v int64) *int64 { return param.Int64(v) }

func Float64(v float64) *float64 { return param.Float64(v) }

func Bool(v bool) *bool { return param.Bool(v) }
