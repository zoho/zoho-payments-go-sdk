package services

import (
	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/exception"
	"github.com/zoho/zoho-payments-go-sdk/internal"
	"github.com/zoho/zoho-payments-go-sdk/model"
	"github.com/zoho/zoho-payments-go-sdk/param"
)

const refundEnvelope = "refund"

// RefundService provides refund operations (all editions).
type RefundService struct {
	caller  *internal.Caller
	edition edition.Edition
}

func NewRefundService(c *internal.Caller, ed edition.Edition) *RefundService {
	return &RefundService{caller: c, edition: ed}
}

func (s *RefundService) Create(paymentID string, params *param.RefundCreateParams) (*model.Refund, error) {
	if err := param.Require("paymentId", paymentID); err != nil {
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
	var result model.Refund
	path := "/payments/" + internal.EncodePathSegment(paymentID) + "/refunds"
	if err := s.caller.PostInto(path, body, &result, refundEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *RefundService) Get(refundID string) (*model.Refund, error) {
	if err := param.Require("refundId", refundID); err != nil {
		return nil, err
	}
	var result model.Refund
	path := "/refunds/" + internal.EncodePathSegment(refundID)
	if err := s.caller.GetInto(path, nil, &result, refundEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}
