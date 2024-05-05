package api

import (
	"github.com/carterjackson/ranked-pick-api/internal/api/handlers/auth"
	"github.com/carterjackson/ranked-pick-api/internal/api/handlers/surveys"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/jwt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	Get(router, "/status").Handler(func(ctx *common.Context) (interface{}, error) {
		return "ready", nil
	})

	// Auth
	Post(router, "/signup").Handler(auth.SignupHandler, &auth.SignupParams{})
	Post(router, "/signin").Handler(auth.SigninHandler, &auth.SigninParams{})

	// Protected Routes
	router.Group(func(authedRouter chi.Router) {
		jwt.AddMiddleware(authedRouter)

		// Surveys
		Post(router, "/surveys").Handler(surveys.Create, &surveys.CreateParams{})
		Get(router, "/surveys").Handler(surveys.List)
		Get(router, "/surveys/{id}").Handler(surveys.Read)
		Post(router, "/surveys/{id}").Handler(surveys.Update, &surveys.UpdateParams{})
		Delete(router, "/surveys/{id}").Handler(surveys.Delete)
	})

	return router
}
