package model

import "time"

type Session struct {
	ID               int64  `gorm:"primaryKey;autoIncrement"`
	UserID           int64  `gorm:"not null;index"`
	RefreshTokenHash string `gorm:"uniqueIndex;not null"`
	UserAgent        *string
	IP               *string
	ExpiresAt        time.Time `gorm:"not null"`
	RevokedAt        *time.Time
	CreatedAt        time.Time `gorm:"not null;default:now()"`
}

func (Session) TableName() string { return "sessions" }
