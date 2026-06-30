package services

import (
	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/exception"
	"github.com/zoho/zoho-payments-go-sdk/internal"
	"github.com/zoho/zoho-payments-go-sdk/model"
	"github.com/zoho/zoho-payments-go-sdk/param"
)

const (
	mandateEnvelope             = "mandate"
	mandateNotificationEnvelope = "mandate_notification"
)

// MandateService provides recurring mandate operations (IN editions only).
type MandateService struct {
	caller  *internal.Caller
	edition edition.Edition
}

func NewMandateService(c *internal.Caller, ed edition.Edition) *MandateService {
	return &MandateService{caller: c, edition: ed}
}

func (s *MandateService) CreateEnrollmentSession(params *param.MandateEnrollmentSessionCreateParams) (*model.PaymentSession, error) {
	if !s.edition.IsIN() {
		return nil, exception.NewUnsupportedEditionError("mandates.CreateEnrollmentSession", s.edition)
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
	var result model.PaymentSession
	if err := s.caller.PostInto("/paymentsessions", body, &result, paymentSessionEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *MandateService) CreateExecutionSession(params *param.MandateExecutionSessionCreateParams) (*model.PaymentSession, error) {
	if !s.edition.IsIN() {
		return nil, exception.NewUnsupportedEditionError("mandates.CreateExecutionSession", s.edition)
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
	var result model.PaymentSession
	if err := s.caller.PostInto("/paymentsessions", body, &result, paymentSessionEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *MandateService) SendNotification(params *param.MandateNotifyParams) (*model.MandateNotification, error) {
	if !s.edition.IsIN() {
		return nil, exception.NewUnsupportedEditionError("mandates.SendNotification", s.edition)
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
	var result model.MandateNotification
	if err := s.caller.PostInto("/mandates/notify", body, &result, mandateNotificationEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *MandateService) Execute(params *param.MandateExecuteParams) (*model.MandatePayment, error) {
	if !s.edition.IsIN() {
		return nil, exception.NewUnsupportedEditionError("mandates.Execute", s.edition)
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
	var result model.MandatePayment
	if err := s.caller.PostInto("/mandates/execute", body, &result, paymentEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *MandateService) GetNotification(mandateNotificationID string) (*model.MandateNotification, error) {
	if !s.edition.IsIN() {
		return nil, exception.NewUnsupportedEditionError("mandates.GetNotification", s.edition)
	}
	if err := param.Require("mandateNotificationId", mandateNotificationID); err != nil {
		return nil, err
	}
	var result model.MandateNotification
	path := "/mandates/notifications/" + internal.EncodePathSegment(mandateNotificationID)
	if err := s.caller.GetInto(path, nil, &result, mandateNotificationEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *MandateService) Get(mandateID string) (*model.Mandate, error) {
	if !s.edition.IsIN() {
		return nil, exception.NewUnsupportedEditionError("mandates.Get", s.edition)
	}
	if err := param.Require("mandateId", mandateID); err != nil {
		return nil, err
	}
	var result model.Mandate
	path := "/mandates/" + internal.EncodePathSegment(mandateID)
	if err := s.caller.GetInto(path, nil, &result, mandateEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}
