package domain

type Payroll struct {
	Header PayrollHeader     `json:"header"`
	Body   []PayrollBodyItem `json:"body"`
}
