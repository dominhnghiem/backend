package repository

import "context"

type CommentRepository interface {
	Create(ctx context.Context, postID, authorID int64, body string) (int64, error)
	DeleteSoft(ctx context.Context, id, authorID int64) error
	ListByPost(ctx context.Context, postID int64, limit, offset int) ([]*CommentView, error)
}

type CommentView struct {
	ID       int64
	PostID   int64
	AuthorID int64
	Body     string
}
