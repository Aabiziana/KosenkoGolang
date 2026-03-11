package domain

type Payroll struct {
	PayrollHeader   PayrollHeader
	PayrollBodyItems []PayrollBodyItem
}
