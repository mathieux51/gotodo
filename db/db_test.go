package db

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	rh := "redis-host"
	os.Setenv("REDIS_HOST", rh)
	got := getEnv("REDIS_HOST", "127.0.0.1")
	want := rh
	if got != want {
		t.Errorf("getEnv(\"REDIS_HOST\", \"127.0.0.1\") = %v, got %v", want, got)
	}
}

func TestGetEnv_fallback(t *testing.T) {
	// Init
	os.Unsetenv("REDIS_HOST")

	got := getEnv("REDIS_HOST", "127.0.0.1")
	want := "127.0.0.1"
	if got != want {
		t.Errorf("getEnv(\"REDIS_HOST\", \"127.0.0.1\") = %v, got %v", want, got)
	}
}

// func TestGetStorage(t *testing.T) {
// 	// Init
// 	os.Unsetenv("REDIS_HOST")

// 	got := NewStorage()
// 	want := Storage{}
// 	if got != want {
// 		t.Errorf("getStorage() = %v, got %v", want, got)
// 	}
// }

// func TestPostTodo(t *testing.T) {
// 	db, err := NewDB()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	if len(db.Todos) != 0 {
// 		t.Errorf("db.Todos should be empty")
// 	}

// 	text := "Do stuff"
// 	completed := false
// 	db.PostTodo(text, completed)

// 	if len(db.Todos) != 1 {
// 		t.Errorf("db.Todos should have one Todo")
// 	}

// 	if db.Todos[0].Text != text {
// 		t.Errorf("db.Todos[0] should not be completed")
// 	}

// 	if db.Todos[0].Completed != completed {
// 		t.Errorf("db.Todos[0] should not be completed")
// 	}
// }

// func TestEditTodo(t *testing.T) {
// 	db := CreateDB()
// 	text := "Do stuff"
// 	completed := false

// }
