package model

import "time"

type Comment struct {
	ID        int64      `gorm:"primaryKey;autoIncrement"`
	PostID    int64      `gorm:"not null;index"`
	AuthorID  int64      `gorm:"not null;index"`
	Body      string     `gorm:"not null"`
	CreatedAt time.Time  `gorm:"not null;default:now()"`
	UpdatedAt time.Time  `gorm:"not null;default:now()"`
	DeletedAt *time.Time `gorm:"index"`
}

func (Comment) TableName() string { return "comments" }
