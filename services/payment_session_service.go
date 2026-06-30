package services

import (
	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/exception"
	"github.com/zoho/zoho-payments-go-sdk/internal"
	"github.com/zoho/zoho-payments-go-sdk/model"
	"github.com/zoho/zoho-payments-go-sdk/param"
)

const paymentSessionEnvelope = "payments_session"

// PaymentSessionService provides payment session operations (all editions).
type PaymentSessionService struct {
	caller  *internal.Caller
	edition edition.Edition
}

func NewPaymentSessionService(c *internal.Caller, ed edition.Edition) *PaymentSessionService {
	return &PaymentSessionService{caller: c, edition: ed}
}

func (s *PaymentSessionService) Create(params *param.PaymentSessionCreateParams) (*model.PaymentSession, error) {
	if params == nil {
		return nil, exception.NewValidationError("", "params must not be nil")
	}
	if err := params.Validate(); err != nil {
		return nil, err
	}
	body, err := internal.MarshalBody(params.ToBody())
	if err != nil {
		return nil, err
	}
	var result model.PaymentSession
	if err := s.caller.PostInto("/paymentsessions", body, &result, paymentSessionEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PaymentSessionService) Get(paymentSessionID string) (*model.PaymentSession, error) {
	if err := param.Require("paymentSessionId", paymentSessionID); err != nil {
		return nil, err
	}
	var result model.PaymentSession
	path := "/paymentsessions/" + internal.EncodePathSegment(paymentSessionID)
	if err := s.caller.GetInto(path, nil, &result, paymentSessionEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}
