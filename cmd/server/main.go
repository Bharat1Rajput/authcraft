package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Bharat1Rajput/authcraft/internal/config"
)

func main() {
	cfg := config.Load()

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("auth service starting on port %s", cfg.Port)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
