package model

import (
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
func (db *DB) AddTodo(text string, completed bool) uuid.UUID {
	id := uuid.New()
	db.Todos = append(db.Todos, Todo{id, text, completed})
	return id
}

// EditTodo ...
// func (db *DB) EditTodo(id uuid.UUID, ⚠️) {
// 	for _, v := range db.Todos {
// 		if v.ID == id {

// 		}
// 	}
// }
