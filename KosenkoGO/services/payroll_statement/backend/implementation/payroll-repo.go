package implementation

import (
	"database/sql"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/payroll_statement/backend/repo"
	"github.com/joyzem/documents/services/payroll_statement/domain"
)

type payrollRepo struct {
	db *sql.DB
}

func NewPayrollRepo(db *sql.DB) repo.PayrollRepo {
	return &payrollRepo{
		db: db,
	}
}

func (r *payrollRepo) CreatePayrollHeader(header domain.PayrollHeader) (*domain.PayrollHeader, error) {
	sql := `
		INSERT INTO payrolls 
		(
			organization_id,
			document_number,
			date_of_issue,
			period_start,
			period_end,
			subdivision,
			corresponding_account,
			payment_deadline_start,
			payment_deadline_end,
			total_amount_words,
			total_amount,
			deposited_amount,
			deposited_amount_words,
			chief,
			financial_chief,
			cashier,
			cash_order_number,
			cash_order_date,
			accountant,
			sheets_count
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20) RETURNING id
	`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedId int
	if err := result.QueryRow(
		header.OrganizationId,
		header.DocumentNumber,
		header.DateOfIssue,
		header.PeriodStart,
		header.PeriodEnd,
		header.Subdivision,
		header.CorrespondingAccount,
		header.PaymentDeadlineStart,
		header.PaymentDeadlineEnd,
		header.TotalAmountWords,
		header.TotalAmount,
		header.DepositedAmount,
		header.DepositedAmountWords,
		header.Chief,
		header.FinancialChief,
		header.Cashier,
		header.CashOrderNumber,
		header.CashOrderDate,
		header.Accountant,
		header.SheetsCount,
	).Scan(
		&insertedId,
	); err != nil {
		return nil, err
	}

	header.Id = insertedId
	return &header, nil
}

func (r *payrollRepo) GetPayrolls() ([]domain.PayrollHeader, error) {
	sql := `
		SELECT * FROM payrolls
	`
	rows, err := r.db.Query(sql)
	if err != nil {
		return nil, err
	}

	headers := []domain.PayrollHeader{}
	for rows.Next() {
		header := domain.PayrollHeader{}
		if err := rows.Scan(
			&header.Id,
			&header.OrganizationId,
			&header.DocumentNumber,
			&header.DateOfIssue,
			&header.PeriodStart,
			&header.PeriodEnd,
			&header.Subdivision,
			&header.CorrespondingAccount,
			&header.PaymentDeadlineStart,
			&header.PaymentDeadlineEnd,
			&header.TotalAmountWords,
			&header.TotalAmount,
			&header.DepositedAmount,
			&header.DepositedAmountWords,
			&header.Chief,
			&header.FinancialChief,
			&header.Cashier,
			&header.CashOrderNumber,
			&header.CashOrderDate,
			&header.Accountant,
			&header.SheetsCount,
		); err != nil {
			return nil, err
		}
		header.DateOfIssue, _ = base.ParseTime(header.DateOfIssue)
		header.PeriodStart, _ = base.ParseTime(header.PeriodStart)
		header.PeriodEnd, _ = base.ParseTime(header.PeriodEnd)
		header.PaymentDeadlineStart, _ = base.ParseTime(header.PaymentDeadlineStart)
		header.PaymentDeadlineEnd, _ = base.ParseTime(header.PaymentDeadlineEnd)
		header.CashOrderDate, _ = base.ParseTime(header.CashOrderDate)
		headers = append(headers, header)
	}

	return headers, err
}

func (r *payrollRepo) PayrollById(id int) (*domain.Payroll, error) {
	sql := `
		SELECT * FROM payrolls WHERE id = $1
	`
	payroll := domain.Payroll{}
	if err := r.db.QueryRow(sql, id).Scan(
		&payroll.PayrollHeader.Id,
		&payroll.PayrollHeader.OrganizationId,
		&payroll.PayrollHeader.DocumentNumber,
		&payroll.PayrollHeader.DateOfIssue,
		&payroll.PayrollHeader.PeriodStart,
		&payroll.PayrollHeader.PeriodEnd,
		&payroll.PayrollHeader.Subdivision,
		&payroll.PayrollHeader.CorrespondingAccount,
		&payroll.PayrollHeader.PaymentDeadlineStart,
		&payroll.PayrollHeader.PaymentDeadlineEnd,
		&payroll.PayrollHeader.TotalAmountWords,
		&payroll.PayrollHeader.TotalAmount,
		&payroll.PayrollHeader.DepositedAmount,
		&payroll.PayrollHeader.DepositedAmountWords,
		&payroll.PayrollHeader.Chief,
		&payroll.PayrollHeader.FinancialChief,
		&payroll.PayrollHeader.Cashier,
		&payroll.PayrollHeader.CashOrderNumber,
		&payroll.PayrollHeader.CashOrderDate,
		&payroll.PayrollHeader.Accountant,
		&payroll.PayrollHeader.SheetsCount,
	); err != nil {
		return nil, err
	}

	payroll.PayrollHeader.DateOfIssue, _ = base.ParseTime(payroll.PayrollHeader.DateOfIssue)
	payroll.PayrollHeader.PeriodStart, _ = base.ParseTime(payroll.PayrollHeader.PeriodStart)
	payroll.PayrollHeader.PeriodEnd, _ = base.ParseTime(payroll.PayrollHeader.PeriodEnd)
	payroll.PayrollHeader.PaymentDeadlineStart, _ = base.ParseTime(payroll.PayrollHeader.PaymentDeadlineStart)
	payroll.PayrollHeader.PaymentDeadlineEnd, _ = base.ParseTime(payroll.PayrollHeader.PaymentDeadlineEnd)
	payroll.PayrollHeader.CashOrderDate, _ = base.ParseTime(payroll.PayrollHeader.CashOrderDate)

	getBodySql := `SELECT * FROM payroll_bodies WHERE payroll_id = $1 ORDER BY row_number`
	rows, err := r.db.Query(getBodySql, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		bodyItem := domain.PayrollBodyItem{}
		if err := rows.Scan(
			&bodyItem.Id,
			&bodyItem.PayrollId,
			&bodyItem.RowNumber,
			&bodyItem.TabularNumber,
			&bodyItem.EmployeeName,
			&bodyItem.Amount,
			&bodyItem.Signature,
			&bodyItem.Note,
		); err != nil {
			return nil, err
		}
		payroll.PayrollBodyItems = append(payroll.PayrollBodyItems, bodyItem)
	}

	return &payroll, nil
}

func (r *payrollRepo) UpdatePayrollHeader(header domain.PayrollHeader) (*domain.Payroll, error) {
	sql := `
		UPDATE payrolls SET 
			organization_id = $1,
			document_number = $2,
			date_of_issue = $3,
			period_start = $4,
			period_end = $5,
			subdivision = $6,
			corresponding_account = $7,
			payment_deadline_start = $8,
			payment_deadline_end = $9,
			total_amount_words = $10,
			total_amount = $11,
			deposited_amount = $12,
			deposited_amount_words = $13,
			chief = $14,
			financial_chief = $15,
			cashier = $16,
			cash_order_number = $17,
			cash_order_date = $18,
			accountant = $19,
			sheets_count = $20
		WHERE id = $21
	`
	rows, err := r.db.Query(
		sql,
		header.OrganizationId,
		header.DocumentNumber,
		header.DateOfIssue,
		header.PeriodStart,
		header.PeriodEnd,
		header.Subdivision,
		header.CorrespondingAccount,
		header.PaymentDeadlineStart,
		header.PaymentDeadlineEnd,
		header.TotalAmountWords,
		header.TotalAmount,
		header.DepositedAmount,
		header.DepositedAmountWords,
		header.Chief,
		header.FinancialChief,
		header.Cashier,
		header.CashOrderNumber,
		header.CashOrderDate,
		header.Accountant,
		header.SheetsCount,
		header.Id,
	)

	if err != nil {
		return nil, err
	}
	rows.Close()

	return r.PayrollById(header.Id)
}

func (r *payrollRepo) DeletePayroll(id int) error {
	sql := `
		DELETE FROM payrolls WHERE id = $1
	`
	_, err := r.db.Exec(sql, id)
	return err
}

func (r *payrollRepo) recalcTotalAmount(payrollId int) error {
	var total float64
	if err := r.db.QueryRow(`SELECT COALESCE(SUM(amount), 0) FROM payroll_bodies WHERE payroll_id = $1`, payrollId).Scan(&total); err != nil {
		return err
	}
	_, err := r.db.Exec(`UPDATE payrolls SET total_amount = $1 WHERE id = $2`, total, payrollId)
	return err
}

func (r *payrollRepo) CreatePayrollBodyItem(item domain.PayrollBodyItem) (*domain.Payroll, error) {
	sql := `
		INSERT INTO payroll_bodies
		(
			payroll_id, 
			row_number,
			tabular_number,
			employee_name,
			amount,
			signature,
			note
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	result, err := r.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var insertedId int
	if err := result.QueryRow(item.PayrollId, item.RowNumber, item.TabularNumber, item.EmployeeName, item.Amount, item.Signature, item.Note).Scan(&insertedId); err != nil {
		return nil, err
	}

	if err := r.recalcTotalAmount(item.PayrollId); err != nil {
		return nil, err
	}

	return r.PayrollById(item.PayrollId)
}

func (r *payrollRepo) DeletePayrollBodyItem(id int) error {
	var payrollId int
	_ = r.db.QueryRow(`SELECT payroll_id FROM payroll_bodies WHERE id = $1`, id).Scan(&payrollId)
	_, err := r.db.Exec(`DELETE FROM payroll_bodies WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if payrollId != 0 {
		return r.recalcTotalAmount(payrollId)
	}
	return nil
}
