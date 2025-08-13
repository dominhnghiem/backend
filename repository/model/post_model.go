package model

import (
	"time"

	"gorm.io/gorm"
)

// GORM model cho posts
type Post struct {
	ID        int64          `gorm:"primaryKey;autoIncrement"`
	AuthorID  int64          `gorm:"not null;index"`
	Body      string         `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Post) TableName() string { return "posts" }
