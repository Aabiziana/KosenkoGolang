package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joyzem/documents/services/payroll_statement/frontend/handlers"
	"github.com/rs/cors"
)

func GetRouter() http.Handler {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("../../../../static"))))
	router.PathPrefix("/documents/static/").Handler(http.StripPrefix("/documents/static/", http.FileServer(http.Dir("../../../documents/static"))))

	router.HandleFunc("/documents/payrolls", handlers.PayrollsHandler)

	router.HandleFunc("/documents/payrolls/create", handlers.CreatePayrollGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/documents/payrolls/create", handlers.CreatePayrollPostHandler).Methods(http.MethodPost)

	router.HandleFunc("/documents/payrolls/update/{id:[0-9]+}", handlers.UpdatePayrollGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/documents/payrolls/update", handlers.UpdatePayrollPostHandler).Methods(http.MethodPost)

	router.HandleFunc("/documents/payrolls/delete", handlers.DeletePayrollHandler).Methods(http.MethodPost)

	router.HandleFunc("/documents/payrolls/details/{id:[0-9]+}", handlers.PayrollDetailsHandler).Methods(http.MethodGet)

	router.HandleFunc("/documents/payrolls/{id:[0-9]+}/body/create", handlers.CreatePayrollBodyGetHandler).Methods(http.MethodGet)
	router.HandleFunc("/documents/payrolls/body/create", handlers.CreatePayrollBodyPostHandler).Methods(http.MethodPost)
	router.HandleFunc("/documents/payrolls/body/delete", handlers.DeletePayrollBodyHandler).Methods(http.MethodPost)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
	})

	handler := c.Handler(router)
	return handler
}
