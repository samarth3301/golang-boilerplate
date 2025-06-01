package main

import (
	"log"
	"go-template/internal/server"
	"go-template/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize all services
	if err := service.InitServices(); err != nil {
		log.Fatalf("Failed to initialize services: %v", err)
	}
	defer service.CloseServices()

	// Create and start server
	srv := server.NewServer()
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 