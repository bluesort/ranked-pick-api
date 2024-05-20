package users

import (
	"database/sql"

	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
)

type UpdateParams struct {
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
}

func Update(ctx *common.Context, tx *db.Queries, id int64, iparams interface{}) (interface{}, error) {
	params := iparams.(*UpdateParams)

	if ctx.UserId != id {
		return nil, errors.NewAuthError()
	}

	existingUser, err := ctx.Queries.ReadUser(ctx, id)
	if err != nil {
		return nil, err
	}

	if params.Username != "" && params.Username != existingUser.Username {
		_, err = ctx.Queries.ReadUserByUsername(ctx, params.Username)
		if err == nil {
			return nil, errors.NewInputError("username unavailable")
		} else if err != sql.ErrNoRows {
			return nil, err
		}
	}

	dbUser, err := tx.UpdateUser(ctx, db.UpdateUserParams{
		ID:          ctx.UserId,
		Username:    db.NewNullString(params.Username),
		DisplayName: db.NewNullString(params.DisplayName),
	})
	if err != nil {
		return nil, err
	}

	return db.NewUser(&dbUser), nil
}
