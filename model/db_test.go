package model

import (
	"fmt"
	"testing"
)

func TestAddTodo(t *testing.T) {
	db := CreateDB()

	if len(db.Todos) != 0 {
		t.Errorf("db.Todos should be empty")
	}

	text := "Play music"
	id := db.AddTodo(text, false)
	fmt.Println(id)

	if len(db.Todos) != 1 {
		t.Errorf("db.Todos should have one Todo")
	}

	// if id {
	// 	t.Errorf("db.Todos should have one Todo")
	// }

	// if db.Todos[0].Completed != todo.Completed {
	// 	t.Errorf("db.Todos[0] should not be completed")
	// }

}

// func TestEditTodo(t *testing.T) {
// 	db := CreateDB()
// 	todo1 := Todo{uuid.New(), "Play music", false}

// 	db.AddTodo(todo1)
// 	db.EditTodo()

// 	// if len(db.Todos) != 0 {
// 	// 	t.Errorf("db.Todos should be empty")
// 	// }

// }
