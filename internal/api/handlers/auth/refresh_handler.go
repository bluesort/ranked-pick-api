package auth

import (
	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
)

func RefreshHandler(ctx *common.Context, tx *db.Queries) (interface{}, error) {
	userId := ctx.UserId
	refreshCookie, err := ctx.Req.Cookie("jwt")
	if err != nil {
		return nil, err
	}
	if refreshCookie.Value == "" {
		return nil, errors.NewAuthError()
	}

	err = verifyRefreshToken(ctx, tx, refreshCookie.Value)
	if err != nil {
		return nil, err
	}

	accessToken, accessTokenExp, err := auth.NewAccessToken(userId)
	if err != nil {
		return nil, err
	}

	user, err := tx.ReadUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		AccessToken: &TokenResponse{
			Token: accessToken,
			Exp:   accessTokenExp,
		},
		User: newUserResp(&user),
	}, nil
}
