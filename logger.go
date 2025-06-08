package fwsample

import (
	"log"
)

type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, args ...any)
}

type StdLogger struct{}

func (l *StdLogger) Info(msg string, args ...any) {
	log.Printf("[INFO] "+msg, args...)
}

func (l *StdLogger) Error(msg string, args ...any) {
	log.Printf("[ERROR] "+msg, args...)
}
