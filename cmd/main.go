package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mathieux51/gotodo/handlers"
	"github.com/mathieux51/gotodo/model"

	"github.com/gorilla/mux"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	redisHost := getEnv("REDIS_HOST", "127.0.0.1")
	redisURL := fmt.Sprintf("redis://%v:6379", redisHost)
	c, err := model.NewDB(redisURL)
	if err != nil {
		log.Panic(err)
	}

	todoService := handlers.NewTodoService(c)

	r := mux.NewRouter()

	r.HandleFunc("/todos", todoService.TodoHander).Methods("GET", "POST")
	r.HandleFunc("/todos/{id}", todoService.TodosByIDHandler).Methods("GET", "POST", "PUT", "DELETE")
	log.Println("> Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
