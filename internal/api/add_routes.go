package api

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/resources/surveys"
	"github.com/go-chi/chi/v5"
)

func AddRoutes(router *chi.Mux) {
	Get(router, "/status").Handler(func(ctx *common.Context) (interface{}, error) {
		return "ready", nil
	})

	// Surveys
	Post(router, "/surveys").Handler(surveys.Create, &surveys.CreateParams{})
	Get(router, "/surveys").Handler(surveys.List)
	Get(router, "/surveys/{id}").Handler(surveys.Read)
	Post(router, "/surveys/{id}").Handler(surveys.Update, &surveys.UpdateParams{})
	Delete(router, "/surveys/{id}").Handler(surveys.Delete)
}
