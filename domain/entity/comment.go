package entity

import "time"

// Entity thuần cho bảng comments
type Comment struct {
	ID        int64
	PostID    int64
	AuthorID  int64
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
