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

	text := "Do stuff"
	completed := false
	id := db.AddTodo(text, completed)
	fmt.Println(id)

	if len(db.Todos) != 1 {
		t.Errorf("db.Todos should have one Todo")
	}

	if db.Todos[0].Text != text {
		t.Errorf("db.Todos[0] should not be completed")
	}

	if db.Todos[0].Completed != completed {
		t.Errorf("db.Todos[0] should not be completed")
	}
}

func TestEditTodo(t *testing.T) {
	db := CreateDB()
	text := "Do stuff"
	completed := false
	id := db.AddTodo(text, completed)

	// db.EditTodo()

	// if len(db.Todos) != 0 {
	// 	t.Errorf("db.Todos should be empty")
	// }

}
