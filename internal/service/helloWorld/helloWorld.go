package helloWorld

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iambpn/go-http-template/internal/config"
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

func (h *helloWorld) SayHelloWorld() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			h.logger.Println("Inside of Hello world service: " + r.URL.Path)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("Hello World from %s:%s", h.appConfig.Host, h.appConfig.Port)))
		},
	)
}
