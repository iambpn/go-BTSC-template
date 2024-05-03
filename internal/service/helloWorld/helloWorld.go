package helloWorld

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iambpn/go-http-template/internal/config"
	"github.com/iambpn/go-http-template/internal/helper"
)

type helloWorld struct {
	logger    *log.Logger
	appConfig config.AppConfig
}

func New(
	a config.AppConfig,
	logger *log.Logger,
) *helloWorld {
	return &helloWorld{appConfig: a, logger: logger}
}

// TODO: implement error validation

func (h *helloWorld) SayHelloWorld() http.Handler {
	type response struct {
		Greeting string `json:"greeting"`
	}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			h.logger.Println("Inside of Hello world service: " + r.URL.Path)

			name := r.URL.Query().Get("name")

			helper.JsonEncode(w, r, http.StatusInternalServerError, response{
				Greeting: fmt.Sprintf("Hello %s from %s:%s", name, h.appConfig.Host, h.appConfig.Port),
			})
		},
	)
}
