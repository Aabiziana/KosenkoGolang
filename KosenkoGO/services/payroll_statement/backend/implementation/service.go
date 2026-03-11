package implementation

import (
	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/payroll_statement/backend/repo"
	"github.com/joyzem/documents/services/payroll_statement/backend/service"
	"github.com/joyzem/documents/services/payroll_statement/domain"
)

type payrollService struct {
	payrollRepo repo.PayrollRepo
}

func NewPayrollService(payrollRepo repo.PayrollRepo) service.PayrollService {
	return &payrollService{
		payrollRepo: payrollRepo,
	}
}

func (s *payrollService) CreatePayrollHeader(organizationId int, documentNumber string, dateOfIssue string, periodStart string, periodEnd string, subdivision string, correspondingAccount string, paymentDeadlineStart string, paymentDeadlineEnd string, totalAmountWords string, totalAmount float64, depositedAmount float64, depositedAmountWords string, chief string, financialChief string, cashier string, cashOrderNumber string, cashOrderDate string, accountant string, sheetsCount int) (*domain.PayrollHeader, error) {
	header := domain.PayrollHeader{
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
	createdHeader, err := s.payrollRepo.CreatePayrollHeader(header)
	base.LogError(err)
	if err != nil {
		return nil, err
	}
	return createdHeader, nil
}

func (s *payrollService) CreatePayrollBodyItem(payrollId int, rowNumber int, tabularNumber string, employeeName string, amount float64, signature string, note string) (*domain.Payroll, error) {
	item := domain.PayrollBodyItem{
		PayrollId:     payrollId,
		RowNumber:     rowNumber,
		TabularNumber: tabularNumber,
		EmployeeName:  employeeName,
		Amount:        amount,
		Signature:     signature,
		Note:          note,
	}
	payroll, err := s.payrollRepo.CreatePayrollBodyItem(item)
	base.LogError(err)
	return payroll, err
}

func (s *payrollService) GetPayrolls() ([]domain.PayrollHeader, error) {
	payrolls, err := s.payrollRepo.GetPayrolls()
	base.LogError(err)
	return payrolls, err
}

func (s *payrollService) PayrollById(id int) (*domain.Payroll, error) {
	payroll, err := s.payrollRepo.PayrollById(id)
	base.LogError(err)
	return payroll, err
}

func (s *payrollService) UpdatePayrollHeader(header domain.PayrollHeader) (*domain.Payroll, error) {
	payroll, err := s.payrollRepo.UpdatePayrollHeader(header)
	base.LogError(err)
	return payroll, err
}

func (s *payrollService) DeletePayroll(id int) error {
	err := s.payrollRepo.DeletePayroll(id)
	base.LogError(err)
	return err
}

func (s *payrollService) DeletePayrollBodyItem(itemId int) error {
	err := s.payrollRepo.DeletePayrollBodyItem(itemId)
	base.LogError(err)
	return err
}
