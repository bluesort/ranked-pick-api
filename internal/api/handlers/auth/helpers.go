package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/carterjackson/ranked-pick-api/internal/auth"
	"github.com/carterjackson/ranked-pick-api/internal/common"
	"github.com/carterjackson/ranked-pick-api/internal/db"
	"github.com/carterjackson/ranked-pick-api/internal/env"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
	"github.com/carterjackson/ranked-pick-api/internal/resources"
)

type TokenResponse struct {
	Token string    `json:"token"`
	Exp   time.Time `json:"exp"`
}

type AuthResponse struct {
	User *resources.User `json:"user"`
}

func createRefreshToken(ctx *common.Context, tx *db.Queries) (string, time.Time, error) {
	token, expiresAt, err := auth.NewRefreshToken(ctx.UserId)
	if err != nil {
		return "", time.Time{}, err
	}

	tokenHash, err := auth.Hash(token)
	if err != nil {
		return "", time.Time{}, err
	}

	_, err = tx.UpsertTokenHash(ctx, db.UpsertTokenHashParams{
		UserID:    ctx.UserId,
		Hash:      string(tokenHash),
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return "", time.Time{}, err
	}

	return token, expiresAt, nil
}

func setRefreshToken(ctx *common.Context, tx *db.Queries, resp http.ResponseWriter) error {
	token, exp, err := createRefreshToken(ctx, tx)
	if err != nil {
		return err
	}
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  exp,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   env.GetBool("SECURE_STRICT", true),
		Path:     "/api/auth",
	}
	http.SetCookie(resp, &cookie)

	return nil
}

func verifyRefreshToken(ctx *common.Context, tx *db.Queries, token string) error {
	dbTokenHash, err := tx.ReadTokenHashByUserId(ctx, ctx.UserId)
	if err != nil {
		return errors.NewAuthError()
	}

	err = auth.VerifyPlainWithHash(token, dbTokenHash.Hash)
	if err != nil {
		log.Println(err)
		return errors.NewAuthError()
	}

	return nil
}
