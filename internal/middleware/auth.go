package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware logs HTTP requests
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start time
		start := time.Now()
		path := c.Request.URL.Path
		
		// Process request
		c.Next()
		
		// End time
		end := time.Now()
		latency := end.Sub(start)
		
		// Log request
		fmt.Printf("[%s] %s %s %d %s\n",
			end.Format("2006-01-02 15:04:05"),
			c.Request.Method,
			path,
			c.Writer.Status(),
			latency,
		)
	}
}
