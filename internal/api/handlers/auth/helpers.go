package auth

import "github.com/carterjackson/ranked-pick-api/internal/resources"

type AuthResponse struct {
	Token string          `json:"token"`
	User  *resources.User `json:"user"`
}
