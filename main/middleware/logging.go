package middleware

import (
	"strconv"
	"time"

	"golang-boilerplate/pkg/metrics"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

// LoggingMiddleware logs HTTP requests with metrics
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Add request ID for tracing
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
			c.Header("X-Request-ID", requestID)
		}
		c.Set("request_id", requestID)

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		// Record metrics
		metrics.HTTPRequestDuration.WithLabelValues(method, path, strconv.Itoa(status)).Observe(latency.Seconds())
		metrics.HTTPRequestTotal.WithLabelValues(method, path, strconv.Itoa(status)).Inc()

		// Structured logging
		logger.Info("request completed",
			zap.String("request_id", requestID),
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
			zap.Duration("latency", latency),
			zap.String("user_agent", c.GetHeader("User-Agent")),
			zap.String("ip", c.ClientIP()),
		)
	}
}

func generateRequestID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 36)
}
