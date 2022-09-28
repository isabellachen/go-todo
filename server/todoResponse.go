package main

import (
	todo "Develop/go-projects/todo/api"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type todoResponse struct {
	Results todo.Todos `json:"results"`
}

func (res *todoResponse) MarshalJson() ([]byte, error) {
	data := struct {
		Results todo.Todos `json:"results"`
		Date    int64      `json:"date"`
		Count   int        `json:"total_results"`
	}{
		Results: res.Results,
		Date:    time.Now().Unix(),
		Count:   len(res.Results),
	}

	return json.Marshal(data)
}

func replyJsonContent(res *todoResponse, w http.ResponseWriter, req *http.Request, statusCode int) {
	body, err := res.MarshalJson()
	if err != nil {
		log.Printf("%s:%s Error %s, %d", req.URL, req.Method, http.StatusText(statusCode), statusCode)
		replyError(w, req, statusCode, fmt.Sprintf("%v", err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}

func replyTextContent(w http.ResponseWriter, req *http.Request, statusCode int, content string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(statusCode)
	w.Write([]byte(content))
}
