package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go/config"
	"go.uber.org/zap"
)

var tracer opentracing.Tracer

func init() {
	cfg := config.Configuration{
		ServiceName: "golang-boilerplate",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}

	var err error
	tracer, _, err = cfg.NewTracer()
	if err != nil {
		logger.Error("failed to create tracer", zap.Error(err))
	}
	opentracing.SetGlobalTracer(tracer)
}

// TracingMiddleware adds distributed tracing to requests
func TracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if tracer == nil {
			c.Next()
			return
		}

		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		span := tracer.StartSpan(fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path), opentracing.ChildOf(spanCtx))
		defer span.Finish()

		ext.HTTPMethod.Set(span, c.Request.Method)
		ext.HTTPUrl.Set(span, c.Request.URL.String())
		ext.Component.Set(span, "gin")

		c.Set("span", span)
		c.Next()

		ext.HTTPStatusCode.Set(span, uint16(c.Writer.Status()))
		if c.Writer.Status() >= 400 {
			ext.Error.Set(span, true)
		}
	}
}
