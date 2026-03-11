package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/documents/services/base"
	"github.com/joyzem/documents/services/payroll_statement/frontend/router"
)

func main() {

	handler := router.GetRouter()

	fmt.Println("Listening on 8087...")
	if err := http.ListenAndServe(":8087", handler); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
