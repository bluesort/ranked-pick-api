package api

import (
	auth_handlers "github.com/carterjackson/ranked-pick-api/internal/api/handlers/auth"
	"github.com/carterjackson/ranked-pick-api/internal/api/handlers/surveys"
	"github.com/carterjackson/ranked-pick-api/internal/api/handlers/users"
	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	// TODO: Add rate limiting
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	Get(router, "/status").Handler(func(ctx *common.Context) (interface{}, error) {
		return "ready", nil
	})

	// Auth
	Post(router, "/auth/signup").Handler(auth_handlers.Signup, &auth_handlers.SignupParams{})
	Post(router, "/auth/signin").Handler(auth_handlers.Signin, &auth_handlers.SigninParams{})
	router.Group(func(refreshRouter chi.Router) {
		auth.AddRefreshTokenMiddleware(refreshRouter)
		Post(refreshRouter, "/auth/refresh").Handler(auth_handlers.Refresh)
		Post(refreshRouter, "/auth/signout").Handler(auth_handlers.Signout)
	})

	// Protected Routes
	router.Group(func(authedRouter chi.Router) {
		auth.AddAccessTokenMiddleware(authedRouter)

		// Surveys
		Post(authedRouter, "/surveys").Handler(surveys.Create, &surveys.CreateParams{})
		Get(authedRouter, "/surveys").Handler(surveys.List)
		Get(authedRouter, "/surveys/{id}").Handler(surveys.Read)
		Post(authedRouter, "/surveys/{id}").Handler(surveys.Update, &surveys.UpdateParams{})
		Delete(authedRouter, "/surveys/{id}").Handler(surveys.Delete)
		Get(authedRouter, "/surveys/{id}/options").Handler(surveys.ListOptions)
		Get(authedRouter, "/surveys/{id}/results").Handler(surveys.Results)
		Post(authedRouter, "/surveys/{id}/vote").Handler(surveys.Vote, &surveys.VoteParams{})
		Get(authedRouter, "/surveys/{id}/responses").Handler(surveys.ListResponses, &surveys.ListResponsesParams{})

		// Users
		Put(authedRouter, "/users/{id}").Handler(users.Update, &users.UpdateParams{})
		Delete(authedRouter, "/users/{id}").Handler(users.Delete)
		Get(authedRouter, "/users/{id}/created_surveys").Handler(users.ListCreatedSurveys)
		Get(authedRouter, "/users/{id}/responded_surveys").Handler(users.ListRespondedSurveys)
	})

	return router
}
