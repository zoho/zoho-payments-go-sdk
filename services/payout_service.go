package services

import (
	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/internal"
	"github.com/zoho/zoho-payments-go-sdk/model"
	"github.com/zoho/zoho-payments-go-sdk/param"
)

const (
	payoutEnvelope          = "payout"
	payoutListEnvelope      = "payouts"
	transactionListEnvelope = "transactions"
)

// PayoutService provides payout operations (all editions).
type PayoutService struct {
	caller  *internal.Caller
	edition edition.Edition
}

func NewPayoutService(caller *internal.Caller, ed edition.Edition) *PayoutService {
	return &PayoutService{caller: caller, edition: ed}
}

func (s *PayoutService) Get(payoutID string) (*model.PayoutDetail, error) {
	if err := param.Require("payoutId", payoutID); err != nil {
		return nil, err
	}
	var result model.PayoutDetail
	path := "/payouts/" + internal.EncodePathSegment(payoutID)
	if err := s.caller.GetInto(path, nil, &result, payoutEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *PayoutService) List(params *param.PayoutListParams) (*model.ListResponse[model.Payout], error) {
	return internal.DoList[model.Payout](s.caller, "/payouts", params.ToQuery(), payoutListEnvelope)
}

func (s *PayoutService) ListTransactions(payoutID string,
	params *param.PayoutTransactionListParams) (*model.ListResponse[model.PayoutTransaction], error) {
	if err := param.Require("payoutId", payoutID); err != nil {
		return nil, err
	}
	path := "/payouts/" + internal.EncodePathSegment(payoutID) + "/transactions"
	return internal.DoList[model.PayoutTransaction](s.caller, path, params.ToQuery(), transactionListEnvelope)
}
