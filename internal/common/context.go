package common

import (
	"context"

	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/carterjackson/ranked-pick-api/internal/db"
)

type Context struct {
	context.Context
	*config.AppConfig
	User *db.User
}

func NewContext() *Context {
	return &Context{
		Context:   context.Background(),
		AppConfig: config.Config,
	}
}
