package config

import (
	"database/sql"
	"flag"

	"github.com/carterjackson/ranked-pick-api/internal/db"
)

type AppConfig struct {
	Port    int
	Env     string
	Db      *sql.DB
	Queries *db.Queries
}

var Config *AppConfig

func ParseConfig() {
	Config = &AppConfig{}
	flag.IntVar(&Config.Port, "port", 3000, "Port for the server to listen on")
	flag.StringVar(&Config.Env, "env", "development", "Environment the server is being run in, e.g. 'development'")
	flag.Parse()
}
