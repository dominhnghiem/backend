package entity

import "time"

// Entity thuần cho bảng likes (PK composite)
type Like struct {
	UserID    int64
	PostID    int64
	CreatedAt time.Time
}
