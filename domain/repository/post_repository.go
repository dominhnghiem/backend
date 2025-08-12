package repository

import "context"

// Cổng (port) repo cho posts
type PostRepository interface {
	Create(ctx context.Context, authorID int64, body string) (int64, error)
	Get(ctx context.Context, id int64) (*PostView, error)
	UpdateBody(ctx context.Context, id, authorID int64, body string) error
	SoftDelete(ctx context.Context, id, authorID int64) error
	ListFeedByAuthorIDs(ctx context.Context, authorIDs []int64, limit, offset int) ([]*PostView, error)
}

// View đơn giản (có thể mở rộng)
type PostView struct {
	ID       int64
	AuthorID int64
	Body     string
}
