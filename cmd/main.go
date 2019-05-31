package main

import (
	"log"
	"net/http"

	"github.com/mathieux51/gotodo/handlers"
	"github.com/mathieux51/gotodo/model"

	"github.com/gorilla/mux"
)

func main() {

	c, err := model.NewDB("redis://127.0.0.1:6379")
	if err != nil {
		log.Panic(err)
	}

	todosHandler := handlers.NewTodosHandler(c)

	r := mux.NewRouter()

	r.HandleFunc("/todos", todosHandler).Methods("GET", "POST")
	// r.HandleFunc("/todos/{id}", handlers.TodosByIDHandler).Methods("GET", "POST", "PUT", "DELETE")
	log.Println("> Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
