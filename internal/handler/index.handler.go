package handler

import (
	"github.com/iambpn/go-http-template/cmd/web/view/page"
	"github.com/iambpn/go-http-template/internal/logger"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
)

type indexHandler struct {
	// service *Service.Service
}

func (h *indexHandler) postIndexHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		logger.Log.Error("Error while parsing request", err)
	}

	body := r.FormValue("name")

	w.Write([]byte(body))
}

func (h *indexHandler) RegisterRoutes(router chi.Router) {
	// View Routes
	router.Get("/", templ.Handler(page.HelloForm()).ServeHTTP)

	// Handler routes
	router.Post("/", h.postIndexHandler)
}

func NewIndexHandler() Handler {
	return &indexHandler{}
}
