package auth

import (
	"database/sql"

	"github.com/carterjackson/ranked-pick-api/internal/api/errors"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/jwt"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

type SignupParams struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"displayName"`
}

func SignupHandler(ctx *common.Context, tx *db.Queries, iparams interface{}) (interface{}, error) {
	params := iparams.(*SignupParams)

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

	passwordHash, err := HashPassword([]byte(params.Password))
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

	token, err := jwt.NewToken(user.ID)
	if err != nil {
		return nil, err
	}

	return token, nil
}
