package entity

import "time"

type User struct {
	ID              int64      `json:"id"`
	Email           string     `json:"email"`
	PasswordHash    string     `json:"-"`
	Name            *string    `json:"name,omitempty"`
	EmailVerifiedAt *time.Time `json:"email_verified_at,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
