package utils

import (
	"fmt"

	"github.com/joyzem/documents/services/base"
)

func GetPayrollsAddress() string {
	return fmt.Sprintf("http://localhost:%s", base.GetEnv("PAYROLL_PORT", "7077"))
}
