package controller

import (
	"log"
	"net/http"

	"github.com/iambpn/go-http-template/internal/config"
	"github.com/iambpn/go-http-template/internal/service/helloWorld"
)

func New(mux *http.ServeMux,
	logger *log.Logger,
	config config.AppConfig) {

	hw := helloWorld.New(config, logger)

	mux.Handle("GET /", hw.SayHelloWorld())
}
