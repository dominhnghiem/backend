package repository

import "context"

type FollowRepository interface {
	Put(ctx context.Context, followerID, followeeID int64) error
	Delete(ctx context.Context, followerID, followeeID int64) error
	ListFolloweeIDs(ctx context.Context, followerID int64, limit, offset int) ([]int64, error)
}
