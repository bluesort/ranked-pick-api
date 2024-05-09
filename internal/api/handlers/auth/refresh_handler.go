package auth

import (
	"fmt"

	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
)

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

func RefreshHandler(ctx *common.Context, tx *db.Queries, iparams interface{}) (interface{}, error) {
	refreshCookie, err := ctx.Req.Cookie("refresh_token")
	if err != nil {
		return nil, err
	}

	fmt.Println(refreshCookie)

	// err = verifyRefreshToken(ctx, tx, params.RefreshToken)
	// if err != nil {
	// 	// TODO: return unauth error
	// 	return err, nil
	// }

	// issue new access token

	return nil, nil
	// return &RefreshResponse{
	// 	AccessToken: accessToken,
	// }, nil
}
