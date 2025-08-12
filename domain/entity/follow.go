package entity

import "time"

// Entity thuần cho bảng follows (PK composite)
type Follow struct {
	FollowerID int64
	FolloweeID int64
	CreatedAt  time.Time
}
