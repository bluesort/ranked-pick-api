package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carterjackson/ranked-pick-api/internal/api"
	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("Starting ranked-pick-api")

	cfg := config.ParseConfig()
	config.PrepareDatabase()

	r := chi.NewRouter()
	api.AddMiddleware(cfg, r)
	api.AddRoutes(cfg, r)

	log.Printf("Router listening on port %d\n", cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), r)
	if err != nil {
		log.Println(err)
	}
}
