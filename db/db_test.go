package db

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	rh := "redis-host"
	os.Setenv("REDIS_HOST", rh)
	defer os.Unsetenv("REDIS_HOST")

	got := getEnv("REDIS_HOST", "127.0.0.1")
	want := rh
	if got != want {
		t.Errorf("getEnv(\"REDIS_HOST\", \"127.0.0.1\") = %v, got %v", want, got)
	}
}

func TestGetEnv_fallback(t *testing.T) {
	got := getEnv("REDIS_HOST", "127.0.0.1")
	want := "127.0.0.1"
	if got != want {
		t.Errorf("getEnv(\"REDIS_HOST\", \"127.0.0.1\") = %v, got %v", want, got)
	}
}

func TestGetStorage(t *testing.T) {
	s, err := NewStorage()
	if err != nil {
		t.Fatal(err)
	}
	got, err := s.Conn.Do("PING")
	if err != nil {
		t.Fatal(err)
	}
	want := "PONG"
	if got != want {
		t.Errorf("getStorage() = %v, got %v", want, got)
	}
}

func TestGetStorage_err(t *testing.T) {
	// Init
	os.Setenv("REDIS_HOST", "foo")
	defer os.Unsetenv("REDIS_HOST")

	_, err := NewStorage()
	if err == nil {
		t.Fatal(err)
		t.Error("GetStorage should return an error if REDIS_HOST invalid")
	}
}

func TestGetTodoKey(t *testing.T) {
	want := "todo:1"
	got := getTodoKey(1)
	if got != want {
		t.Errorf("getTodoKey() = %v, got %v", want, got)
	}
}
func TestGetTodos(t *testing.T) {
	s, err := NewStorage()
	if err != nil {
		t.Fatal(err)
	}
	todos, err := s.GetTodos()
	if err != nil {
		t.Fatal(err)
	}
	got := len(*todos)
	want := 2
	if got != want {
		t.Errorf("getTodoKey() = %v, got %v", want, got)
	}
}
