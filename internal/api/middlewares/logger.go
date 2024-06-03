package middlewares

import (
	"log"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := newResponseWriter(w)
		log.Printf("%s: %s, %d", r.Method, r.RequestURI, rw.statusCode)
		next.ServeHTTP(w, r)
	})
}
