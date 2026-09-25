package services

import (
	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/exception"
	"github.com/zoho/zoho-payments-go-sdk/internal"
	"github.com/zoho/zoho-payments-go-sdk/model"
	"github.com/zoho/zoho-payments-go-sdk/param"
)

const (
	transferDetailEnvelope         = "transfer_details"
	transferListEnvelope           = "transfers"
	transferReversalDetailEnvelope = "transfer_reversal_details"
	transferReversalListEnvelope   = "transfer_reversals"
	transferReversalEnvelope       = "transfer_reversal"
	dataEnvelope                   = "data"
	connectedAccountEnvelope       = "connected_account"
	connectedAccountListEnvelope   = "connected_accounts"
)

// SplitSettlementService provides transfer, transfer reversal and connected account operations (IN editions only).
type SplitSettlementService struct {
	caller  *internal.Caller
	edition edition.Edition
}

func NewSplitSettlementService(caller *internal.Caller, ed edition.Edition) *SplitSettlementService {
	return &SplitSettlementService{caller: caller, edition: ed}
}

// supported reports whether split settlement is available on this edition.
// It is an IN-only API.
func (s *SplitSettlementService) supported() bool {
	return s.edition.IsIN()
}

// CreateTransfer splits a payment across one or more connected accounts.
func (s *SplitSettlementService) CreateTransfer(params *param.TransferCreateParams) (*model.TransferCreateResponse, error) {
	if !s.supported() {
		return nil, exception.NewUnsupportedEditionError("splitSettlement.CreateTransfer", s.edition)
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
	var result model.TransferCreateResponse
	if err := s.caller.PostInto("/transfers", body, &result, dataEnvelope, transferListEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SplitSettlementService) GetTransfer(transferID string) (*model.Transfer, error) {
	if !s.supported() {
		return nil, exception.NewUnsupportedEditionError("splitSettlement.GetTransfer", s.edition)
	}
	if err := param.Require("transferId", transferID); err != nil {
		return nil, err
	}
	var result model.Transfer
	path := "/transfers/" + internal.EncodePathSegment(transferID)
	if err := s.caller.GetInto(path, nil, &result, transferDetailEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SplitSettlementService) ListTransfers(params *param.TransferListParams) (*model.ListResponse[model.TransferSummary], error) {
	if !s.supported() {
		return nil, exception.NewUnsupportedEditionError("splitSettlement.ListTransfers", s.edition)
	}
	return internal.DoList[model.TransferSummary](s.caller, "/transfers", params.ToQuery(), transferListEnvelope)
}

// CreateTransferReversal reverses all or part of a transfer.
func (s *SplitSettlementService) CreateTransferReversal(
	params *param.TransferReversalCreateParams) (*model.TransferReversalCreateResponse, error) {
	if !s.supported() {
		return nil, exception.NewUnsupportedEditionError("splitSettlement.CreateTransferReversal", s.edition)
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
	var result model.TransferReversalCreateResponse
	if err := s.caller.PostInto("/transferreversals", body, &result, dataEnvelope, transferReversalEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SplitSettlementService) GetTransferReversal(transferReversalID string) (*model.TransferReversalDetail, error) {
	if !s.supported() {
		return nil, exception.NewUnsupportedEditionError("splitSettlement.GetTransferReversal", s.edition)
	}
	if err := param.Require("transferReversalId", transferReversalID); err != nil {
		return nil, err
	}
	var result model.TransferReversalDetail
	path := "/transferreversals/" + internal.EncodePathSegment(transferReversalID)
	if err := s.caller.GetInto(path, nil, &result, transferReversalDetailEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SplitSettlementService) ListTransferReversals(
	params *param.TransferReversalListParams) (*model.ListResponse[model.TransferReversal], error) {
	if !s.supported() {
		return nil, exception.NewUnsupportedEditionError("splitSettlement.ListTransferReversals", s.edition)
	}
	return internal.DoList[model.TransferReversal](s.caller, "/transferreversals", params.ToQuery(), transferReversalListEnvelope)
}

// CreateConnectedAccount onboards a connected account. The API returns no resource body.
func (s *SplitSettlementService) CreateConnectedAccount(params *param.ConnectedAccountCreateParams) error {
	if !s.supported() {
		return exception.NewUnsupportedEditionError("splitSettlement.CreateConnectedAccount", s.edition)
	}
	if params == nil {
		return exception.NewValidationError("", "params must not be nil")
	}
	if err := params.Validate(); err != nil {
		return err
	}
	body, err := internal.MarshalBody(params.ToBody())
	if err != nil {
		return err
	}
	return s.caller.Post("/connectedaccounts", body)
}

func (s *SplitSettlementService) GetConnectedAccount(connectedAccountID string) (*model.ConnectedAccount, error) {
	if !s.supported() {
		return nil, exception.NewUnsupportedEditionError("splitSettlement.GetConnectedAccount", s.edition)
	}
	if err := param.Require("connectedAccountId", connectedAccountID); err != nil {
		return nil, err
	}
	var result model.ConnectedAccount
	path := "/connectedaccounts/" + internal.EncodePathSegment(connectedAccountID)
	if err := s.caller.GetInto(path, nil, &result, connectedAccountEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SplitSettlementService) ListConnectedAccounts(
	params *param.ConnectedAccountListParams) (*model.ListResponse[model.ConnectedAccountSummary], error) {
	if !s.supported() {
		return nil, exception.NewUnsupportedEditionError("splitSettlement.ListConnectedAccounts", s.edition)
	}
	return internal.DoList[model.ConnectedAccountSummary](s.caller, "/connectedaccounts", params.ToQuery(),
		connectedAccountListEnvelope)
}

func (s *SplitSettlementService) ListConnectedAccountPayouts(
	connectedAccountID string) (*model.ListResponse[model.ConnectedAccountPayoutSummary], error) {
	if !s.supported() {
		return nil, exception.NewUnsupportedEditionError("splitSettlement.ListConnectedAccountPayouts", s.edition)
	}
	if err := param.Require("connectedAccountId", connectedAccountID); err != nil {
		return nil, err
	}
	path := "/connectedaccounts/" + internal.EncodePathSegment(connectedAccountID) + "/payouts"
	return internal.DoList[model.ConnectedAccountPayoutSummary](s.caller, path, nil, payoutListEnvelope)
}

func (s *SplitSettlementService) GetConnectedAccountPayout(connectedAccountID,
	payoutID string) (*model.ConnectedAccountPayout, error) {
	if !s.supported() {
		return nil, exception.NewUnsupportedEditionError("splitSettlement.GetConnectedAccountPayout", s.edition)
	}
	if err := param.Require("connectedAccountId", connectedAccountID); err != nil {
		return nil, err
	}
	if err := param.Require("payoutId", payoutID); err != nil {
		return nil, err
	}
	var result model.ConnectedAccountPayout
	path := "/connectedaccounts/" + internal.EncodePathSegment(connectedAccountID) +
		"/payouts/" + internal.EncodePathSegment(payoutID)
	if err := s.caller.GetInto(path, nil, &result, payoutEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SplitSettlementService) ListConnectedAccountPayoutTransactions(connectedAccountID,
	payoutID string) (*model.ListResponse[model.ConnectedAccountPayoutTransaction], error) {
	if !s.supported() {
		return nil, exception.NewUnsupportedEditionError("splitSettlement.ListConnectedAccountPayoutTransactions", s.edition)
	}
	if err := param.Require("connectedAccountId", connectedAccountID); err != nil {
		return nil, err
	}
	if err := param.Require("payoutId", payoutID); err != nil {
		return nil, err
	}
	path := "/connectedaccounts/" + internal.EncodePathSegment(connectedAccountID) +
		"/payouts/" + internal.EncodePathSegment(payoutID) + "/transactions"
	return internal.DoList[model.ConnectedAccountPayoutTransaction](s.caller, path, nil, transactionListEnvelope)
}

func (s *SplitSettlementService) ListConnectedAccountTransactions(connectedAccountID string,
	params *param.ConnectedAccountTransactionListParams) (*model.ListResponse[model.ConnectedAccountTransaction], error) {
	if !s.supported() {
		return nil, exception.NewUnsupportedEditionError("splitSettlement.ListConnectedAccountTransactions", s.edition)
	}
	if err := param.Require("connectedAccountId", connectedAccountID); err != nil {
		return nil, err
	}
	path := "/connectedaccounts/" + internal.EncodePathSegment(connectedAccountID) + "/transactions"
	return internal.DoList[model.ConnectedAccountTransaction](s.caller, path, params.ToQuery(), transactionListEnvelope)
}
