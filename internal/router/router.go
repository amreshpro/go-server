package router

import (
	"net/http"

	"github.com/amreshpro/go-server/internal/handler"
	"github.com/amreshpro/go-server/internal/middleware"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// Apply middleware
	handlerWithMiddleware := middleware.LoggingMiddleware(handler.HealthCheck)
	mux.Handle("/api/health", handlerWithMiddleware)

	return mux
}
