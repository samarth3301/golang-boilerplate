package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthHandler returns a 200 OK response if the service is healthy
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"service":   "go-template",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// PingHandler returns a simple pong response
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// NotFoundHandler handles 404 errors
