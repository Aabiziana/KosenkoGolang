package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/payroll_statement/backend/implementation"
	"github.com/joyzem/documents/services/payroll_statement/backend/transport"
	httptransport "github.com/joyzem/documents/services/payroll_statement/backend/transport/http"

	kithttp "github.com/go-kit/kit/transport/http"
)

func main() {

	db, err := base.ConnectToDb()
	if err != nil {
		base.LogError(err)
		os.Exit(-1)
	}

	defer db.Close()

	payrollRepo := implementation.NewPayrollRepo(db)

	svc := implementation.NewPayrollService(payrollRepo)

	endpoints := transport.MakeEndpoints(svc)

	h := httptransport.NewService(endpoints, []kithttp.ServerOption{})

	fmt.Println("Listening on 7077...")
	if err := http.ListenAndServe(":7077", h); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}

}
