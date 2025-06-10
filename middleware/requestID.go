package middleware

import (
	"context"
	"net/http"

	fwsample "github.com/k-tsurumaki/fw-sample"
)

type contextKey string

const requestIDKey contextKey = fwsample.HeaderXRequestID

func RequestID(next fwsample.HandlerFunc) fwsample.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(fwsample.HeaderXRequestID)

		// if the request doesn't have request-ID
		// create new request-ID
		if requestID == "" {
			requestID = generator()
		}

		// set request-ID to response header
		w.Header().Set(fwsample.HeaderXRequestID, requestID)

		// set request-ID to context
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		next(w, r.WithContext(ctx))
	}
}

func generator() string {
	return randomString(32)
}