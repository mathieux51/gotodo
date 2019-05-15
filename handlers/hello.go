package handlers

import (
	"io"
	"net/http"
)

// HelloHandler handler that returns Hello World
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
