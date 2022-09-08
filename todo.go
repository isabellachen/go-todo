package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type todo struct {
	Id          int
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []todo

func (todos *Todos) Add(task string) bool {
	id := len(*todos)
	todo := todo{
		Id:          id,
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*todos = append(*todos, todo)

	return true
}

func (todos *Todos) Complete(i int) error {
	todoList := *todos
	if i < 0 || i > len(todoList) {
		return fmt.Errorf("item %d does not exist", i)
	}
	todoList[i].Done = true
	todoList[i].CompletedAt = time.Now()

	return nil
}

func (todos *Todos) Delete(i int) error {
	todoList := *todos
	if i < 0 || i > len(todoList) {
		return fmt.Errorf("item %d does not exist", i)
	}

	*todos = append(todoList[:i], todoList[i+1:]...)

	return nil
}

func (todos *Todos) Save(filename string) error {
	js, err := json.Marshal(todos)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

func (todos *Todos) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, todos)
}
