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
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"-"`
}

// func (e Employee) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf(`{
// 		"id": %d,
// 		"name": "%s",
// 		"speciality": "%s",
// 	}`, e.ID, e.Name, e.Department)

// 	return []byte(jsonString), nil
// }

var employees = []Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Lakki", "Cloud", 2001},
	{3, "Muthu", "SRE", 10001},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmp.ID = len(employees) + 1
	employees = append(employees, newEmp)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmp)
}

// func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == "POST" {
// 		EmployeeCreateHandler(w, req)
// 	} else {
// 		EmployeesIndexHandler(w, req)
// 	}
// }

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
	// r := http.NewServeMux()
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello, World!"
		// w.Write([]byte(msg))

		fmt.Fprintln(w, msg)
	}).Methods("GET")

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")
	// r.HandleFunc("/employees", EmployeesHandler)

	// http.ListenAndServe(":8000", LoggingMiddleware(r))
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
}
