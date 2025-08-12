package dto

type LikeRequest struct {
	PostID int64 `json:"post_id" validate:"required"`
}
type FollowRequest struct {
	UserID int64 `json:"user_id" validate:"required"`
}
type SimpleAck struct {
	OK bool `json:"ok"`
}
