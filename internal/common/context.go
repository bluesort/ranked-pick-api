package common

import (
	"context"

	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/config"
)

type Context struct {
	context.Context
	*config.AppConfig
	UserId int64
}

func NewContext(reqCtx context.Context) (*Context, error) {
	claims, err := auth.ParseAccessClaims(reqCtx)
	if err != nil {
		// TODO: return auth error
		return nil, err
	}

	return &Context{
		Context:   reqCtx,
		AppConfig: config.Config,
		UserId:    claims.UserId,
	}, nil
}
