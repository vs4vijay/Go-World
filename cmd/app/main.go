package main

import (
	"fmt"
	"log"
	"time"
	"net/http"

	"go-world/config"
)

func main() {
	appConfig := config.AppConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", GoWorld)
	mux.HandleFunc("/go", GoWorld)

	port := appConfig.Server.Port

	log.Println("Started server on 0.0.0.0:%s\n", port)
	
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
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