package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	empHTTP "algogrit.com/empserver/employee/http"
	"algogrit.com/empserver/employee/repository"
	"algogrit.com/empserver/employee/service"
)

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

	var empRepo = repository.NewInMemRepository()
	var empSvcV1 = service.NewV1(empRepo)
	var empHandler = empHTTP.NewHandler(empSvcV1)

	empHandler.SetupRoutes(r)

	log.Println("Starting server on port: 8000...")

	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
	// http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, empHandler))
}
