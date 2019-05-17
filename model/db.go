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
func (db *DB) AddTodo(text string, completed bool) uuid.UUID {
	id := uuid.New()
	db.Todos = append(db.Todos, Todo{id, text, completed})
	return id
}

// EditTodo ...
func (db *DB) EditTodo(id uuid.UUID, newTodo Todo) error {
	todo, err := db.findTodoByID(id)
	if err != nil {
		return err
	}
	return nil
}

// findTodoByID ...
func (db *DB) findTodoByID(id uuid.UUID) (Todo, error) {
	for _, v := range db.Todos {
		if v.ID == id {
			return v, nil
		}
	}
	return nil, errors.New("Todo not found")
	// cannot use nil as type Todo in return argument
}
