package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type HMuxHandler struct {}

func (h HMuxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello 42!\npath: %v", r.URL.Path)
}

func main() {
	var hmuxHandler HMuxHandler
	server := &http.Server{
		Addr:           ":8080",
		Handler:        hmuxHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
