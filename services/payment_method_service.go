package services

import (
	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/exception"
	"github.com/zoho/zoho-payments-go-sdk/internal"
	"github.com/zoho/zoho-payments-go-sdk/model"
	"github.com/zoho/zoho-payments-go-sdk/param"
)

const paymentMethodEnvelope = "payment_method"

// PaymentMethodService provides saved payment method operations (US only).
type PaymentMethodService struct {
	caller  *internal.Caller
	edition edition.Edition
}

func NewPaymentMethodService(c *internal.Caller, ed edition.Edition) *PaymentMethodService {
	return &PaymentMethodService{caller: c, edition: ed}
}

func (s *PaymentMethodService) Get(paymentMethodID string) (*model.PaymentMethod, error) {
	if !s.edition.IsUS() {
		return nil, exception.NewUnsupportedEditionError("paymentMethods.Get", s.edition)
	}
	if err := param.Require("paymentMethodId", paymentMethodID); err != nil {
		return nil, err
	}
	var result model.PaymentMethod
	path := "/paymentmethods/" + internal.EncodePathSegment(paymentMethodID)
	if err := s.caller.GetInto(path, nil, &result, paymentMethodEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PaymentMethodService) Update(paymentMethodID string, params *param.PaymentMethodUpdateParams) (*model.PaymentMethod, error) {
	if !s.edition.IsUS() {
		return nil, exception.NewUnsupportedEditionError("paymentMethods.Update", s.edition)
	}
	if err := param.Require("paymentMethodId", paymentMethodID); err != nil {
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
	var result model.PaymentMethod
	path := "/paymentmethods/" + internal.EncodePathSegment(paymentMethodID)
	if err := s.caller.PutInto(path, body, &result, paymentMethodEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PaymentMethodService) Delete(paymentMethodID string) error {
	if !s.edition.IsUS() {
		return exception.NewUnsupportedEditionError("paymentMethods.Delete", s.edition)
	}
	if err := param.Require("paymentMethodId", paymentMethodID); err != nil {
		return err
	}
	return s.caller.DeleteResource("/paymentmethods/" + internal.EncodePathSegment(paymentMethodID))
}
