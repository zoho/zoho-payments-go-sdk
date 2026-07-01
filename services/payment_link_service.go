package services

import (
	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/exception"
	"github.com/zoho/zoho-payments-go-sdk/internal"
	"github.com/zoho/zoho-payments-go-sdk/model"
	"github.com/zoho/zoho-payments-go-sdk/param"
)

const paymentLinkEnvelope = "payment_links"

// PaymentLinkService provides payment link operations (all editions).
type PaymentLinkService struct {
	caller  *internal.Caller
	edition edition.Edition
}

func NewPaymentLinkService(caller *internal.Caller, ed edition.Edition) *PaymentLinkService {
	return &PaymentLinkService{caller: caller, edition: ed}
}

func (s *PaymentLinkService) Create(params *param.PaymentLinkCreateParams) (*model.PaymentLink, error) {
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
	var result model.PaymentLink
	if err := s.caller.PostInto("/paymentlinks", body, &result, paymentLinkEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PaymentLinkService) Get(paymentLinkID string) (*model.PaymentLink, error) {
	if err := param.Require("paymentLinkId", paymentLinkID); err != nil {
		return nil, err
	}
	var result model.PaymentLink
	path := "/paymentlinks/" + internal.EncodePathSegment(paymentLinkID)
	if err := s.caller.GetInto(path, nil, &result, paymentLinkEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PaymentLinkService) Update(paymentLinkID string, params *param.PaymentLinkUpdateParams) (*model.PaymentLink, error) {
	if err := param.Require("paymentLinkId", paymentLinkID); err != nil {
		return nil, err
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
	var result model.PaymentLink
	path := "/paymentlinks/" + internal.EncodePathSegment(paymentLinkID)
	if err := s.caller.PutInto(path, body, &result, paymentLinkEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PaymentLinkService) Cancel(paymentLinkID string) (*model.PaymentLink, error) {
	if err := param.Require("paymentLinkId", paymentLinkID); err != nil {
		return nil, err
	}
	var result model.PaymentLink
	path := "/paymentlinks/" + internal.EncodePathSegment(paymentLinkID) + "/cancel"
	if err := s.caller.PutInto(path, nil, &result, paymentLinkEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}
