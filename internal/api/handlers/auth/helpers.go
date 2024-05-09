package auth

import (
	"errors"
	"time"

	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

type TokenResponse struct {
	Token string    `json:"token"`
	Exp   time.Time `json:"exp"`
}

type AuthResponse struct {
	AccessToken  *TokenResponse  `json:"access_token"`
	RefreshToken *TokenResponse  `json:"refresh_token"`
	User         *resources.User `json:"user"`
}

func newRefreshToken(ctx *common.Context, tx *db.Queries, userId int64) (string, time.Time, error) {
	token, expiresAt, err := auth.NewRefreshToken(userId)
	if err != nil {
		return "", time.Time{}, err
	}

	tokenHash, err := auth.Hash([]byte(token))
	if err != nil {
		return "", time.Time{}, err
	}

	_, err = tx.CreateTokenHash(ctx, db.CreateTokenHashParams{
		UserID:    userId,
		Hash:      string(tokenHash),
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiresAt, nil
}

func verifyRefreshToken(ctx *common.Context, tx *db.Queries, token string, userId int64) error {
	tokenHash, err := auth.Hash([]byte(token))
	if err != nil {
		return err
	}

	dbTokenHash, err := tx.ReadTokenHashByHash(ctx, string(tokenHash))
	if err != nil {
		return err
	}

	if dbTokenHash.UserID != userId {
		// TODO: auth error
		return errors.New("invalid user")
	}

	return nil
}
