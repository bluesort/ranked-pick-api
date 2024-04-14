package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carterjackson/ranked-pick-api/pkg/config"
	"github.com/carterjackson/ranked-pick-api/pkg/db"
	"github.com/carterjackson/ranked-pick-api/pkg/router"
	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("Starting ranked-pick-api")

	cfg := config.ParseConfig()
	db.PrepareDatabase()

	r := chi.NewRouter()
	router.AddMiddleware(r)
	router.AddRoutes(r)

	log.Printf("Router listening on port %d\n", cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
	if err != nil {
		log.Println(err)
	}
}
