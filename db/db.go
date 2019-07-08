package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

// Storage ...
type Storage struct {
	Conn redis.Conn
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// NewStorage ...
func NewStorage() (*Storage, error) {
	redisHost := getEnv("REDIS_HOST", "127.0.0.1")
	redisURL := fmt.Sprintf("redis://%v:6379", redisHost)
	conn, err := NewDB(redisURL)
	if err != nil {
		return nil, err
	}
	return &Storage{Conn: conn}, nil
}

// Todo ...
type Todo struct {
	ID        int    `json:"id" redis:"id"`
	Text      string `json:"text" redis:"text"`
	Completed bool   `json:"completed" redis:"completed"`
}

// Todos ...
type Todos []Todo

func getTodoKey(id int) string {
	return fmt.Sprintf("todo:%v", id)
}

// NewDB ...
func NewDB(dataSource string) (redis.Conn, error) {
	// redis://user:secret@localhost:6379/0?foo=bar&qux=baz
	c, err := redis.DialURL(dataSource)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// GetTodos ...
func (s Storage) GetTodos() (*Todos, error) {
	// ⚠️ Could be good to have an example with hscan to compare perfs

	// Should return the last 10 todos
	// or it should accept parameter to get a specific range
	todoList, err := redis.Strings(s.Conn.Do("smembers", "todos"))
	if err != nil {
		return nil, err
	}

	var todos Todos
	for _, v := range todoList {
		t, err := redis.Values(s.Conn.Do("hgetall", v))
		if err != nil {
			return nil, err
		}
		var todo Todo
		if err := redis.ScanStruct(t, &todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)

	}
	log.Println("> GET Todos", todos)
	return &todos, nil
}

// GetTodoByID ...
func (s Storage) GetTodoByID(id int) (*Todo, error) {
	t, err := redis.Values(s.Conn.Do("hgetall", fmt.Sprintf("todo:%v", id)))
	if err != nil {
		return nil, err
	}
	// Test me
	if t == nil {
		return nil, errors.New("todo not found")
	}
	var todo Todo
	if err := redis.ScanStruct(t, &todo); err != nil {
		return nil, err
	}
	return &todo, nil
}

// GetID ...
func (s Storage) GetID() (int, error) {
	return redis.Int(s.Conn.Do("incr", "todos:"))
}

// PostTodo ...
func (s Storage) PostTodo(t Todo) error {
	key := getTodoKey(t.ID)

	// Keep track of all the todos
	err := s.Conn.Send("sadd", "todos", key)
	if err != nil {
		return err
	}
	// Create a todo hash, redis-cli: todo:id text Hey completed false id someID
	if _, err := s.Conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(&t)...); err != nil {
		return err
	}

	// logging
	log.Println("> POST Todo: ", t.ID)
	return nil
}

// DeleteTodoByID ...
func (s Storage) DeleteTodoByID(id int) error {
	key := getTodoKey(id)
	if _, err := s.Conn.Do("srem", "todos", key); err != nil {
		return err
	}
	if _, err := s.Conn.Do("del", key); err != nil {
		return err
	}

	return nil
}

// PutTodoByID ...
func (s Storage) PutTodoByID(todo Todo) error {
	key := getTodoKey(todo.ID)
	if _, err := s.Conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(&todo)...); err != nil {
		return err
	}
	return nil
}
