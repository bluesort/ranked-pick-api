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
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	DisplayName          string `json:"display_name"`
	AcceptedTos          bool   `json:"accepted_tos"`
}

func Signup(ctx *common.Context, tx *db.Queries, iparams interface{}) (interface{}, error) {
	params := iparams.(*SignupParams)

	if params.Password != params.PasswordConfirmation {
		return nil, errors.NewInputError("password confirmation does not match")
	}

	if !params.AcceptedTos {
		return nil, errors.NewInputError("you must accept the terms of service")
	}

	err := resources.ValidateUsername(params.Username)
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

	_, err = tx.ReadUserByUsername(ctx, params.Username)
	if err == nil {
		return nil, errors.NewInputError("username already in use")
	} else if err != sql.ErrNoRows {
		return nil, err
	}

	passwordHash, err := auth.Hash(params.Password)
	if err != nil {
		return nil, err
	}

	user, err := tx.CreateUser(ctx, db.CreateUserParams{
		Username:     params.Username,
		DisplayName:  db.NewNullString(params.DisplayName),
		PasswordHash: passwordHash,
	})
	if err != nil {
		return nil, err
	}

	ctx.UserId = user.ID
	err = setRefreshToken(ctx, tx, ctx.Resp)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		User: db.NewUser(&user),
	}, nil
}
