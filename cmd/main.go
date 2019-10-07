package main

import (
	"fmt"
	"log"
	"time"
	"net/http"

	"go-world/config"
	"go-world/router"
)

func main() {
	appConfig := config.AppConfig()

	appRouter := router.New()

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", GoWorld)
	// mux.HandleFunc("/go", GoWorld)

	port := fmt.Sprintf("%d", appConfig.Server.Port)

	log.Printf("Started server on 0.0.0.0:%s\n", port)
	
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      appRouter,
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