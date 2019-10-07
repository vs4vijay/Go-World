package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GoWorld)
	mux.HandleFunc("/go", GoWorld)

	log.Println("Started server on :9999")
	
	s := &http.Server{
		Addr:         ":9999",
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Failed to start server")
	}
}

func GoWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go World!")
}