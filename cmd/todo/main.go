package main

import (
	"Develop/go-projects/todo"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "todos.json"
	todos := &todo.Todos{}

	progArgs := os.Args[1:]

	err := todos.Get(filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(progArgs) < 1 {
		for _, todo := range *todos {
			fmt.Printf("Task: %q\n", todo.Task)
			fmt.Printf("Completed: %t\n", todo.Done)
		}
	} else {
		task := strings.Join(progArgs, " ")
		todos.Add(task)
		if err := todos.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
