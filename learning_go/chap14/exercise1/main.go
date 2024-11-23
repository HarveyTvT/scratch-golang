package main

import (
	"context"
	"net/http"
	"time"
)

func TimedHandler(timeoutMilli int64) func(http.Handler) http.Handler {
	return func(r http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			ctx, cancel := context.WithTimeout(ctx, time.Duration(timeoutMilli)*time.Millisecond)
			defer cancel()
			req = req.WithContext(ctx)
			r.ServeHTTP(w, req)
		})

	}
}
