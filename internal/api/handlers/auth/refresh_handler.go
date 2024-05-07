package auth

import (
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
)

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

func RefreshHandler(ctx *common.Context, tx *db.Queries, iparams interface{}) (interface{}, error) {
	// read refresh token from req

	err := verifyRefreshToken(ctx, tx, params.RefreshToken)
	if err != nil {
		// TODO: return unauth error
		return err, nil
	}

	// issue new access token

	return &RefreshResponse{
		AccessToken: accessToken,
	}, nil
}
