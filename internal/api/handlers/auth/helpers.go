package auth

import (
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

func newUserResp(user *db.User) *resources.User {
	return &resources.User{
		Id:          user.ID,
		Email:       user.Email,
		DisplayName: user.DisplayName.String,
	}
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
		Secure:   false, // TODO: set to false before deployment
		Path:     "/api/auth/refresh",
	}
	http.SetCookie(resp, &cookie)

	return nil
}

func verifyRefreshToken(ctx *common.Context, tx *db.Queries, token string) error {
	dbTokenHash, err := tx.ReadTokenHashByUserId(ctx, ctx.UserId)
	if err != nil {
		return err
	}

	err = auth.VerifyPlainWithHash(token, dbTokenHash.Hash)
	if err != nil {
		return err
	}

	return nil
}
