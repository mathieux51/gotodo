package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	"github.com/mathieux51/gotodo/model"
)

func getTodoFromBody(r *http.Request) (model.Todo, error) {
	var t model.Todo

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
	c redis.Conn
}

// NewTodoService ...
func NewTodoService(c redis.Conn) *TodoService {
	return &TodoService{c: c}
}

// TodoHander ...
func (s TodoService) TodoHander(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		// Response
		todos, err := model.GetTodos(s.c)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		jsonTodos, err := json.Marshal(todos)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(jsonTodos)

	case http.MethodPost:

		t, err := getTodoFromBody(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Generate id
		id, err := redis.Int(s.c.Do("incr", "todos:"))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		t.ID = id

		// Save to db
		if err = model.PostTodo(s.c, t); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Response
		jsonTodo, err := json.Marshal(t)
		if err != nil {
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
		http.Error(w, err.Error(), 500)
		return
	}
	switch r.Method {
	case http.MethodGet:
		todo, err := model.GetTodoByID(s.c, id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Response
		jsonTodo, err := json.Marshal(todo)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(jsonTodo)

	case http.MethodPut:

		t, err := getTodoFromBody(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		todo := model.Todo{ID: id, Text: t.Text, Completed: t.Completed}
		err = model.PutTodoByID(s.c, todo)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(200)

	case http.MethodDelete:
		err := model.DeleteTodoByID(s.c, id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(200)

	}

}
