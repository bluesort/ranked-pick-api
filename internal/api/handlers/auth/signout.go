package auth

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
)

func Signout(ctx *common.Context, tx *db.Queries) error {
	userId := ctx.UserId
	refreshCookie, err := ctx.Req.Cookie("jwt")
	if err != nil {
		return err
	}

	err = verifyRefreshToken(ctx, tx, refreshCookie.Value)
	if err != nil {
		return err
	}

	return tx.DeleteTokenHashByUserId(ctx, userId)
}
