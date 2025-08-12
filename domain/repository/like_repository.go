package repository

import "context"

type LikeRepository interface {
	Put(ctx context.Context, userID, postID int64) error    // upsert (like)
	Delete(ctx context.Context, userID, postID int64) error // unlike
	CountByPost(ctx context.Context, postID int64) (int64, error)
}
