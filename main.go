package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mathieux51/gotodo/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.TodosHandler).Methods("GET", "POST")
	r.HandleFunc("/todos/{id}", handlers.TodosByIDHandler).Methods("GET", "POST", "PUT", "DELETE")

	log.Println("> Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
