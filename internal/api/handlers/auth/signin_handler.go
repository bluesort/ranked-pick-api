package auth

import (
	"database/sql"

	"github.com/carterjackson/ranked-pick-api/internal/api/errors"
	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
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

	err = auth.VerifyPassword([]byte(user.PasswordHash), []byte(params.Password))
	if err != nil {
		return nil, invalidCredsErr
	}

	token, err := auth.NewToken(user.ID)
	if err != nil {
		return nil, err
	}

	return token, nil
}
