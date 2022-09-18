package main

import (
	todo "Develop/go-projects/todo/api"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var filename = "todos.json"

func getTask(reader io.Reader, args ...string) (string, error) {
	var task string
	if len(args) > 0 {
		task = strings.Join(args, " ")
		return task, nil
	}

	scanner := bufio.NewScanner(reader)

	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(scanner.Text()) == 0 {
		return "", fmt.Errorf("task cannot be blank")
	}

	return scanner.Text(), nil
}

func main() {
	if os.Getenv("TODO_FILENAME") != "" {
		filename = os.Getenv("TODO_FILENAME")
	}

	todos := &todo.Todos{}

	err := todos.Get(filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Err: todos.GET", err)
		os.Exit(1)
	}

	listTodosFlag := flag.Bool("list", false, "list todos")
	addTodoFlag := flag.Bool("add", false, "add a todo")
	deleteTodoFlag := flag.Int("delete", -1, "delete a todo")
	completeTodoFlag := flag.Int("complete", -1, "complete a todo")

	flag.Parse()

	switch {
	case *listTodosFlag:
		formatted := todos.String()
		fmt.Printf("%s", formatted)
	case *addTodoFlag:
		task, err := getTask(os.Stdin, flag.Args()...)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Err: getTask", err)
			os.Exit(1)
		}

		todos.Add(task)
		if err := todos.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, "Err: todos.Save", err)
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
	case *deleteTodoFlag > -1:
		index := *deleteTodoFlag
		if err := todos.Delete(index); err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}

		if err := todos.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option, use go run main.go -h for help")
		os.Exit(1)
	}

}
