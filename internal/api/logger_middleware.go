package api

import (
	"net/http"
	"time"

	"github.com/stpiech/gosolve-task/internal/logger"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger.InfoLogger("Incoming request: " + r.Method + " " + r.URL.String())

		ww := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(ww, r)

		logger.DebugLogger("Response status: " + http.StatusText(ww.status))
		logger.DebugLogger("Request processed in: " + time.Since(start).String())
	})
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
