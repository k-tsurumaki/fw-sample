package fwsample

import (
	"log"
	"fmt"
	"time"
)

type StdLogger struct{}

func (l *StdLogger) Info(ctx Context, msg string) {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	log.Printf("[INFO] timestamp=%v %s", timestamp, msg)
}

func (l *StdLogger) Error(ctx Context, msg string) {
	timestamp := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
	log.Printf("[ERROR] timestamp=%v %s", timestamp, msg)
}

type StdLoggerWithRequestID struct{}

func (l *StdLoggerWithRequestID) Info(ctx Context, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000Z07:00")
    requestID := ctx.ResponseWriter().Header().Get(HeaderXRequestID) 	
    fmt.Printf("[INFO] timestamp=%v request_id=%v %s\n", timestamp, requestID, msg)
}

func (l *StdLoggerWithRequestID) Error(ctx Context, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000Z07:00")
	requestID := ctx.ResponseWriter().Header().Get(HeaderXRequestID) 
	fmt.Printf("[ERROR] timestamp=%v request_id=%v %s\n", timestamp, requestID, msg)
}