package repo

import "github.com/joyzem/documents/services/payroll_statement/domain"

type PayrollRepo interface {
	CreatePayrollHeader(header domain.PayrollHeader) (*domain.PayrollHeader, error)
	CreatePayrollBodyItem(item domain.PayrollBodyItem) (*domain.Payroll, error)
	GetPayrolls() ([]domain.PayrollHeader, error)
	PayrollById(id int) (*domain.Payroll, error)
	UpdatePayrollHeader(header domain.PayrollHeader) (*domain.Payroll, error)
	DeletePayroll(id int) error
	DeletePayrollBodyItem(itemId int) error
}
