package main

import (
	"log"
	"net/http"

	"github.com/mathieux51/mem-crud/handlers"
)

func main() {
	http.HandleFunc("/todos/", handlers.TodosHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
