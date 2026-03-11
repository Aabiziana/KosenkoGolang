package utils

import (
	"fmt"

	"github.com/joyzem/documents/services/base"
	employeeDto "github.com/joyzem/documents/services/employee/dto"
	"github.com/levigross/grequests"
)

func GetEmployees() (*employeeDto.GetEmployeesResponse, error) {
	employeesUrl := fmt.Sprintf("http://localhost:%s/employees", base.GetEnv("EMPLOYEE_PORT", "7074"))
	resp, err := grequests.Get(employeesUrl, nil)
	if err != nil {
		return nil, err
	}
	var employees employeeDto.GetEmployeesResponse
	resp.JSON(&employees)
	return &employees, nil
}
