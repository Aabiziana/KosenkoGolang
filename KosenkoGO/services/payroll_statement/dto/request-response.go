package dto

import "github.com/joyzem/documents/services/payroll_statement/domain"

type CreatePayrollHeaderRequest struct {
	OrganizationId       int     `json:"organization_id"`
	DocumentNumber       string  `json:"document_number"`
	DateOfIssue          string  `json:"date_of_issue"`
	PeriodStart          string  `json:"period_start"`
	PeriodEnd            string  `json:"period_end"`
	Subdivision          string  `json:"subdivision"`
	CorrespondingAccount string  `json:"corresponding_account"`
	PaymentDeadlineStart string  `json:"payment_deadline_start"`
	PaymentDeadlineEnd   string  `json:"payment_deadline_end"`
	TotalAmountWords     string  `json:"total_amount_words"`
	TotalAmount          float64 `json:"total_amount"`
	DepositedAmount      float64 `json:"deposited_amount"`
	DepositedAmountWords string  `json:"deposited_amount_words"`
	Chief                string  `json:"chief"`
	FinancialChief       string  `json:"financial_chief"`
	Cashier              string  `json:"cashier"`
	CashOrderNumber      string  `json:"cash_order_number"`
	CashOrderDate        string  `json:"cash_order_date"`
	Accountant           string  `json:"accountant"`
	SheetsCount          int     `json:"sheets_count"`
}

type CreatePayrollHeaderResponse struct {
	PayrollHeader *domain.PayrollHeader `json:"payroll_header"`
	Err           string                `json:"err,omitempty"`
}

type CreatePayrollBodyItemRequest struct {
	PayrollId     int     `json:"payroll_id"`
	RowNumber     int     `json:"row_number"`
	TabularNumber string  `json:"tabular_number"`
	EmployeeName  string  `json:"employee_name"`
	Amount        float64 `json:"amount"`
	Signature     string  `json:"signature"`
	Note          string  `json:"note"`
}

type CreatePayrollBodyItemResponse struct {
	Payroll *domain.Payroll `json:"payroll"`
	Err     string          `json:"err,omitempty"`
}

type GetPayrollsRequest struct{}

type GetPayrollsResponse struct {
	Payrolls []domain.PayrollHeader `json:"payrolls"`
	Err      string                 `json:"err,omitempty"`
}

type PayrollByIdRequest struct {
	Id int `json:"id"`
}

type PayrollByIdResponse struct {
	Payroll *domain.Payroll `json:"payroll"`
	Err     string          `json:"err,omitempty"`
}

type UpdatePayrollHeaderRequest struct {
	Header domain.PayrollHeader `json:"header"`
}

type UpdatePayrollResponse struct {
	Payroll *domain.Payroll `json:"payroll"`
	Err     string          `json:"err,omitempty"`
}

type DeletePayrollRequest struct {
	Id int `json:"id"`
}

type DeletePayrollResponse struct {
	Err string `json:"err,omitempty"`
}

type DeletePayrollBodyItemRequest struct {
	ItemId int `json:"item_id"`
}

type DeletePayrollBodyItemResponse struct {
	Err string `json:"err,omitempty"`
}
