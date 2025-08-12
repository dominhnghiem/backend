package model

import "time"

type User struct {
	ID              int64  `gorm:"primaryKey;autoIncrement"`
	Email           string `gorm:"uniqueIndex;not null"`
	PasswordHash    string `gorm:"not null"`
	Name            *string
	EmailVerifiedAt *time.Time
	CreatedAt       time.Time  `gorm:"not null;default:now()"`
	UpdatedAt       time.Time  `gorm:"not null;default:now()"`
	DeletedAt       *time.Time `gorm:"index"`
}

func (User) TableName() string { return "users" }
