package jwt

import (
	"context"
	"encoding/json"
	"time"

	"github.com/carterjackson/ranked-pick-api/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type Claims struct {
	UserId  int64     `json:"user_id"`
	Expires time.Time `json:"expires"`
}

func AddMiddleware(router *chi.Mux) {
	// Seek, verify and validate JWT tokensx
	router.Use(jwtauth.Verifier(config.Config.Auth))

	// Handle valid / invalid tokens
	router.Use(jwtauth.Authenticator(config.Config.Auth))
}

func NewToken(userId int64) (string, error) {
	claims := &Claims{UserId: userId}
	_, tokenString, err := config.Config.Auth.Encode(map[string]interface{}{
		"user_id": claims.UserId,
	})
	if err != nil {
		return "", err
	}

	return tokenString, nil
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
