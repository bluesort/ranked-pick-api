package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/carterjackson/ranked-pick-api/internal/api"
	"github.com/carterjackson/ranked-pick-api/internal/config"
)

func main() {
	log.Println("Starting ranked-pick-api")
	config.InitConfig()
	r := api.NewRouter()

	log.Printf("Router listening on port %d\n", config.Config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), r)
	if err != nil {
		log.Println(err)
	}
}
