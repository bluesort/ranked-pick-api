package api

import (
	"github.com/carterjackson/ranked-pick-api/internal/resources/surveys"
	"github.com/go-chi/chi/v5"
)

func AddRoutes(router *chi.Mux) {
	Get(router, "/status").Handler(StatusHandler)

	// Surveys
	Get(router, "/surveys").Handler(surveys.List)
	Post(router, "/surveys").Handler(surveys.Create, &surveys.CreateParams{})
	Post(router, "/surveys/{id}").Handler(surveys.Update, &surveys.UpdateParams{})
	Get(router, "/surveys/{id}").Handler(surveys.Get)
	Delete(router, "/surveys/{id}").Handler(surveys.Delete)
}
