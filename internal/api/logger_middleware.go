package api

import (
	"net/http"
	"time"

	"github.com/stpiech/gosolve-task/internal/logger"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger.InfoLogger("Incoming request: " + r.Method + " " + r.URL.String())

		ww := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(ww, r)

		logger.DebugLogger("Response status: " + http.StatusText(ww.status))
		logger.InfoLogger("Request processed in: " + time.Since(start).String())
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}
