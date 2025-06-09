package middleware

import (
	"net/http"

	fwsample "github.com/k-tsurumaki/fw-sample"
)

func RequestID(next fwsample.HandlerFunc) fwsample.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(fwsample.HeaderXRequestID)
		if requestID == "" {
			requestID = generator()
		}

		w.Header().Set(fwsample.HeaderXRequestID, requestID)
		next(w, r)
	}
}

func generator() string {
	return randomString(32)
}