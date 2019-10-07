package main

import (
	"fmt"
	"time"
	"net/http"

	"go-world/config"
	"go-world/router"
	"go-world/logger"
)

func main() {
	appConfig := config.AppConfig()

	appLogger := logger.New(true)

	appRouter := router.New()

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", GoWorld)
	// mux.HandleFunc("/go", GoWorld)

	port := fmt.Sprintf("%d", appConfig.Server.Port)

	appLogger.Info().Msgf("Started server on 0.0.0.0:%s", port)
	// logger.Info().Msg("Server starting on 0.0.0.0")

	
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      appRouter,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		appLogger.Fatal().Err(err).Msg("Failed to start server")
	}
}

func GoWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go World!")
}