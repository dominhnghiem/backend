package model

import "time"

type Like struct {
	UserID    int64     `gorm:"primaryKey"`
	PostID    int64     `gorm:"primaryKey;index"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
}

func (Like) TableName() string { return "likes" }
