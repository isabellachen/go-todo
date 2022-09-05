package todo

import (
	"fmt"
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

func (todos *Todos) Add(task string) todo {
	id := len(*todos)
	todo := todo{
		Id:          id,
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*todos = append(*todos, todo)

	return todo
}

func (todos *Todos) Complete(i int) (*todo, error) {
	todoList := *todos
	if i < 0 || i > len(todoList) {
		return nil, fmt.Errorf("item %d does not exist", i)
	}
	todoList[i].Done = true
	todoList[i].CompletedAt = time.Now()
	return &todoList[i], nil
}
