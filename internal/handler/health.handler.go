package handler

import (
	"encoding/json"
	"github.com/iambpn/go-http-template/internal/database"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type healthHandler struct{}

func (h *healthHandler) checkHealth(w http.ResponseWriter, r *http.Request) {
	db := database.NewDb()
	jsonResp, _ := json.Marshal(db.Health())
	_, _ = w.Write(jsonResp)
}

func (h *healthHandler) RegisterRoutes(router chi.Router) {
	router.Get("/db", h.checkHealth)
}

func NewHealthHandler() Handler {
	return &healthHandler{}
}
