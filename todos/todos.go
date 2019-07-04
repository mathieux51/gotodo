package todos

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mathieux51/gotodo/db"
)

func getTodoFromBody(r *http.Request) (db.Todo, error) {
	var t db.Todo

	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return t, nil
	}

	// Unmarshal
	err = json.Unmarshal(b, &t)
	if err != nil {
		return t, err
	}
	return t, nil
}

// TodoService ...
type TodoService struct {
	storage *db.Storage
}

// NewTodoService ...
func NewTodoService(s *db.Storage) *TodoService {
	return &TodoService{storage: s}
}

// TodoHander ...
func (s TodoService) TodoHander(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		// Response
		todos, err := s.storage.GetTodos()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		jsonTodos, err := json.Marshal(todos)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(jsonTodos)

	case http.MethodPost:

		t, err := getTodoFromBody(r)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}

		// New id
		id, err := s.storage.GetID()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		t.ID = id

		// Save to db
		if err = s.storage.PostTodo(t); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}

		// Response
		jsonTodo, err := json.Marshal(t)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(jsonTodo)
	}
}

// TodosByIDHandler ...
func (s TodoService) TodosByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 500)
		return
	}
	switch r.Method {
	case http.MethodGet:
		todo, err := s.storage.GetTodoByID(id)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}

		// Response
		jsonTodo, err := json.Marshal(todo)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(jsonTodo)

	case http.MethodPut:

		t, err := getTodoFromBody(r)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		todo := db.Todo{ID: id, Text: t.Text, Completed: t.Completed}
		err = s.storage.PutTodoByID(todo)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(200)

	case http.MethodDelete:
		err := s.storage.DeleteTodoByID(id)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(200)
	}

}
