package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"algogrit.com/empserver/entities"

	"algogrit.com/empserver/employee/repository"
	"algogrit.com/empserver/employee/service"
)

var empRepo = repository.NewInMemRepository()
var empSvcV1 = service.NewV1(empRepo)

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	employees, err := empSvcV1.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := empSvcV1.Create(newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // May be an application error => Validations
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdEmp)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	h := func(w http.ResponseWriter, req *http.Request) {
		begin := time.Now()

		next.ServeHTTP(w, req)

		dur := time.Since(begin)

		log.Printf("%s %s took %s\n", req.Method, req.URL, dur)
	} // Type: func(http.ResponseWrite, *http.Request)

	return http.HandlerFunc(h)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	}).Methods("GET")

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	log.Println("Starting server on port: 8000...")

	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
}
