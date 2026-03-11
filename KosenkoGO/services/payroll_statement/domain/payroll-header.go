package domain

type PayrollHeader struct {
	Id                    int     `json:"id"`
	OrganizationId        int     `json:"organization_id"`
	DocumentNumber        string  `json:"document_number"`
	DateOfIssue           string  `json:"date_of_issue"`
	PeriodStart           string  `json:"period_start"`
	PeriodEnd             string  `json:"period_end"`
	Subdivision           string  `json:"subdivision"`
	CorrespondingAccount  string  `json:"corresponding_account"`
	PaymentDeadlineStart  string  `json:"payment_deadline_start"`
	PaymentDeadlineEnd    string  `json:"payment_deadline_end"`
	TotalAmountWords      string  `json:"total_amount_words"`
	TotalAmount           float64 `json:"total_amount"`
	DepositedAmount       float64 `json:"deposited_amount"`
	DepositedAmountWords  string  `json:"deposited_amount_words"`
	Chief                 string  `json:"chief"`
	FinancialChief        string  `json:"financial_chief"`
	Cashier               string  `json:"cashier"`
	CashOrderNumber       string  `json:"cash_order_number"`
	CashOrderDate         string  `json:"cash_order_date"`
	Accountant            string  `json:"accountant"`
	SheetsCount           int     `json:"sheets_count"`
}
