package gormrepo

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
	"your.module/name/domain/repository"
	"your.module/name/repository/model"
)

type SessionRepository struct{ db *gorm.DB }

func NewSessionRepository(db *gorm.DB) *SessionRepository { return &SessionRepository{db: db} }

func (r *SessionRepository) Create(ctx context.Context, s *repository.Session) (int64, error) {
	m := model.Session{
		UserID:           s.UserID,
		RefreshTokenHash: s.RefreshTokenHash,
		UserAgent:        s.UserAgent,
		IP:               s.IP,
		ExpiresAt:        s.ExpiresAt,
		RevokedAt:        s.RevokedAt,
	}
	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return 0, err
	}
	return m.ID, nil
}

func (r *SessionRepository) FindActiveByHash(ctx context.Context, hash string) (*repository.Session, error) {
	var m model.Session
	err := r.db.WithContext(ctx).
		Where("refresh_token_hash = ? AND revoked_at IS NULL AND expires_at > now()", hash).
		First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &repository.Session{
		ID:               m.ID,
		UserID:           m.UserID,
		RefreshTokenHash: m.RefreshTokenHash,
		UserAgent:        m.UserAgent,
		IP:               m.IP,
		ExpiresAt:        m.ExpiresAt,
		RevokedAt:        m.RevokedAt,
		CreatedAt:        m.CreatedAt,
	}, nil
}

func (r *SessionRepository) RevokeByID(ctx context.Context, id int64) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&model.Session{}).
		Where("id = ? AND revoked_at IS NULL", id).
		Update("revoked_at", &now).Error
}

func (r *SessionRepository) RevokeByHash(ctx context.Context, hash string) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&model.Session{}).
		Where("refresh_token_hash = ? AND revoked_at IS NULL", hash).
		Update("revoked_at", &now).Error
}
