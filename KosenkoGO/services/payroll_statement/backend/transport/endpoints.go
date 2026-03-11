package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/joyzem/documents/services/payroll_statement/backend/service"
	"github.com/joyzem/documents/services/payroll_statement/dto"
)

type Endpoints struct {
	CreatePayrollHeader   endpoint.Endpoint
	CreatePayrollBodyItem endpoint.Endpoint
	GetPayrolls           endpoint.Endpoint
	PayrollById           endpoint.Endpoint
	UpdatePayrollHeader   endpoint.Endpoint
	DeletePayroll         endpoint.Endpoint
	DeletePayrollBodyItem endpoint.Endpoint
}

func MakeEndpoints(s service.PayrollService) Endpoints {
	return Endpoints{
		CreatePayrollHeader:   makeCreatePayrollHeaderEndpoint(s),
		CreatePayrollBodyItem: makeCreatePayrollBodyItemEndpoint(s),
		GetPayrolls:           makeGetPayrollsEndpoint(s),
		PayrollById:           makePayrollByIdEndpoint(s),
		UpdatePayrollHeader:   makeUpdatePayrollHeaderEndpoint(s),
		DeletePayroll:         makeDeletePayrollEndpoint(s),
		DeletePayrollBodyItem: makeDeletePayrollBodyItemEndpoint(s),
	}
}

func makeCreatePayrollHeaderEndpoint(s service.PayrollService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.CreatePayrollHeaderRequest)
		header, err := s.CreatePayrollHeader(
			req.OrganizationId,
			req.DocumentNumber,
			req.DateOfIssue,
			req.PeriodStart,
			req.PeriodEnd,
			req.Subdivision,
			req.CorrespondingAccount,
			req.PaymentDeadlineStart,
			req.PaymentDeadlineEnd,
			req.TotalAmountWords,
			req.TotalAmount,
			req.DepositedAmount,
			req.DepositedAmountWords,
			req.Chief,
			req.FinancialChief,
			req.Cashier,
			req.CashOrderNumber,
			req.CashOrderDate,
			req.Accountant,
			req.SheetsCount,
		)
		return dto.CreatePayrollHeaderResponse{PayrollHeader: header}, err
	}
}

func makeCreatePayrollBodyItemEndpoint(s service.PayrollService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.CreatePayrollBodyItemRequest)
		payroll, err := s.CreatePayrollBodyItem(req.PayrollId, req.RowNumber, req.TabularNumber, req.EmployeeName, req.Amount, req.Signature, req.Note)
		return dto.CreatePayrollBodyItemResponse{Payroll: payroll}, err
	}
}

func makeGetPayrollsEndpoint(s service.PayrollService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		payrolls, err := s.GetPayrolls()
		return dto.GetPayrollsResponse{Payrolls: payrolls}, err
	}
}

func makePayrollByIdEndpoint(s service.PayrollService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.PayrollByIdRequest)
		payroll, err := s.PayrollById(req.Id)
		return dto.PayrollByIdResponse{Payroll: payroll}, err
	}
}

func makeUpdatePayrollHeaderEndpoint(s service.PayrollService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.UpdatePayrollHeaderRequest)
		payroll, err := s.UpdatePayrollHeader(req.Header)
		return dto.UpdatePayrollResponse{Payroll: payroll}, err
	}
}

func makeDeletePayrollBodyItemEndpoint(s service.PayrollService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.DeletePayrollBodyItemRequest)
		err = s.DeletePayrollBodyItem(req.ItemId)
		return dto.DeletePayrollResponse{}, err
	}
}

func makeDeletePayrollEndpoint(s service.PayrollService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(dto.DeletePayrollRequest)
		err = s.DeletePayroll(req.Id)
		return dto.DeletePayrollResponse{}, err
	}
}
