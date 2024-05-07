package auth

import (
	"database/sql"

	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

type SignupParams struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	DisplayName          string `json:"display_name"`
	AcceptedTos          bool   `json:"accepted_tos"`
}

// TODO: Confirm email
func SignupHandler(ctx *common.Context, tx *db.Queries, iparams interface{}) (interface{}, error) {
	params := iparams.(*SignupParams)

	if params.Password != params.PasswordConfirmation {
		return nil, errors.NewInputError("password confirmation does not match")
	}

	if !params.AcceptedTos {
		return nil, errors.NewInputError("you must accept the terms of service")
	}

	err := resources.ValidateEmail(params.Email)
	if err != nil {
		return nil, err
	}
	err = resources.ValidatePassword(params.Password)
	if err != nil {
		return nil, err
	}
	if params.DisplayName != "" {
		err = resources.ValidateDisplayName(params.DisplayName)
		if err != nil {
			return nil, err
		}
	}

	_, err = tx.ReadUserByEmail(ctx, params.Email)
	if err == nil {
		return nil, errors.NewInputError("email already in use")
	} else if err != sql.ErrNoRows {
		return nil, err
	}

	passwordHash, err := auth.Hash([]byte(params.Password))
	if err != nil {
		return nil, err
	}

	user, err := tx.CreateUser(ctx, db.CreateUserParams{
		Email:        params.Email,
		DisplayName:  db.NewNullString(params.DisplayName),
		PasswordHash: string(passwordHash),
	})
	if err != nil {
		return nil, err
	}

	accessToken, err := auth.NewAccessToken(user.ID)
	if err != nil {
		return nil, err
	}

	refreshToken, err := newRefreshToken(tx)
	if err != nil {
		return nil, err
	}

	userResp := resources.NewUser(user)
	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         &userResp,
	}, nil
}
