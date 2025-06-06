package fwsample

import (
	"time"
)

type Config struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

var DefaultConfig = Config{
	Addr: ":8080",
	ReadTimeout: 5 *time.Second,
	WriteTimeout: 10 *time.Second,
	IdleTimeout: 120 *time.Second,
}