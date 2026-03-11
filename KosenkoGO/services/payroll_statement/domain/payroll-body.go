package domain

type PayrollBodyItem struct {
	Id            int     `json:"id"`
	PayrollId     int     `json:"payroll_id"`
	RowNumber     int     `json:"row_number"`
	TabularNumber string  `json:"tabular_number"`
	EmployeeName  string  `json:"employee_name"`
	Amount        float64 `json:"amount"`
	Signature     string  `json:"signature"`
	Note          string  `json:"note"`
}
