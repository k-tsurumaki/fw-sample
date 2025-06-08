package middleware

import (
	"time"
	"fmt"

	fwsample "github.com/k-tsurumaki/fw-sample"
)

type Logger interface {
	Info(ctx fwsample.Context, msg string)
	Error(ctx fwsample.Context, msg string)
}

type LoggingMiddleware struct {
	Logger Logger
}

func (m *LoggingMiddleware) Logging(next fwsample.HandlerFunc) fwsample.HandlerFunc {
	return func(ctx fwsample.Context) {
		start := time.Now()
		m.Logger.Info(ctx, "Request start")
		next(ctx)
		duration := time.Since(start)
		formattedMsg := fmt.Sprintf("Request complete (duration=%s)", duration)
		m.Logger.Info(ctx, formattedMsg)
	}
}
