package api

import (
	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/resources/surveys"
	"github.com/go-chi/chi/v5"
)

// TODO: move handlers to api/handlers package

func AddRoutes(router *chi.Mux) {
	Get(router, "/status").Handler(func(ctx *common.Context) (interface{}, error) {
		return "ready", nil
	})

	// Auth
	Post(router, "/signin").Handler(auth.SigninHandler, &auth.SigninParams{})

	// Surveys
	Post(router, "/surveys").Handler(surveys.Create, &surveys.CreateParams{})
	Get(router, "/surveys").Handler(surveys.List)
	Get(router, "/surveys/{id}").Handler(surveys.Read)
	Post(router, "/surveys/{id}").Handler(surveys.Update, &surveys.UpdateParams{})
	Delete(router, "/surveys/{id}").Handler(surveys.Delete)
}
