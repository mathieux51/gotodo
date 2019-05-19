package model

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

// Todo ...
type Todo struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
}

// Todos ...
type Todos []Todo

// DB ...
type DB struct {
	Todos `json:"todos"`
}

// CreateDB returns a DB
// func CreateDB() *DB {
// 	return &DB{}
// }

// AddTodo append a todo to the todos array
func AddTodo(t Todo) error {
	j, err := os.Open("model/db.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}
	defer j.Close()
	b, _ := ioutil.ReadAll(j)
	// var todos Todos
	db := DB{}

	if err = json.Unmarshal(b, &db); err != nil {
		return err
	}
	db.Todos = append(db.Todos, t)
	f, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile("model/db.json", f, 0644); err != nil {
		return err
	}

	return nil
}

// UpdateTodo updates a todo text and completed fields given the id.
// func UpdateTodo(id uuid.UUID, todo Todo) error {
// 	for _, t := range db.Todos {
// 		if t.ID == id {
// 			t = todo
// 			t.ID = id
// 			return nil
// 		}
// 	}
// 	return errors.New("Todo not found")
// }
