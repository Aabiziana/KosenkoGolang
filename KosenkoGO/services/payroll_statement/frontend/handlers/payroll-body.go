package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/payroll_statement/dto"
	"github.com/joyzem/documents/services/payroll_statement/frontend/utils"
	"github.com/levigross/grequests"
)

func CreatePayrollBodyGetHandler(w http.ResponseWriter, r *http.Request) {
	payrollIdStr := mux.Vars(r)["id"]
	payrollId, err := strconv.Atoi(payrollIdStr)
	if err != nil || payrollId == 0 {
		payrollId, _ = strconv.Atoi(r.URL.Query().Get("payroll_id"))
	}

	nextRowNumber := 1
	if payrollId != 0 {
		payroll, _ := utils.GetPayrollById(payrollId)
		if payroll != nil && payroll.Payroll != nil {
			max := 0
			for _, item := range payroll.Payroll.PayrollBodyItems {
				if item.RowNumber > max {
					max = item.RowNumber
				}
			}
			nextRowNumber = max + 1
			if nextRowNumber <= 0 {
				nextRowNumber = 1
			}
		}
	}

	type bodyTemplate struct {
		PayrollId      int
		NextRowNumber  int
		EmployeeNames  []string
	}

	templateData := bodyTemplate{
		PayrollId:     payrollId,
		NextRowNumber: nextRowNumber,
	}

	employees, _ := utils.GetEmployees()
	if employees != nil && employees.Err == "" {
		names := []string{}
		for _, employee := range employees.Employees {
			names = append(names, employeeFullname(employee))
		}
		templateData.EmployeeNames = names
	}

	tmpl, _ := template.ParseFiles("../static/html/create-payroll-body.html")
	tmpl.Execute(w, templateData)
}

func CreatePayrollBodyPostHandler(w http.ResponseWriter, r *http.Request) {
	payrollId, _ := strconv.Atoi(r.FormValue("payroll_id"))
	rowNumber, _ := strconv.Atoi(r.FormValue("row_number"))
	tabularNumber := r.FormValue("tabular_number")
	employeeName := r.FormValue("employee_name")
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	signature := r.FormValue("signature")
	note := r.FormValue("note")

	body := dto.CreatePayrollBodyItemRequest{
		PayrollId:     payrollId,
		RowNumber:     rowNumber,
		TabularNumber: tabularNumber,
		EmployeeName:  employeeName,
		Amount:        amount,
		Signature:     signature,
		Note:          note,
	}

	url := fmt.Sprintf("%s/payroll/body", utils.GetPayrollsAddress())
	resp, _ := grequests.Post(url, &grequests.RequestOptions{
		JSON: body,
	})

	var bodyResp dto.CreatePayrollBodyItemResponse
	resp.JSON(&bodyResp)

	if bodyResp.Err != "" {
		http.Error(w, bodyResp.Err, http.StatusInternalServerError)
		return
	}

	redirectAddress := fmt.Sprintf("/documents/payrolls/update/%d", payrollId)
	http.Redirect(w, r, redirectAddress, http.StatusSeeOther)
}

func DeletePayrollBodyHandler(w http.ResponseWriter, r *http.Request) {
	itemId, _ := strconv.Atoi(r.FormValue("item_id"))
	payrollId, _ := strconv.Atoi(r.FormValue("payroll_id"))

	body := dto.DeletePayrollBodyItemRequest{ItemId: itemId}
	url := fmt.Sprintf("%s/payroll/body", utils.GetPayrollsAddress())
	resp, _ := grequests.Delete(url, &grequests.RequestOptions{
		JSON: body,
	})

	var deleteResponse dto.DeletePayrollBodyItemResponse
	resp.JSON(&deleteResponse)
	if deleteResponse.Err != "" {
		http.Error(w, deleteResponse.Err, http.StatusInternalServerError)
		return
	}

	redirectAddress := fmt.Sprintf("/documents/payrolls/update/%d", payrollId)
	http.Redirect(w, r, redirectAddress, http.StatusSeeOther)
}
