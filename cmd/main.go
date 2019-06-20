package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mathieux51/gotodo/db"
	"github.com/mathieux51/gotodo/handlers"

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
	conn, err := db.NewDB(redisURL)
	if err != nil {
		log.Panic(err)
	}
	s := db.Storage{Conn: conn}

	todoService := handlers.NewTodoService(s)

	r := mux.NewRouter()

	r.HandleFunc("/todos", todoService.TodoHander).Methods("GET", "POST")
	r.HandleFunc("/todos/{id}", todoService.TodosByIDHandler).Methods("GET", "POST", "PUT", "DELETE")

	port := "3001"
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
