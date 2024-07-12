package http

import (
	"github.com/gorilla/mux"

	"algogrit.com/empserver/employee/service"
)

type EmployeeHandler struct {
	svcV1 service.EmployeeService
	*mux.Router
}

func (h *EmployeeHandler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/v1/employees", h.IndexV1).Methods("GET")
	r.HandleFunc("/v1/employees", h.CreateV1).Methods("POST")

	h.Router = r
}

func NewHandler(svcV1 service.EmployeeService) EmployeeHandler {
	h := EmployeeHandler{svcV1: svcV1}

	h.SetupRoutes(mux.NewRouter())

	return h
}
