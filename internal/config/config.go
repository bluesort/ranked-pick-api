package config

import (
	"database/sql"

	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/go-chi/jwtauth/v5"
)

type AppConfig struct {
	Port    int
	Env     string
	Db      *sql.DB
	Queries *db.Queries
	Auth    *jwtauth.JWTAuth
}

var Config *AppConfig

func InitConfig() {
	Config = &AppConfig{
		Auth: jwtauth.New("HS256", []byte("secret"), nil), // TODO: Move secret to env
	}
	ParseFlags()
	PrepareDatabase()
}
