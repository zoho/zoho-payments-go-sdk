package services

import (
	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/exception"
	"github.com/zoho/zoho-payments-go-sdk/internal"
	"github.com/zoho/zoho-payments-go-sdk/model"
	"github.com/zoho/zoho-payments-go-sdk/param"
)

const (
	virtualAccountEnvelope         = "virtual_account"
	virtualAccountPaymentsEnvelope = "payments"
)

// CollectService provides virtual account (Collect) operations (IN editions only).
type CollectService struct {
	caller  *internal.Caller
	edition edition.Edition
}

func NewCollectService(c *internal.Caller, ed edition.Edition) *CollectService {
	return &CollectService{caller: c, edition: ed}
}

func (s *CollectService) Create(params *param.VirtualAccountCreateParams) (*model.VirtualAccount, error) {
	if !s.edition.IsIN() {
		return nil, exception.NewUnsupportedEditionError("collect.Create", s.edition)
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
	var result model.VirtualAccount
	if err := s.caller.PostInto("/virtualaccounts", body, &result, virtualAccountEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *CollectService) Update(virtualAccountID string, params *param.VirtualAccountUpdateParams) (*model.VirtualAccount, error) {
	if !s.edition.IsIN() {
		return nil, exception.NewUnsupportedEditionError("collect.Update", s.edition)
	}
	if err := param.Require("virtualAccountId", virtualAccountID); err != nil {
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
	var result model.VirtualAccount
	path := "/virtualaccounts/" + internal.EncodePathSegment(virtualAccountID)
	if err := s.caller.PutInto(path, body, &result, virtualAccountEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *CollectService) Get(virtualAccountID string) (*model.VirtualAccount, error) {
	if !s.edition.IsIN() {
		return nil, exception.NewUnsupportedEditionError("collect.Get", s.edition)
	}
	if err := param.Require("virtualAccountId", virtualAccountID); err != nil {
		return nil, err
	}
	var result model.VirtualAccount
	path := "/virtualaccounts/" + internal.EncodePathSegment(virtualAccountID)
	if err := s.caller.GetInto(path, nil, &result, virtualAccountEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *CollectService) ListPayments(virtualAccountID string, params *param.VirtualAccountPaymentListParams) (*model.ListResponse[model.VirtualAccountPayment], error) {
	if !s.edition.IsIN() {
		return nil, exception.NewUnsupportedEditionError("collect.ListPayments", s.edition)
	}
	if err := param.Require("virtualAccountId", virtualAccountID); err != nil {
		return nil, err
	}
	path := "/virtualaccounts/" + internal.EncodePathSegment(virtualAccountID) + "/payments"
	return internal.DoList[model.VirtualAccountPayment](s.caller, path, params.ToQuery(), virtualAccountPaymentsEnvelope)
}

func (s *CollectService) Close(virtualAccountID string) error {
	if !s.edition.IsIN() {
		return exception.NewUnsupportedEditionError("collect.Close", s.edition)
	}
	if err := param.Require("virtualAccountId", virtualAccountID); err != nil {
		return err
	}
	return s.caller.Put("/virtualaccounts/" + internal.EncodePathSegment(virtualAccountID) + "/close")
}
