package model

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type persistToRedis interface {
	Save() error
}

// Todo ...
type Todo struct {
	ID        uuid.UUID `json:"id" redis:"id"`
	Text      string    `json:"text" redis:"text"`
	Completed bool      `json:"completed" redis:"completed"`
}

// Todos ...
type Todos []Todo

// NewDB ...
func NewDB(dataSource string) (redis.Conn, error) {
	// redis://user:secret@localhost:6379/0?foo=bar&qux=baz
	c, err := redis.DialURL(dataSource)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// // DB ...
// type DB struct {
// 	c     redis.Conn
// 	Todos `json:"todos"`
// }

// Save ...
// func (d DB) Save() error {
// 	f, err := json.MarshalIndent(d, "", " ")
// 	if err != nil {
// 		return err
// 	}
// 	if _, err = d.c.Do("set", "db", f); err != nil {
// 		return err
// 	}

// 	// log
// 	log.Println("> SAVED")
// 	return nil
// }

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
		var todo Todo
		t, err := redis.Values(c.Do("hgetall", v))
		if err != nil {
			return nil, err
		}

		if err := redis.ScanStruct(t, &todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)

	}
	log.Println("> GET Todos", todos)
	return &todos, nil
}

// GetTodoByID ...
// func GetTodoByID(id uuid.UUID) (Todo, error) {
// 	var todo Todo
// 	db, err := GetDB()
// 	if err != nil {
// 		return todo, err
// 	}

// 	for _, v := range db.Todos {
// 		if v.ID == id {
// 			todo = v
// 			return todo, nil
// 		}
// 	}

// 	return todo, errors.New("Todo not found")
// }

// PostTodo append a todo to the todos array
func PostTodo(c redis.Conn, t Todo) error {
	todoKey := fmt.Sprintf("todo:%v", t.ID)

	// Keep track of all the todos
	c.Send("sadd", "todos", todoKey)

	// Create a todo hash, redis-cli: todo:id text Hey completed false id someID
	if _, err := c.Do("HMSET", redis.Args{}.Add(todoKey).AddFlat(&t)...); err != nil {
		return err
	}

	log.Println("> POST Todo: ", t.ID)
	return nil
}

func remove(t Todos, i int) Todos {
	t[i] = t[len(t)-1]
	return t[:len(t)-1]
}

// DeleteTodoByID ...
// func DeleteTodoByID(id uuid.UUID) error {
// 	db, err := GetDB()
// 	if err != nil {
// 		return err
// 	}

// 	for i, v := range db.Todos {
// 		if v.ID == id {
// 			todos := remove(db.Todos, i)
// 			db.Todos = todos

// 			err := db.Save()
// 			if err != nil {
// 				return err
// 			}
// 			// log
// 			log.Println("> DELETE Todo")
// 			return nil
// 		}
// 	}

// 	return errors.New("Todo not found")
// }

// PutTodoByID ...
// func PutTodoByID(todo Todo) error {
// 	db, err := GetDB()
// 	if err != nil {
// 		return err
// 	}

// 	for i, v := range db.Todos {
// 		if v.ID == todo.ID {
// 			db.Todos[i].Text = todo.Text
// 			db.Todos[i].Completed = todo.Completed

// 			err := db.Save()
// 			if err != nil {
// 				return err
// 			}

// 			// log
// 			log.Println("> PUT Todo")

// 			return nil
// 		}
// 	}

// 	return errors.New("Todo not found")
// }
