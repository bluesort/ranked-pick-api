package common

import (
	"context"

	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/jwt"
)

type Context struct {
	context.Context
	*config.AppConfig
	User   *db.User
	Claims *jwt.Claims
}

func NewContext(reqCtx context.Context) (*Context, error) {
	claims, err := jwt.ParseClaims(reqCtx)
	if err != nil {
		return nil, err
	}

	return &Context{
		Context:   reqCtx,
		AppConfig: config.Config,
		Claims:    claims,
	}, nil
}
