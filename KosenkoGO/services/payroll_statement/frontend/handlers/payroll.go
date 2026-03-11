package handlers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/payroll_statement/domain"
	"github.com/joyzem/documents/services/payroll_statement/dto"
	"github.com/joyzem/documents/services/payroll_statement/frontend/utils"
	"github.com/levigross/grequests"

	employeeDomain "github.com/joyzem/documents/services/employee/domain"
	organizationDomain "github.com/joyzem/documents/services/organization/domain"
)

func PayrollsHandler(w http.ResponseWriter, r *http.Request) {
	payrollsUrl := fmt.Sprintf("%s/payroll", utils.GetPayrollsAddress())
	resp, _ := grequests.Get(payrollsUrl, nil)
	var payrolls dto.GetPayrollsResponse
	resp.JSON(&payrolls)
	if payrolls.Err != "" {
		http.Error(w, payrolls.Err, http.StatusInternalServerError)
		return
	}

	type payrollTemplate struct {
		Id             int
		Organization   string
		DocumentNumber string
		DateOfIssue    string
		Subdivision    string
	}

	templateData := []payrollTemplate{}

	for _, payroll := range payrolls.Payrolls {

		organization, _ := utils.GetOrganizationById(payroll.OrganizationId)
		var organizationName string
		if organization.Err != "" {
			organizationName = fmt.Sprintf("Ошибка сервера: %s", organization.Err)
		} else {
			organizationName = organization.Organization.Name
		}

		payrollTmpl := payrollTemplate{
			Id:             payroll.Id,
			Organization:   organizationName,
			DocumentNumber: payroll.DocumentNumber,
			DateOfIssue:    payroll.DateOfIssue,
			Subdivision:    payroll.Subdivision,
		}
		templateData = append(templateData, payrollTmpl)
	}

	tmpl, _ := template.ParseFiles("../static/html/payrolls.html")
	tmpl.Execute(w, templateData)
}

func CreatePayrollGetHandler(w http.ResponseWriter, r *http.Request) {

	type payrollTemplate struct {
		Organizations []organizationDomain.Organization
		EmployeeNames []string
	}

	templateData := payrollTemplate{}

	organizations, _ := utils.GetOrganizations()
	if organizations.Err != "" {
		base.LogError(errors.New(organizations.Err))
	} else {
		templateData.Organizations = organizations.Organizations
	}

	employees, _ := utils.GetEmployees()
	if employees != nil && employees.Err == "" {
		names := []string{}
		for _, employee := range employees.Employees {
			names = append(names, employeeFullname(employee))
		}
		templateData.EmployeeNames = names
	}

	tmpl, _ := template.ParseFiles("../static/html/create-payroll.html")
	tmpl.Execute(w, templateData)
}

func CreatePayrollPostHandler(w http.ResponseWriter, r *http.Request) {
	organizationId, _ := strconv.Atoi(r.FormValue("organization_id"))
	documentNumber := r.FormValue("document_number")
	dateOfIssue := r.FormValue("date_of_issue")
	periodStart := r.FormValue("period_start")
	periodEnd := r.FormValue("period_end")
	subdivision := r.FormValue("subdivision")
	correspondingAccount := r.FormValue("corresponding_account")
	paymentDeadlineStart := r.FormValue("payment_deadline_start")
	paymentDeadlineEnd := r.FormValue("payment_deadline_end")
	totalAmountWords := r.FormValue("total_amount_words")
	totalAmount, _ := strconv.ParseFloat(r.FormValue("total_amount"), 64)
	depositedAmount, _ := strconv.ParseFloat(r.FormValue("deposited_amount"), 64)
	depositedAmountWords := r.FormValue("deposited_amount_words")
	chief := r.FormValue("chief")
	financialChief := r.FormValue("financial_chief")
	cashier := r.FormValue("cashier")
	cashOrderNumber := r.FormValue("cash_order_number")
	cashOrderDate := r.FormValue("cash_order_date")
	accountant := r.FormValue("accountant")
	sheetsCount, _ := strconv.Atoi(r.FormValue("sheets_count"))

	payroll := dto.CreatePayrollHeaderRequest{
		OrganizationId:       organizationId,
		DocumentNumber:       documentNumber,
		DateOfIssue:          dateOfIssue,
		PeriodStart:          periodStart,
		PeriodEnd:            periodEnd,
		Subdivision:          subdivision,
		CorrespondingAccount: correspondingAccount,
		PaymentDeadlineStart: paymentDeadlineStart,
		PaymentDeadlineEnd:   paymentDeadlineEnd,
		TotalAmountWords:     totalAmountWords,
		TotalAmount:          totalAmount,
		DepositedAmount:      depositedAmount,
		DepositedAmountWords: depositedAmountWords,
		Chief:                chief,
		FinancialChief:       financialChief,
		Cashier:              cashier,
		CashOrderNumber:      cashOrderNumber,
		CashOrderDate:        cashOrderDate,
		Accountant:           accountant,
		SheetsCount:          sheetsCount,
	}

	payrollUrl := fmt.Sprintf("%s/payroll", utils.GetPayrollsAddress())
	resp, _ := grequests.Post(payrollUrl, &grequests.RequestOptions{
		JSON: payroll,
	})

	var payrollResp dto.CreatePayrollHeaderResponse
	resp.JSON(&payrollResp)

	if payrollResp.Err != "" {
		http.Error(w, payrollResp.Err, http.StatusInternalServerError)
		return
	}

	redirectAddress := fmt.Sprintf("/documents/payrolls/update/%d", payrollResp.PayrollHeader.Id)
	http.Redirect(w, r, redirectAddress, http.StatusSeeOther)
}

func UpdatePayrollGetHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	payroll, _ := utils.GetPayrollById(id)
	if payroll.Err != "" {
		http.Error(w, payroll.Err, http.StatusInternalServerError)
		return
	}

	type payrollBodyTemplate struct {
		Id            int
		RowNumber     int
		TabularNumber string
		EmployeeName  string
		Amount        float64
		Signature     string
		Note          string
	}

	type updatePayrollTemplate struct {
		Payroll          domain.Payroll
		Organizations    []organizationDomain.Organization
		PayrollBodyItems []payrollBodyTemplate
		EmployeeNames    []string
	}

	templateData := updatePayrollTemplate{
		Payroll: *payroll.Payroll,
	}

	payrollBodyItems := []payrollBodyTemplate{}
	for _, bodyItem := range payroll.Payroll.PayrollBodyItems {
		payrollBodyItems = append(payrollBodyItems, payrollBodyTemplate{
			Id:            bodyItem.Id,
			RowNumber:     bodyItem.RowNumber,
			TabularNumber: bodyItem.TabularNumber,
			EmployeeName:  bodyItem.EmployeeName,
			Amount:        bodyItem.Amount,
			Signature:     bodyItem.Signature,
			Note:          bodyItem.Note,
		})
	}
	templateData.PayrollBodyItems = payrollBodyItems

	organizations, _ := utils.GetOrganizations()
	if organizations.Err != "" {
		base.LogError(errors.New(organizations.Err))
	} else {
		templateData.Organizations = organizations.Organizations
	}

	employees, _ := utils.GetEmployees()
	if employees != nil && employees.Err == "" {
		names := []string{}
		for _, employee := range employees.Employees {
			names = append(names, employeeFullname(employee))
		}
		templateData.EmployeeNames = names
	}

	tmpl, _ := template.ParseFiles("../static/html/update-payroll.html")
	tmpl.Execute(w, templateData)
}

func UpdatePayrollPostHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	organizationId, _ := strconv.Atoi(r.FormValue("organization_id"))
	documentNumber := r.FormValue("document_number")
	dateOfIssue := r.FormValue("date_of_issue")
	periodStart := r.FormValue("period_start")
	periodEnd := r.FormValue("period_end")
	subdivision := r.FormValue("subdivision")
	correspondingAccount := r.FormValue("corresponding_account")
	paymentDeadlineStart := r.FormValue("payment_deadline_start")
	paymentDeadlineEnd := r.FormValue("payment_deadline_end")
	totalAmountWords := r.FormValue("total_amount_words")
	totalAmount, _ := strconv.ParseFloat(r.FormValue("total_amount"), 64)
	depositedAmount, _ := strconv.ParseFloat(r.FormValue("deposited_amount"), 64)
	depositedAmountWords := r.FormValue("deposited_amount_words")
	chief := r.FormValue("chief")
	financialChief := r.FormValue("financial_chief")
	cashier := r.FormValue("cashier")
	cashOrderNumber := r.FormValue("cash_order_number")
	cashOrderDate := r.FormValue("cash_order_date")
	accountant := r.FormValue("accountant")
	sheetsCount, _ := strconv.Atoi(r.FormValue("sheets_count"))

	payroll := dto.UpdatePayrollHeaderRequest{
		Header: domain.PayrollHeader{
			Id:                   id,
			OrganizationId:       organizationId,
			DocumentNumber:       documentNumber,
			DateOfIssue:          dateOfIssue,
			PeriodStart:          periodStart,
			PeriodEnd:            periodEnd,
			Subdivision:          subdivision,
			CorrespondingAccount: correspondingAccount,
			PaymentDeadlineStart: paymentDeadlineStart,
			PaymentDeadlineEnd:   paymentDeadlineEnd,
			TotalAmountWords:     totalAmountWords,
			TotalAmount:          totalAmount,
			DepositedAmount:      depositedAmount,
			DepositedAmountWords: depositedAmountWords,
			Chief:                chief,
			FinancialChief:       financialChief,
			Cashier:              cashier,
			CashOrderNumber:      cashOrderNumber,
			CashOrderDate:        cashOrderDate,
			Accountant:           accountant,
			SheetsCount:          sheetsCount,
		}}

	payrollUrl := fmt.Sprintf("%s/payroll", utils.GetPayrollsAddress())
	resp, _ := grequests.Put(payrollUrl, &grequests.RequestOptions{
		JSON: payroll,
	})

	var payrollResp dto.UpdatePayrollResponse
	resp.JSON(&payrollResp)

	if payrollResp.Err != "" {
		http.Error(w, payrollResp.Err, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/documents/payrolls", http.StatusSeeOther)
}

func DeletePayrollHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	body := dto.DeletePayrollRequest{Id: id}
	url := fmt.Sprintf("%s/payroll", utils.GetPayrollsAddress())
	resp, _ := grequests.Delete(url, &grequests.RequestOptions{
		JSON: body,
	})
	var deleteResponse dto.DeletePayrollResponse
	resp.JSON(&deleteResponse)
	if deleteResponse.Err != "" {
		http.Error(w, deleteResponse.Err, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/documents/payrolls", http.StatusSeeOther)
}

func PayrollDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	payroll, _ := utils.GetPayrollById(id)
	if payroll.Err != "" {
		http.Error(w, payroll.Err, http.StatusInternalServerError)
		return
	}

	type payrollBodyTemplate struct {
		RowNumber     int
		TabularNumber string
		EmployeeName  string
		Amount        float64
		Signature     string
		Note          string
	}

	type payrollDetailTemplate struct {
		PayrollId            int
		Organization         string
		Okud                 string
		Okpo                 string
		DocumentNumber       string
		DateOfIssue          string
		PeriodStart          string
		PeriodEnd            string
		Subdivision          string
		CorrespondingAccount string
		PaymentDeadlineStart string
		PaymentDeadlineEnd   string
		TotalAmountWords     string
		TotalAmount          float64
		DepositedAmount      float64
		DepositedAmountWords string
		Chief                string
		FinancialChief       string
		Cashier              string
		CashOrderNumber      string
		CashOrderDate        string
		Accountant           string
		SheetsCount          int
		PayrollBodyItems     []payrollBodyTemplate
	}

	templateData := payrollDetailTemplate{
		PayrollId:            payroll.Payroll.PayrollHeader.Id,
		DocumentNumber:       payroll.Payroll.PayrollHeader.DocumentNumber,
		DateOfIssue:          payroll.Payroll.PayrollHeader.DateOfIssue,
		PeriodStart:          payroll.Payroll.PayrollHeader.PeriodStart,
		PeriodEnd:            payroll.Payroll.PayrollHeader.PeriodEnd,
		Subdivision:          payroll.Payroll.PayrollHeader.Subdivision,
		CorrespondingAccount: payroll.Payroll.PayrollHeader.CorrespondingAccount,
		PaymentDeadlineStart: payroll.Payroll.PayrollHeader.PaymentDeadlineStart,
		PaymentDeadlineEnd:   payroll.Payroll.PayrollHeader.PaymentDeadlineEnd,
		TotalAmountWords:     payroll.Payroll.PayrollHeader.TotalAmountWords,
		TotalAmount:          payroll.Payroll.PayrollHeader.TotalAmount,
		DepositedAmount:      payroll.Payroll.PayrollHeader.DepositedAmount,
		DepositedAmountWords: payroll.Payroll.PayrollHeader.DepositedAmountWords,
		Chief:                payroll.Payroll.PayrollHeader.Chief,
		FinancialChief:       payroll.Payroll.PayrollHeader.FinancialChief,
		Cashier:              payroll.Payroll.PayrollHeader.Cashier,
		CashOrderNumber:      payroll.Payroll.PayrollHeader.CashOrderNumber,
		CashOrderDate:        payroll.Payroll.PayrollHeader.CashOrderDate,
		Accountant:           payroll.Payroll.PayrollHeader.Accountant,
		SheetsCount:          payroll.Payroll.PayrollHeader.SheetsCount,
	}

	payrollBodyItems := []payrollBodyTemplate{}
	for _, bodyItem := range payroll.Payroll.PayrollBodyItems {
		payrollBodyItems = append(payrollBodyItems, payrollBodyTemplate{
			RowNumber:     bodyItem.RowNumber,
			TabularNumber: bodyItem.TabularNumber,
			EmployeeName:  bodyItem.EmployeeName,
			Amount:        bodyItem.Amount,
			Signature:     bodyItem.Signature,
			Note:          bodyItem.Note,
		})
	}
	templateData.PayrollBodyItems = payrollBodyItems

	organization, _ := utils.GetOrganizationById(payroll.Payroll.PayrollHeader.OrganizationId)
	if organization.Err != "" {
		base.LogError(errors.New(organization.Err))
		templateData.Organization = fmt.Sprintf("Ошибка: %s", organization.Err)
	} else {
		templateData.Organization = organization.Organization.Name
		templateData.Okud = organization.Organization.Okud
		templateData.Okpo = organization.Organization.Okpo
	}

	tmpl, _ := template.ParseFiles("../static/html/payroll-details.html")
	tmpl.Execute(w, templateData)
}

func employeeFullname(employee employeeDomain.Employee) string {
	if employee.MiddleName == "" {
		return fmt.Sprintf("%s %s", employee.LastName, employee.FirstName)
	}
	return fmt.Sprintf("%s %s %s", employee.LastName, employee.FirstName, employee.MiddleName)
}
