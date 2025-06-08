package middleware

import (
	fwsample "github.com/k-tsurumaki/fw-sample"
)

func RequestID(next fwsample.HandlerFunc) fwsample.HandlerFunc {
	return func(ctx fwsample.Context) {
		requestID := ctx.Request().Header.Get(fwsample.HeaderXRequestID)
		if requestID == "" {
			requestID = generator()
		}

		ctx.ResponseWriter().Header().Set(fwsample.HeaderXRequestID, requestID)
		next(ctx)
	}
}

func generator() string {
	return randomString(32)
}
