package services

import (
	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/exception"
	"github.com/zoho/zoho-payments-go-sdk/internal"
	"github.com/zoho/zoho-payments-go-sdk/model"
	"github.com/zoho/zoho-payments-go-sdk/param"
)

const paymentMethodSessionEnvelope = "payment_method_session"

// PaymentMethodSessionService provides payment method session operations (US only).
type PaymentMethodSessionService struct {
	caller  *internal.Caller
	edition edition.Edition
}

func NewPaymentMethodSessionService(c *internal.Caller, ed edition.Edition) *PaymentMethodSessionService {
	return &PaymentMethodSessionService{caller: c, edition: ed}
}

func (s *PaymentMethodSessionService) Create(params *param.PaymentMethodSessionCreateParams) (*model.PaymentMethodSession, error) {
	if !s.edition.IsUS() {
		return nil, exception.NewUnsupportedEditionError("paymentMethodSessions.Create", s.edition)
	}
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
	var result model.PaymentMethodSession
	if err := s.caller.PostInto("/paymentmethodsessions", body, &result, paymentMethodSessionEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PaymentMethodSessionService) Get(paymentMethodSessionID string) (*model.PaymentMethodSession, error) {
	if !s.edition.IsUS() {
		return nil, exception.NewUnsupportedEditionError("paymentMethodSessions.Get", s.edition)
	}
	if err := param.Require("paymentMethodSessionId", paymentMethodSessionID); err != nil {
		return nil, err
	}
	var result model.PaymentMethodSession
	path := "/paymentmethodsessions/" + internal.EncodePathSegment(paymentMethodSessionID)
	if err := s.caller.GetInto(path, nil, &result, paymentMethodSessionEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}
