package fwsample

import (
	"log"
)

func Logger(next HandlerFunc) HandlerFunc {
	return func(ctx Context) {
		log.Printf("Request: %s %s", ctx.Request().Method, ctx.Request().URL.Path)
		next(ctx)
	}
}
