package utils

import (
	"fmt"

	"github.com/joyzem/documents/services/payroll_statement/dto"
	"github.com/levigross/grequests"
)

func GetPayrollById(id int) (*dto.PayrollByIdResponse, error) {
	payrollUrl := fmt.Sprintf("%s/payroll/%d", GetPayrollsAddress(), id)
	resp, err := grequests.Get(payrollUrl, &grequests.RequestOptions{
		JSON: dto.PayrollByIdRequest{
			Id: id,
		}})
	if err != nil {
		return nil, err
	}
	var payroll dto.PayrollByIdResponse
	resp.JSON(&payroll)
	return &payroll, nil
}
