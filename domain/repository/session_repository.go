package repository

import (
	"context"
	"time"
)

type Session struct {
	ID               int64
	UserID           int64
	RefreshTokenHash string
	UserAgent        *string
	IP               *string
	ExpiresAt        time.Time
	RevokedAt        *time.Time
	CreatedAt        time.Time
}

type SessionRepository interface {
	Create(ctx context.Context, s *Session) (int64, error)
	FindActiveByHash(ctx context.Context, hash string) (*Session, error)
	RevokeByID(ctx context.Context, id int64) error
	RevokeByHash(ctx context.Context, hash string) error
}
