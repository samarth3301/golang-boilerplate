package handlers

import (
	"context"
	"golang-boilerplate/main/config"
	"golang-boilerplate/main/models"
	"golang-boilerplate/main/repo"
	"golang-boilerplate/main/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var userRepo = repo.NewUserRepo()

// HealthHandler returns a 200 OK response if the service is healthy
func HealthHandler(c *gin.Context) {
	checks := checkServices()
	health := map[string]interface{}{
		"status":    "ok",
		"service":   "golang-boilerplate",
		"timestamp": time.Now().Format(time.RFC3339),
		"checks":    checks,
	}

	status := http.StatusOK
	for _, check := range checks {
		if check == "error" {
			status = http.StatusServiceUnavailable
			health["status"] = "degraded"
			break
		}
	}

	c.JSON(status, health)
}

func checkServices() map[string]string {
	checks := make(map[string]string)

	// Check database
	if service.DB == nil {
		checks["database"] = "not initialized"
	} else if err := service.DB.Ping(); err != nil {
		checks["database"] = "error"
	} else {
		checks["database"] = "ok"
	}

	// Check Redis
	if service.RedisClient == nil {
		checks["redis"] = "not initialized"
	} else if err := service.RedisClient.Ping(context.Background()).Err(); err != nil {
		checks["redis"] = "error"
	} else {
		checks["redis"] = "ok"
	}

	return checks
}

// PingHandler returns a simple pong response
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// RegisterHandler handles user registration
func RegisterHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := &models.User{
		Username:  req.Username,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := userRepo.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginHandler handles user login
func LoginHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userRepo.GetUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.AppConfig.JWT.Secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// ProtectedHandler is an example protected route
func ProtectedHandler(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Welcome to protected route", "user_id": userID})
}
