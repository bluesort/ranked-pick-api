package common

import (
	"context"
	"log"
	"net/http"

	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
)

type Context struct {
	context.Context
	*config.AppConfig
	Req    *http.Request
	Resp   http.ResponseWriter
	UserId int64
}

func NewContext(w http.ResponseWriter, r *http.Request) (*Context, error) {
	reqCtx := r.Context()
	claims, err := auth.ParseClaims(reqCtx)
	if err != nil {
		log.Printf("Error parsing JWT claims: %s", err)
		return nil, errors.NewAuthError()
	}

	return &Context{
		Req:       r,
		Resp:      w,
		Context:   reqCtx,
		AppConfig: config.Config,
		UserId:    claims.UserId,
	}, nil
}
