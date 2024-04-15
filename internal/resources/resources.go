package resources

import "time"

type User struct {
	Id           int       `json:"id"`
	PasswordHash string    `json:"password_hash"`
	DisplayName  string    `json:"display_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"created_at"`
}
