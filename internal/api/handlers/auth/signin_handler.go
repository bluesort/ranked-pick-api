package auth

import (
	"database/sql"

	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

type SigninParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SigninHandler(ctx *common.Context, tx *db.Queries, iparams interface{}) (interface{}, error) {
	params := iparams.(*SigninParams)
	invalidCredsErr := errors.NewInputError("invalid email or password")

	user, err := tx.ReadUserByEmail(ctx, params.Email)
	if err == sql.ErrNoRows {
		return nil, invalidCredsErr
	} else if err != nil {
		return nil, err
	}

	err = auth.VerifyPassword(user.PasswordHash, params.Password)
	if err != nil {
		return nil, invalidCredsErr
	}

	accessToken, accessTokenExp, err := auth.NewAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, refreshTokenExp, err := newRefreshToken(ctx, tx, user.ID)
	if err != nil {
		return nil, err
	}

	userResp := resources.NewUser(user)
	return &AuthResponse{
		AccessToken: &TokenResponse{
			Token: accessToken,
			Exp:   accessTokenExp,
		},
		RefreshToken: &TokenResponse{
			Token: refreshToken,
			Exp:   refreshTokenExp,
		},
		User: &userResp,
	}, nil
}
