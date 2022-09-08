package main

import (
	"Develop/go-projects/todo"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := "todos.json"
	todos := &todo.Todos{}

	err := todos.Get(filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	listTodosFlag := flag.Bool("list", false, "list todos")
	addTodoFlag := flag.String("add", "", "add a todo")
	completeTodoFlag := flag.Int("complete", -1, "complete a todo")

	flag.Parse()

	switch {
	case *listTodosFlag:
		for _, todo := range *todos {
			if !todo.Done {
				fmt.Println(todo.Task)
			}
		}
	case *addTodoFlag != "":
		task := *addTodoFlag
		todos.Add(task)
		if err := todos.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *completeTodoFlag > -1:
		index := *completeTodoFlag
		if err := todos.Complete(index); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := todos.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)

	}

}
