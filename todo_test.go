package todo_test

import (
	"Develop/go-projects/todo"

	"testing"
)

func TestAddTodo(t *testing.T) {
	todos := todo.Todos{}
	title := "Vanquish foes"
	todos.Add(title)
	todos.Add("Make hummus")

	if len(todos) < 1 {
		t.Errorf("No todos found, expected at least %d", 1)
	}

	if len(todos) < 2 {
		t.Errorf("Not enough todos found, expected at least %d", 2)
	}

	if todos[0].Task != title {
		t.Errorf("Expected todo with title %q but got %q", title, todos[0].Task)
	}

	if todos[0].Id != 0 {
		t.Errorf("Expected todo with id %d but got %d", 0, todos[0].Id)
	}

	if todos[1].Id != 1 {
		t.Errorf("Expected todo with id %d but got %d", 1, todos[1].Id)
	}
}

func TestCompleteTodo(t *testing.T) {
	todos := todo.Todos{}
	todos.Add("Vanquish foes")
	completedTodo, err := todos.Complete(0)

	if err != nil {
		t.Errorf("%v", err)
	}

	if todos[0].Done != completedTodo.Done {
		t.Errorf("Error updating Completed todo Done status, expected %t, got %t", true, completedTodo.Done)
	}

	if todos[0].CompletedAt != completedTodo.CompletedAt {
		t.Errorf("Completed todo has invalid timestamp, expected %v, got %v", completedTodo.CompletedAt, todos[0].CompletedAt)
	}

}
