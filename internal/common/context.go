package common

import (
	"context"

	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/carterjackson/ranked-pick-api/internal/jwt"
)

type Context struct {
	context.Context
	*config.AppConfig
	UserId int64
}

func NewContext(reqCtx context.Context) (*Context, error) {
	claims, err := jwt.ParseClaims(reqCtx)
	if err != nil {
		// TODO: return auth error
		return nil, err
	}

	// TODO: Read user?

	return &Context{
		Context:   reqCtx,
		AppConfig: config.Config,
		UserId:    claims.UserId,
	}, nil
}
