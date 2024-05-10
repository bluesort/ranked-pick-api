package auth

import (
	"errors"
	"net/http"
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
	AccessToken *TokenResponse  `json:"access_token"`
	User        *resources.User `json:"user"`
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

func setRefreshToken(ctx *common.Context, tx *db.Queries, resp http.ResponseWriter, userId int64) error {
	token, exp, err := newRefreshToken(ctx, tx, userId)
	if err != nil {
		return err
	}
	resp.Header().Set("Set-Cookie", "refresh_token="+token+"; Path=/auth/refresh; Expires="+exp.Format(time.RFC1123)+"; HttpOnly")

	return nil
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
