package todo_test

import (
	todo "Develop/go-projects/todo/api"
	"io/ioutil"
	"os"
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

	if todos[0].Id != 1 {
		t.Errorf("Expected todo with id %d but got %d", 1, todos[0].Id)
	}

	if todos[1].Id != 2 {
		t.Errorf("Expected todo with id %d but got %d", 2, todos[1].Id)
	}
}

func TestCompleteTodo(t *testing.T) {
	todos := todo.Todos{}
	todos.Add("Vanquish foes")
	err := todos.Complete(1)

	if err != nil {
		t.Errorf("%v", err)
	}

	if todos[0].Done != true {
		t.Errorf("Error updating Completed todo Done status, expected %t, got %t", true, false)
	}
}

func TestDeleteTodo(t *testing.T) {
	todos := todo.Todos{}
	todos.Add("Vanquish foes")
	todos.Add("Make hummus")
	todos.Add("Do the dishes")
	err := todos.Delete(2)
	if err != nil {
		t.Errorf("%v", err)
	}

	if len(todos) != 2 {
		t.Errorf("Error deleting todo, expected %d, got %d", 2, len(todos))
	}
}

func TestSaveGet(t *testing.T) {
	t1 := todo.Todos{}
	t2 := todo.Todos{}

	t1.Add("Vanquish foes")
	t1.Add("Make hummus")

	tmpFile, err := ioutil.TempFile("", "test_todos")

	if err != nil {
		t.Errorf("Error creating temp file %s", err)
	}

	defer os.Remove(tmpFile.Name())

	err = t1.Save(tmpFile.Name())

	if err != nil {
		t.Errorf("Error saving list fo file %s", err)
	}

	err = t2.Get(tmpFile.Name())

	if err != nil {
		t.Errorf("Error getting list from file %s", err)
	}

	if t1[0].Task != t2[0].Task {
		t.Errorf("Expected task from saved list to be %q, but got %q", t1[0].Task, t2[0].Task)
	}
}
