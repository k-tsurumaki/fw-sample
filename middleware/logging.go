package middleware

import (
	"fmt"
	"net/http"
	"time"

	fwsample "github.com/k-tsurumaki/fw-sample"
)

type Logger interface {
	Info(w http.ResponseWriter, r *http.Request, msg string)
	Error(w http.ResponseWriter, r *http.Request, msg string)
}

type LoggingMiddleware struct {
	Logger Logger
}

func (m *LoggingMiddleware) Logging(next fwsample.HandlerFunc) fwsample.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		m.Logger.Info(w, r, "Request start")
		next(w, r)
		duration := time.Since(start)
		formattedMsg := fmt.Sprintf("Request complete (duration=%s)", duration)
		m.Logger.Info(w, r, formattedMsg)
	}
}