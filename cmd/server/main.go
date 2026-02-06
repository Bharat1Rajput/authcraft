package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Bharat1Rajput/authcraft/internal/config"
	authhttp "github.com/Bharat1Rajput/authcraft/internal/http/handlers"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()

	// Auth routes
	mux.HandleFunc("/auth/login", authhttp.Login)
	mux.HandleFunc("/auth/refresh", authhttp.Refresh)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("auth service starting on port %s", cfg.Port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
