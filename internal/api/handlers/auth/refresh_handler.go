package auth

import (
	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

func RefreshHandler(ctx *common.Context, tx *db.Queries) (interface{}, error) {
	userId := ctx.UserId
	refreshCookie, err := ctx.Req.Cookie("jwt")
	if err != nil {
		return nil, err
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

	userResp := resources.NewUser(user)
	return &AuthResponse{
		AccessToken: &TokenResponse{
			Token: accessToken,
			Exp:   accessTokenExp,
		},
		User: &userResp,
	}, nil
}
