package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type HMuxHandler struct{}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello 42!\npath: %v", r.URL.Path)
}

func main() {
	mux := http.FileServer(http.Dir("."))

	server := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}



	log.Println("Server starting on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
