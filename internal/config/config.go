package config

import (
	"database/sql"

	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/go-chi/jwtauth/v5"
)

type AppConfig struct {
	Port             int
	Env              string
	Db               *sql.DB
	Queries          *db.Queries
	AccessTokenAuth  *jwtauth.JWTAuth
	RefreshTokenAuth *jwtauth.JWTAuth
}

var Config *AppConfig

func InitConfig() {
	Config = &AppConfig{
		AccessTokenAuth:  jwtauth.New("HS256", []byte("secret1"), nil), // TODO: Move secret to env
		RefreshTokenAuth: jwtauth.New("HS256", []byte("secret2"), nil), // TODO: Move secret to env
	}
	ParseFlags()
	PrepareDatabase()
}
