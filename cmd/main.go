package main

import (
	"golang-boilerplate/main/config"
	"golang-boilerplate/main/server"
	"golang-boilerplate/main/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize all services
	if err := service.InitServices(); err != nil {
		log.Fatalf("Failed to initialize services: %v", err)
	}
	defer service.CloseServices()

	// Create and start server
	srv := server.NewServer()
	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	if err := srv.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
