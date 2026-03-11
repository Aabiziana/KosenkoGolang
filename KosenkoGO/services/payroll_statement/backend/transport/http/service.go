package http

import (
	"context"
	"net/http"
	"strconv"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/payroll_statement/backend/transport"
	"github.com/joyzem/documents/services/payroll_statement/dto"
)

func NewService(
	svcEndpoints transport.Endpoints,
	options []kithttp.ServerOption,
) http.Handler {
	router := mux.NewRouter()
	errorEncoder := kithttp.ServerErrorEncoder(base.EncodeErrorResponse)
	options = append(options, errorEncoder)
	
	router.Methods("POST").Path("/payroll").Handler(
		kithttp.NewServer(
			svcEndpoints.CreatePayrollHeader,
			decodeCreatePayrollHeaderRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/payroll").Handler(
		kithttp.NewServer(
			svcEndpoints.GetPayrolls,
			decodeGetPayrollsRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("GET").Path("/payroll/{id:[0-9]+}").Handler(
		kithttp.NewServer(
			svcEndpoints.PayrollById,
			decodePayrollByIdRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("PUT").Path("/payroll").Handler(
		kithttp.NewServer(
			svcEndpoints.UpdatePayrollHeader,
			decodeUpdatePayrollHeaderRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/payroll").Handler(
		kithttp.NewServer(
			svcEndpoints.DeletePayroll,
			decodeDeletePayrollRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("POST").Path("/payroll/body").Handler(
		kithttp.NewServer(
			svcEndpoints.CreatePayrollBodyItem,
			decodeCreatePayrollBodyItemRequest,
			base.EncodeResponse,
			options...,
		))

	router.Methods("DELETE").Path("/payroll/body").Handler(
		kithttp.NewServer(
			svcEndpoints.DeletePayrollBodyItem,
			decodeDeletePayrollBodyItemRequest,
			base.EncodeResponse,
			options...,
		))

	return router
}

func decodeCreatePayrollHeaderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreatePayrollHeaderRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeGetPayrollsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.GetPayrollsRequest
	return req, nil
}

func decodePayrollByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	return dto.PayrollByIdRequest{Id: id}, err
}

func decodeUpdatePayrollHeaderRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.UpdatePayrollHeaderRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeDeletePayrollRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.DeletePayrollRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeCreatePayrollBodyItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.CreatePayrollBodyItemRequest
	err := base.DecodeBody(r, &req)
	return req, err
}

func decodeDeletePayrollBodyItemRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req dto.DeletePayrollBodyItemRequest
	err := base.DecodeBody(r, &req)
	return req, err
}
