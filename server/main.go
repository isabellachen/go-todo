package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	portNumber := 8080
	host := flag.String("host", "localhost", "Server host")
	port := flag.Int("port", portNumber, "Server port")
	todoFileName := flag.String("file", "todos.json", "todo JSON filename")

	fmt.Printf("Listening on :%d...", portNumber)
	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", *host, *port),
		Handler:      newMux(*todoFileName),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.ListenAndServe()
}
