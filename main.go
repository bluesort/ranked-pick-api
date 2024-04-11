package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carterjackson/ranked-choice/api"
	"github.com/carterjackson/ranked-choice/db"
	"github.com/go-chi/chi/v5"
)

func main() {
	log.Println("Starting ranked-choice API")

	cfg := api.ParseConfig()
	db.InitDb()

	router := chi.NewRouter()
	api.AddMiddleware(router)
	api.AddRoutes(router)

	log.Printf("Router listening on port %d\n", cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
	if err != nil {
		log.Println(err)
	}
}
