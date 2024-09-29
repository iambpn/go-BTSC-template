package handler

import (
	"github.com/go-chi/chi/v5"
)

type Handler interface {
	RegisterRoutes(r chi.Router)
}
