package model

import "time"

type Follow struct {
	FollowerID int64     `gorm:"primaryKey"`
	FolloweeID int64     `gorm:"primaryKey;index"`
	CreatedAt  time.Time `gorm:"not null;default:now()"`
}

func (Follow) TableName() string { return "follows" }
