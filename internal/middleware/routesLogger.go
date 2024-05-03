package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Before %s", r.URL.String())
		next.ServeHTTP(w, r)
		logger.Printf("After %s", r.URL.String())
	})
}
