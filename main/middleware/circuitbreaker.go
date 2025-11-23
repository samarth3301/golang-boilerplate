package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eapache/go-resiliency/breaker"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	postgresBreaker *breaker.Breaker
	redisBreaker    *breaker.Breaker
)

func init() {
	// Initialize circuit breakers with error threshold and success threshold
	postgresBreaker = breaker.New(3, 1, 5*time.Second) // 3 failures, 1 success, 5s timeout
	redisBreaker = breaker.New(3, 1, 5*time.Second)
}

// CircuitBreakerMiddleware wraps handlers with circuit breaker protection
func CircuitBreakerMiddleware(service string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b *breaker.Breaker
		switch service {
		case "postgres":
			b = postgresBreaker
		case "redis":
			b = redisBreaker
		default:
			c.Next()
			return
		}

		err := b.Run(func() error {
			c.Next()
			if c.Writer.Status() >= 500 {
				return fmt.Errorf("service error: %d", c.Writer.Status())
			}
			return nil
		})

		if err != nil {
			logger.Error("circuit breaker tripped",
				zap.String("service", service),
				zap.Error(err),
			)
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
				"error":   "Service temporarily unavailable",
				"service": service,
			})
			return
		}
	}
}
