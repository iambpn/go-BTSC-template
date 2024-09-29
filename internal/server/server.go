package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/iambpn/go-http-template/internal/database"
	"github.com/iambpn/go-http-template/internal/route"
)

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	// initialize database
	database.NewDb()

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      route.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
