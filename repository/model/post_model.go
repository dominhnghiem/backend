package model

import "time"

// GORM model cho posts
type Post struct {
	ID        int64      `gorm:"primaryKey;autoIncrement"`
	AuthorID  int64      `gorm:"not null;index"`
	Body      string     `gorm:"not null"`
	CreatedAt time.Time  `gorm:"not null;default:now()"`
	UpdatedAt time.Time  `gorm:"not null;default:now()"`
	DeletedAt *time.Time `gorm:"index"`
}

func (Post) TableName() string { return "posts" }
