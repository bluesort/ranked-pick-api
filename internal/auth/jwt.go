package auth

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/carterjackson/ranked-pick-api/internal/errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type Claims struct {
	UserId int64     `json:"user_id"`
	Exp    time.Time `json:"exp"`
	Iat    time.Time `json:"iat"`
}

var AccessTokenTTL = 15 * time.Minute
var RefreshTokenTTL = 24 * time.Hour * 7

func ParseClaims(ctx context.Context) (*Claims, error) {
	_, claimsMap, _ := jwtauth.FromContext(ctx)

	var claims Claims
	encodedClaims, err := json.Marshal(claimsMap)
	if err != nil {
		log.Printf("Error parsing JWT claims: %s", err)
		return nil, errors.NewAuthError()
	}
	err = json.Unmarshal(encodedClaims, &claims)
	if err != nil {
		log.Printf("Error parsing JWT claims: %s", err)
		return nil, errors.NewAuthError()
	}

	return &claims, nil
}

func AddAccessTokenMiddleware(router chi.Router) {
	addMiddleware(router, config.Config.AccessTokenAuth)
}

func AddRefreshTokenMiddleware(router chi.Router) {
	addMiddleware(router, config.Config.RefreshTokenAuth)
}

func NewAccessToken(userId int64) (string, time.Time, error) {
	return newToken(config.Config.AccessTokenAuth, AccessTokenTTL, userId)
}

func NewRefreshToken(userId int64) (string, time.Time, error) {
	return newToken(config.Config.RefreshTokenAuth, RefreshTokenTTL, userId)
}

func addMiddleware(router chi.Router, auth *jwtauth.JWTAuth) {
	// Seek, verify and validate JWT tokens
	router.Use(jwtauth.Verifier(auth))

	// Handle valid/invalid tokens
	router.Use(jwtauth.Authenticator(auth))
}

func newToken(auth *jwtauth.JWTAuth, ttl time.Duration, userId int64) (string, time.Time, error) {
	expiresAtUnix := jwtauth.ExpireIn(ttl)
	_, tokenString, err := auth.Encode(map[string]interface{}{
		"user_id": userId,
		"exp":     expiresAtUnix,
		"iat":     time.Now().Unix(),
	})
	if err != nil {
		return "", time.Time{}, err
	}
	expiresAt := time.Unix(expiresAtUnix, 0)
	return tokenString, expiresAt, nil
}
