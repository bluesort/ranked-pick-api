package api

import (
	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/go-chi/chi/v5"
)

func AddRoutes(cfg *config.AppConfig, router *chi.Mux) {
	Get(cfg, router, "/status").Handler(StatusHandler)
}
