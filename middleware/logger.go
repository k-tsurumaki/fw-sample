package middleware

import (
	"fmt"
	"net/http"
	"time"
	"log"

	fwsample "github.com/k-tsurumaki/fw-sample"
)

type Logger interface {
	Info(w fwsample.ResponseWriterWithStatus, r *http.Request, msg string)
	Error(w fwsample.ResponseWriterWithStatus, r *http.Request, msg string)
}

type LoggingMiddleware struct {
	Logger Logger
}

func (m *LoggingMiddleware) Logging(next fwsample.HandlerFunc) fwsample.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wrappedWriter := &fwsample.ResponseWriterWithStatus{
			ResponseWriter: w,
			StatusCode:     200,
		}

		start := time.Now()
		m.Logger.Info(*wrappedWriter, r, "Request start")

		next(wrappedWriter, r)

		duration := time.Since(start)
		formattedMsg := fmt.Sprintf("Request complete (duration=%s)", duration)

		if wrappedWriter.StatusCode >= 500 {
			m.Logger.Error(*wrappedWriter, r, formattedMsg)
		} else {
			m.Logger.Info(*wrappedWriter, r, formattedMsg)
		}
	}
}

type StdLogger struct{}

func (l *StdLogger) Info(w fwsample.ResponseWriterWithStatus, r *http.Request, msg string) {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	log.Printf("[INFO] timestamp=%v %s", timestamp, msg)
}

func (l *StdLogger) Error(w fwsample.ResponseWriterWithStatus, r *http.Request, msg string) {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	log.Printf("[ERROR] timestamp=%v %s", timestamp, msg)
}

type StdLoggerWithRequestID struct{}

func (l *StdLoggerWithRequestID) Info(w fwsample.ResponseWriterWithStatus, r *http.Request, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000Z07:00")
	requestID := r.Context().Value(requestIDKey)
	fmt.Printf("[INFO] timestamp=%v request_id=%v code=%v %s\n", timestamp, requestID, w.StatusCode, msg)
}

func (l *StdLoggerWithRequestID) Error(w fwsample.ResponseWriterWithStatus, r *http.Request, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000Z07:00")
	requestID := r.Context().Value(requestIDKey)
	fmt.Printf("[ERROR] timestamp=%v request_id=%v code-%v %s\n", timestamp, requestID, w.StatusCode, msg)
}

