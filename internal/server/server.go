package server

import (
	"fmt"
	"os"
	"go-template/internal/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		router: gin.Default(),
	}
}

func (s *Server) Start() error {
	// Setup routes
	routes.SetupRoutes(s.router)

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	return s.router.Run(fmt.Sprintf(":%s", port))
} 