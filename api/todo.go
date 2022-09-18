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
	id := len(*todos) + 1
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

func (todos *Todos) Complete(id int) error {
	index := id - 1
	todoList := *todos
	if index < 0 || index > len(todoList) {
		return fmt.Errorf("item %d does not exist", id)
	}
	todoList[index].Done = true
	todoList[index].CompletedAt = time.Now()

	return nil
}

func (todos *Todos) Delete(id int) error {
	index := id - 1
	todoList := *todos
	if index < 0 || index > len(todoList) {
		return fmt.Errorf("item %d does not exist", id)
	}

	*todos = append(todoList[:index], todoList[index+1:]...)

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

func (todos *Todos) String() string {
	formatted := ""

	for index, todo := range *todos {
		prefix := " "
		if todo.Done {
			prefix = "X"
		}
		formatted = formatted + fmt.Sprintf("%s %d: %s\n", prefix, index+1, todo.Task)
	}

	return formatted
}
