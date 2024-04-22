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

	config.ParseConfig()
	config.PrepareDatabase()

	r := chi.NewRouter()
	api.AddMiddleware(r)
	api.AddRoutes(r)

	log.Printf("Router listening on port %d\n", config.Config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), r)
	if err != nil {
		log.Println(err)
	}
}
