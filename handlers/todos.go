package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/mathieux51/mem-crud/model"
)

// Todo ...
type Todo struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
}

// TodosHandler ...
func TodosHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:

		// Response
		io.WriteString(w, "GET")

	case http.MethodPost:

		// Read body
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal
		var t model.Todo
		err = json.Unmarshal(b, &t)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		// Generate uuid
		id, err := uuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		t.ID = id

		// db.Todos.AddTodo()

		// Response
		output, err := json.Marshal(t)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)

	}

}

// TodosByIDHandler ...
func TodosByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		io.WriteString(w, "GET")
	case http.MethodPost:
		io.WriteString(w, "POST")
	case http.MethodPut:
		io.WriteString(w, "PUT")
	case http.MethodDelete:
		io.WriteString(w, "DELETE")

	}

}
