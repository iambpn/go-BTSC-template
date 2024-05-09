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

func (h *helloWorld) SayHelloWorld() http.Handler {
	type response struct {
		Greeting string `json:"greeting"`
	}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			h.logger.Println("Inside of Hello world service: " + r.URL.Path)

			q := Query{
				name: r.URL.Query().Get("name"),
			}

			problems := q.Valid(r.Context())

			if len(problems) > 0 {
				errorDetails := []string{}

				for _, v := range problems {
					errorDetails = append(errorDetails, v)
				}

				helper.JsonEncodeError(w, r, http.StatusBadRequest, helper.HttpError{
					Message: errorDetails,
				})
				return
			}

			helper.JsonEncode(w, r, http.StatusInternalServerError, response{
				Greeting: fmt.Sprintf("Hello %s from %s:%s", q.name, h.appConfig.Host, h.appConfig.Port),
			})
		},
	)
}
