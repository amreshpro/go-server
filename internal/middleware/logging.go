package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("[%s] %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next(w, r)

		log.Printf("Completed in %v", time.Since(start))
	}
}
