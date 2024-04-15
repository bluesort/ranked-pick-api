package api

import (
	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AddMiddleware(cfg *config.AppConfig, router *chi.Mux) {
	// TODO: CORS, rate limit, timeout
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
}
