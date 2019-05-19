package model

import (
	"errors"

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
func CreateDB() *DB {
	return &DB{}
}

// AddTodo append a todo to the todos array
func (db *DB) AddTodo(text string, completed bool) Todo {
	id := uuid.New()
	d := Todo{id, text, completed}
	db.Todos = append(db.Todos, d)
	return d
}

// UpdateTodo updates a todo text and completed fields given the id.
func (db *DB) UpdateTodo(id uuid.UUID, todo Todo) error {
	for _, t := range db.Todos {
		if t.ID == id {
			t = todo
			t.ID = id
			return nil
		}
	}
	return errors.New("Todo not found")
}
