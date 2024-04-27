package jwt

import (
	"context"
	"encoding/json"

	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type Claims struct {
	UserId int64 `json:"user_id"`
}

func AddMiddleware(router *chi.Mux) {
	// Seek, verify and validate JWT tokensx
	router.Use(jwtauth.Verifier(config.Config.Auth))

	// Handle valid / invalid tokens
	router.Use(jwtauth.Authenticator(config.Config.Auth))
}

func ParseClaims(ctx context.Context) (*Claims, error) {
	_, claimsMap, _ := jwtauth.FromContext(ctx)

	var claims *Claims
	encodedClaims, err := json.Marshal(claimsMap)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(encodedClaims, claims)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
