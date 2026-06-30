package services

import (
	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/exception"
	"github.com/zoho/zoho-payments-go-sdk/internal"
	"github.com/zoho/zoho-payments-go-sdk/model"
	"github.com/zoho/zoho-payments-go-sdk/param"
)

const (
	paymentEnvelope     = "payment"
	paymentListEnvelope = "payments"
)

// PaymentService provides payment operations.
type PaymentService struct {
	caller  *internal.Caller
	edition edition.Edition
}

func NewPaymentService(c *internal.Caller, ed edition.Edition) *PaymentService {
	return &PaymentService{caller: c, edition: ed}
}

func (s *PaymentService) Create(params *param.PaymentCreateParams) (*model.Payment, error) {
	if !s.edition.IsUS() {
		return nil, exception.NewUnsupportedEditionError("payments.Create", s.edition)
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
	var result model.Payment
	if err := s.caller.PostInto("/payments", body, &result, paymentEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PaymentService) Get(paymentID string) (*model.Payment, error) {
	if err := param.Require("paymentId", paymentID); err != nil {
		return nil, err
	}
	var result model.Payment
	path := "/payments/" + internal.EncodePathSegment(paymentID)
	if err := s.caller.GetInto(path, nil, &result, paymentEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PaymentService) List(params *param.PaymentListParams) (*model.ListResponse[model.PaymentSummary], error) {
	return internal.DoList[model.PaymentSummary](s.caller, "/payments", params.ToQuery(), paymentListEnvelope)
}
