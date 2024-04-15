package common

import (
	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

// TODO: Extend go context
type Context struct {
	Config *config.AppConfig
	User   *resources.User
}

func NewContext(config *config.AppConfig) *Context {
	return &Context{
		Config: config,
	}
}
