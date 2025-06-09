package fwsample

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type StdLogger struct{}

func (l *StdLogger) Info(w http.ResponseWriter, r *http.Request, msg string) {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	log.Printf("[INFO] timestamp=%v %s", timestamp, msg)
}

func (l *StdLogger) Error(w http.ResponseWriter, r *http.Request, msg string) {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	log.Printf("[ERROR] timestamp=%v %s", timestamp, msg)
}

type StdLoggerWithRequestID struct{}

func (l *StdLoggerWithRequestID) Info(w http.ResponseWriter, r *http.Request, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000Z07:00")
    requestID := w.Header().Get(HeaderXRequestID) 	
    fmt.Printf("[INFO] timestamp=%v request_id=%v %s\n", timestamp, requestID, msg)
}

func (l *StdLoggerWithRequestID) Error(w http.ResponseWriter, r *http.Request, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000Z07:00")
	requestID := w.Header().Get(HeaderXRequestID) 
	fmt.Printf("[ERROR] timestamp=%v request_id=%v %s\n", timestamp, requestID, msg)
}