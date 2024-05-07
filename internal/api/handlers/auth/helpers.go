package auth

import (
	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

type AuthResponse struct {
	AccessToken  string          `json:"access_token"`
	RefreshToken string          `json:"refresh_token"`
	User         *resources.User `json:"user"`
}

func newRefreshToken(ctx *common.Context, tx *db.Queries, userId int64) (string, error) {
	token, expiresAt, err := auth.NewRefreshToken(userId)
	if err != nil {
		return "", err
	}

	tokenHash, err := auth.Hash([]byte(token))
	if err != nil {
		return "", err
	}

	_, err = tx.CreateTokenHash(ctx, db.CreateTokenHashParams{
		UserID:    userId,
		Hash:      string(tokenHash),
		ExpiresAt: *expiresAt,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func verifyRefreshToken(ctx *common.Context, tx *db.Queries, token string) error {
	tokenHash, err := auth.Hash([]byte(token))
	if err != nil {
		return err
	}

	_, err = tx.ReadTokenHashByHash(ctx, string(tokenHash))
	if err != nil {
		return err
	}

	return nil
}
