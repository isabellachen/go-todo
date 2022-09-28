package main

import (
	todo "Develop/go-projects/todo/api"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("hello world"))
}

func todoRouter(todoFileName string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		todos := &todo.Todos{}
		urlHasParams := len(req.URL.Path) > 0

		if err := todos.Get(todoFileName); err != nil {
			replyError(w, req, http.StatusInternalServerError, err.Error())
			return
		}

		if err := todos.Get(todoFileName); err != nil {
			replyError(w, req, http.StatusInternalServerError, err.Error())
		}

		if !urlHasParams {
			switch req.Method {
			case "GET":
				response := &todoResponse{
					Results: *todos,
				}

				replyJsonContent(response, w, req, 200)
			case "POST":
				item := &struct {
					Task string `json:"task"`
				}{}

				if err := json.NewDecoder(req.Body).Decode(item); err != nil {
					message := fmt.Sprintf("Invalid JSON: %s", err)
					replyError(w, req, http.StatusInternalServerError, message)
				}

				todos.Add(item.Task)
				if err := todos.Save(todoFileName); err != nil {
					replyError(w, req, http.StatusInternalServerError, err.Error())
				}

				replyTextContent(w, req, http.StatusCreated, "")
			}
		}

		if urlHasParams {
			switch req.Method {
			case "GET":
				fmt.Println("Get one todo")
			case "PATCH":
				fmt.Println("Complete one todo")
			case "DELETE":
				fmt.Println("Delete one todo")
			}
		}

		// 		GET /todo
		// 		GET /todo/{number}
		// 		POST /todo
		// 		PATCH /todo/{number}?complete
		// 		DELETE /todo/{number}

	}
}

func getListTodosHandler(todoFileName string) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/list" {
			http.NotFound(w, req)
		}

		todos := &todo.Todos{}

		err := todos.Get(todoFileName)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Err: todos.GET", err)
			os.Exit(1)
		}

		todoResponse := &todoResponse{
			Results: *todos,
		}

		replyJsonContent(todoResponse, w, req, 200)
	}
}
