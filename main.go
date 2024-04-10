package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carterjackson/ranked-choice/api"
	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	log.Println("Starting ranked-choice API")

	cfg := api.ParseConfig()

	// TODO: Move DB URL into env var, use postgres role, and enable SSL
	log.Println("Running migrations...")
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://localhost:5432/development?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err == migrate.ErrNoChange {
		log.Println("No migrations to be run")
	} else if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Migrations run")
	}

	// TODO: Rate limiting
	log.Println("Creating routes...")
	baseRouter := chi.NewRouter()
	apiRouter := chi.NewRouter()
	api.AddMiddleware(apiRouter)
	api.AddRoutes(apiRouter)
	baseRouter.Mount("/api", apiRouter)
	log.Println("Routes created")

	log.Printf("Listening on port %d\n", cfg.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), baseRouter)
	if err != nil {
		log.Println(err)
	}
}
