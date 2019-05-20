package model

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

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

// GetDB ...
func GetDB() (DB, error) {
	db := DB{}
	j, err := os.Open("model/db.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return db, err
	}
	defer j.Close()
	b, _ := ioutil.ReadAll(j)
	// var todos Todos

	if err = json.Unmarshal(b, &db); err != nil {
		return db, err
	}
	return db, nil
}

// GetTodos ...
func GetTodos() (Todos, error) {
	db, err := GetDB()
	if err != nil {
		return db.Todos, err
	}
	log.Println("> GET Todos")
	return db.Todos, nil
}

// GetTodoByID ...
func GetTodoByID(id uuid.UUID) (Todo, error) {
	var todo Todo
	db, err := GetDB()
	if err != nil {
		return todo, err
	}

	for _, v := range db.Todos {
		if v.ID == id {
			todo = v
			return todo, nil
		}
	}

	return todo, errors.New("Todo not found")
}

// AddTodo append a todo to the todos array
func AddTodo(t Todo) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	db.Todos = append(db.Todos, t)

	// Save to DB
	f, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile("model/db.json", f, 0644); err != nil {
		return err
	}

	// log
	log.Println("> POST Todo: ", t.ID)
	return nil
}

func remove(t Todos, i int) Todos {
	t[i] = t[len(t)-1]
	return t[:len(t)-1]
}

// DeleteTodoByID ...
func DeleteTodoByID(id uuid.UUID) error {
	db, err := GetDB()
	if err != nil {
		return err
	}

	for i, v := range db.Todos {
		if v.ID == id {
			todos := remove(db.Todos, i)
			db.Todos = todos

			// Save to DB
			f, err := json.MarshalIndent(db, "", " ")
			if err != nil {
				return err
			}
			if err = ioutil.WriteFile("model/db.json", f, 0644); err != nil {
				return err
			}

			// log
			log.Println("> DELETE Todo")
			return nil
		}
	}

	return errors.New("Todo not found")
}
