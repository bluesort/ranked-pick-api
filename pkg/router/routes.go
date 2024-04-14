package router

import (
	"github.com/go-chi/chi/v5"
)

func AddRoutes(router *chi.Mux) {
	Get(router, "/status").Handler(StatusHandler)
}
