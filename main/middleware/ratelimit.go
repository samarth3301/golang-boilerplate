package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	limiters  = make(map[string]*rate.Limiter)
	mu        sync.RWMutex
	rateLimit = rate.Every(1 * time.Second) // 1 request per second
	burst     = 5                           // burst of 5
)

func getLimiter(key string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := limiters[key]
	if !exists {
		limiter = rate.NewLimiter(rateLimit, burst)
		limiters[key] = limiter
	}

	return limiter
}

func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Use user ID if authenticated, otherwise use IP
		key := c.ClientIP()
		if userID, exists := c.Get("user_id"); exists {
			key = userID.(string)
		}

		limiter := getLimiter(key)
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
