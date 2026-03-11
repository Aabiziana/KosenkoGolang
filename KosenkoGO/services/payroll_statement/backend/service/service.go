package service

import "github.com/joyzem/documents/services/payroll_statement/domain"

type PayrollService interface {
	CreatePayrollHeader(organizationId int, documentNumber string, dateOfIssue string, periodStart string, periodEnd string, subdivision string, correspondingAccount string, paymentDeadlineStart string, paymentDeadlineEnd string, totalAmountWords string, totalAmount float64, depositedAmount float64, depositedAmountWords string, chief string, financialChief string, cashier string, cashOrderNumber string, cashOrderDate string, accountant string, sheetsCount int) (*domain.PayrollHeader, error)
	CreatePayrollBodyItem(payrollId int, rowNumber int, tabularNumber string, employeeName string, amount float64, signature string, note string) (*domain.Payroll, error)
	GetPayrolls() ([]domain.PayrollHeader, error)
	PayrollById(id int) (*domain.Payroll, error)
	UpdatePayrollHeader(header domain.PayrollHeader) (*domain.Payroll, error)
	DeletePayroll(id int) error
	DeletePayrollBodyItem(itemId int) error
}
