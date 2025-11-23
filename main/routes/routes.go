package routes

import (
	"golang-boilerplate/main/handlers"
	"golang-boilerplate/main/middleware"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine) {
	// Global middlewares
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.SecurityHeadersMiddleware())
	router.Use(middleware.TracingMiddleware())
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.RateLimitMiddleware())

	// Metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Public routes
		public := v1.Group("")
		{
			public.GET("/ping", handlers.PingHandler)
			public.GET("/health", handlers.HealthHandler)
			public.POST("/login", handlers.LoginHandler)
			public.POST("/register", handlers.RegisterHandler)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/protected", handlers.ProtectedHandler)
		}
	}

	// Legacy API routes (for backward compatibility)
	public := router.Group("/api")
	{
		public.GET("/ping", handlers.PingHandler)
		public.GET("/health", handlers.HealthHandler)
		public.POST("/login", handlers.LoginHandler)
		public.POST("/register", handlers.RegisterHandler)
	}

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/protected", handlers.ProtectedHandler)
	}
}
