package common

import (
	"context"

	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/carterjackson/ranked-pick-api/internal/db"
)

type Context struct {
	context.Context
	Config *config.AppConfig
	User   *db.User
	Db     *db.Queries
}

func NewContext() *Context {
	return &Context{
		Context: context.Background(),
		Config:  config.Config,
		Db:      config.Config.Db,
	}
}
