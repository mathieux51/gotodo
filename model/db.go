package model

import (
	"errors"
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

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
func GetTodos(c redis.Conn) (*Todos, error) {
	// ⚠️ Could be good to have an example with hscan to compare perfs

	// Should return the last 10 todos
	// or it should accept parameter to get a specific range
	todoList, err := redis.Strings(c.Do("smembers", "todos"))
	if err != nil {
		return nil, err
	}
	var todos Todos

	for _, v := range todoList {
		t, err := redis.Values(c.Do("hgetall", v))
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
func GetTodoByID(c redis.Conn, id int) (*Todo, error) {
	t, err := redis.Values(c.Do("hgetall", fmt.Sprintf("todo:%v", id)))
	if err != nil {
		return nil, err
	}
	// Test me
	if t == nil {
		return nil, errors.New("Todo not found")
	}
	var todo Todo
	if err := redis.ScanStruct(t, &todo); err != nil {
		return nil, err
	}
	return &todo, nil
}

// PostTodo ...
func PostTodo(c redis.Conn, t Todo) error {
	key := getTodoKey(t.ID)

	// Keep track of all the todos
	c.Send("sadd", "todos", key)

	// Create a todo hash, redis-cli: todo:id text Hey completed false id someID
	if _, err := c.Do("HMSET", redis.Args{}.Add(key).AddFlat(&t)...); err != nil {
		return err
	}

	// logging
	log.Println("> POST Todo: ", t.ID)
	return nil
}

// DeleteTodoByID ...
func DeleteTodoByID(c redis.Conn, id int) error {
	key := getTodoKey(id)
	if _, err := c.Do("srem", "todos", key); err != nil {
		return err
	}
	if _, err := c.Do("del", key); err != nil {
		return err
	}

	return nil
}

// PutTodoByID ...
func PutTodoByID(c redis.Conn, todo Todo) error {
	key := getTodoKey(todo.ID)
	if _, err := c.Do("HMSET", redis.Args{}.Add(key).AddFlat(&todo)...); err != nil {
		return err
	}
	return nil
}
