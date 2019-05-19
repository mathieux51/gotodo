package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mathieux51/mem-crud/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.TodosHandler).Methods("GET", "POST")
	r.HandleFunc("/todos/{id:\\d}", handlers.TodosByIDHandler).Methods("GET", "POST", "PUT", "DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
