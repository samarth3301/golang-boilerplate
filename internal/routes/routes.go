package routes

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/handlers"
	"go-template/internal/middleware"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine) {
	// Add logging middleware
	router.Use(middleware.LoggingMiddleware())
	

	// API routes
	api := router.Group("/api")
	{
		api.GET("/ping", handlers.PingHandler)
		api.GET("/health", handlers.HealthHandler)
	}
}