package server

import (
	"log"
	"net/http"

	"github.com/iambpn/go-http-template/internal/config"
	"github.com/iambpn/go-http-template/internal/router"
)

func NewServer(
	logger *log.Logger,
	config *config.AppConfig,
) http.Handler {
	mux := http.NewServeMux()

	router.AddRoutes(
		mux,
		logger,
		*config,
	)

	// using polymorphism to get http.Handler
	var handler http.Handler = mux

	// Add global middlewares Here
	// handler = someMiddleware(handler)

	return handler
}
