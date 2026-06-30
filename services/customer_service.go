package services

import (
	"github.com/zoho/zoho-payments-go-sdk/edition"
	"github.com/zoho/zoho-payments-go-sdk/exception"
	"github.com/zoho/zoho-payments-go-sdk/internal"
	"github.com/zoho/zoho-payments-go-sdk/model"
	"github.com/zoho/zoho-payments-go-sdk/param"
)

const (
	customerEnvelope     = "customer"
	customerListEnvelope = "customers"
)

// CustomerService provides customer operations.
type CustomerService struct {
	caller  *internal.Caller
	edition edition.Edition
}

func NewCustomerService(c *internal.Caller, ed edition.Edition) *CustomerService {
	return &CustomerService{caller: c, edition: ed}
}

func (s *CustomerService) Create(params *param.CustomerCreateParams) (*model.Customer, error) {
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
	var result model.Customer
	if err := s.caller.PostInto("/customers", body, &result, customerEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *CustomerService) Get(customerID string) (*model.Customer, error) {
	if err := param.Require("customerId", customerID); err != nil {
		return nil, err
	}
	var result model.Customer
	path := "/customers/" + internal.EncodePathSegment(customerID)
	if err := s.caller.GetInto(path, nil, &result, customerEnvelope); err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *CustomerService) List(params *param.CustomerListParams) (*model.ListResponse[model.CustomerSummary], error) {
	if !s.edition.IsUS() {
		return nil, exception.NewUnsupportedEditionError("customers.List", s.edition)
	}
	return internal.DoList[model.CustomerSummary](s.caller, "/customers", params.ToQuery(), customerListEnvelope)
}

func (s *CustomerService) Delete(customerID string) error {
	if !s.edition.IsUS() {
		return exception.NewUnsupportedEditionError("customers.Delete", s.edition)
	}
	if err := param.Require("customerId", customerID); err != nil {
		return err
	}
	return s.caller.DeleteResource("/customers/" + internal.EncodePathSegment(customerID))
}
