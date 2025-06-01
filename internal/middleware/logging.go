package middleware

import (
	"time"
	"go.uber.org/zap"
	"github.com/gin-gonic/gin"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		logger.Info("request completed",
			zap.String("path", path),
			zap.String("method", method),
			zap.Int("status", status),
			zap.Duration("latency", latency),
		)
	}
} 