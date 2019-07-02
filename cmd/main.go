package main

import (
	"log"
	"net/http"

	"github.com/mathieux51/gotodo/db"
	"github.com/mathieux51/gotodo/handlers"

	"github.com/gorilla/mux"
)

const port = "3001"

func main() {
	// Init storage
	s, err := db.NewStorage()
	if err != nil {
		log.Fatal(err)
	}
	// Init services
	todoService := handlers.NewTodoService(s)

	// Router
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/todos", todoService.TodoHander).Methods("GET", "POST")
	r.HandleFunc("/todos/{id}", todoService.TodosByIDHandler).Methods("GET", "POST", "PUT", "DELETE")

	// Start server
	done := make(chan bool)
	go func() {
		err := http.ListenAndServe(":"+port, r)
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("> Listening on port %v", port)
	<-done
}
