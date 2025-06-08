package middleware

import (
	"time"

	fwsample "github.com/k-tsurumaki/fw-sample"
)

type LoggingMiddleware struct {
	Logger fwsample.Logger
}

func (m *LoggingMiddleware) Logging(next fwsample.HandlerFunc) fwsample.HandlerFunc {
	return func(ctx fwsample.Context) {
		start := time.Now()
		m.Logger.Info("Request start", "method", ctx.Request().Method, "path", ctx.Request().URL.Path)
		next(ctx)
		duration := time.Since(start)
		m.Logger.Info("Request complete", "duration", duration)
	}
}
