package handlers

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// TodosHandler handler that returns Hello World
func TodosHandler(w http.ResponseWriter, r *http.Request) {
	// /todos/1
	if r.Method == "GET" {
		p := strings.Split(r.URL.Path, "/")
		log.Println(p)
		io.WriteString(w, "GET")
	}
	if r.Method == "POST" {
		io.WriteString(w, "POST")

		// body, err := req.GetBody()
		// if err != nil {

		// 	// Cannot parse body
		// }

		// Do something with body
	}
	// marshall JSON
	// io.WriteString(w, "Hi")

}
