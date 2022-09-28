package main

import (
	"log"
	"net/http"
)

func newMux(todoFileName string) http.Handler {
	// Anything in Go can be a handler as long as it satisfules the
	// http.Handler interface. That is to say, a handler must have
	// the method ServeHTTP(http.ResponseWriter, *http.Request)

	mux := http.NewServeMux()
	router := todoRouter(todoFileName)
	mux.HandleFunc("/", rootHandler)
	mux.Handle("/todo", http.StripPrefix("/todo", router))
	mux.Handle("/todo/", http.StripPrefix("/todo/", router))

	return mux
}

func replyError(w http.ResponseWriter, req *http.Request, statusCode int, message string) {
	log.Printf("%s:%s Error %d %s", req.URL, req.Method, statusCode, message)
	http.Error(w, http.StatusText(statusCode), statusCode)
}
