package auth

import (
	"database/sql"

	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
)

type SigninParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signin(ctx *common.Context, tx *db.Queries, iparams interface{}) (interface{}, error) {
	params := iparams.(*SigninParams)
	invalidCredsErr := errors.NewInputError("invalid username or password")

	user, err := tx.ReadUserByUsername(ctx, params.Username)
	if err == sql.ErrNoRows {
		return nil, invalidCredsErr
	} else if err != nil {
		return nil, err
	}

	err = auth.VerifyPlainWithHash(params.Password, user.PasswordHash)
	if err != nil {
		return nil, invalidCredsErr
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
