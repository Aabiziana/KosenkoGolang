package domain

type PayrollBodyItem struct {
	Id             int     `json:"id"`
	PayrollId      int     `json:"payroll_id"`
	EmployeeId     int     `json:"employee_id"`
	TabularNumber  string  `json:"tabular_number"`
	Amount         float64 `json:"amount"`
	Signature      string  `json:"signature"`
	Note           string  `json:"note"`
}
