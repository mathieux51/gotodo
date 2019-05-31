package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
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

// NewTodosHandler ...
func NewTodosHandler(c redis.Conn) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			// Response
			todos, err := model.GetTodos(c)
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

			// Generate uuid
			id, err := uuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			t.ID = id

			// Save to db
			if err = model.PostTodo(c, t); err != nil {
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
}

// TodosByIDHandler ...
// func TodosByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := uuid.Parse(vars["id"])
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	switch r.Method {
// 	case http.MethodGet:
// 		todo, err := model.GetTodoByID(id)
// 		if err != nil {
// 			http.Error(w, err.Error(), 500)
// 			return
// 		}

// 		// Response
// 		jsonTodo, err := json.Marshal(todo)
// 		if err != nil {
// 			http.Error(w, err.Error(), 500)
// 			return
// 		}
// 		w.Header().Set("content-type", "application/json")
// 		w.Write(jsonTodo)

// 	case http.MethodPut:

// 		t, err := getTodoFromBody(r)
// 		if err != nil {
// 			http.Error(w, err.Error(), 500)
// 			return
// 		}
// 		todo := model.Todo{ID: id, Text: t.Text, Completed: t.Completed}
// 		err = model.PutTodoByID(todo)
// 		if err != nil {
// 			http.Error(w, err.Error(), 500)
// 			return
// 		}
// 		w.WriteHeader(200)

// 	case http.MethodDelete:
// 		err := model.DeleteTodoByID(id)
// 		if err != nil {
// 			http.Error(w, err.Error(), 500)
// 			return
// 		}
// 		w.WriteHeader(200)

// 	}

// }
