package main

import (
	"Develop/go-projects/todo"
	"flag"
	"fmt"
	"os"
)

func listTodos(todos *todo.Todos) {
	for _, todo := range *todos {
		fmt.Println(todo.Task)
	}
}

func main() {
	filename := "todos.json"
	todos := &todo.Todos{}
	progArgs := os.Args[1:]

	err := todos.Get(filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	listTodosFlag := flag.Bool("list", false, "list todos")
	addTodoFlag := flag.String("add", "", "add a todo")
	completeTodoFlag := flag.Int("complete", -1, "complete a todo")

	flag.Parse()

	if len(progArgs) < 1 {
		listTodos(todos)
	}

	if *listTodosFlag {
		listTodos(todos)
		os.Exit(0)
	}

	if *addTodoFlag != "" {
		task := *addTodoFlag
		todos.Add(task)
		if err := todos.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	if *completeTodoFlag != -1 {
		index := *completeTodoFlag
		todos.Delete(index)
		if err := todos.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

}
