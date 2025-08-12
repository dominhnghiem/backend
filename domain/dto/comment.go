package dto

type CreateCommentRequest struct {
	Body string `json:"body" validate:"required,min=1,max=1000"`
}
type CommentResponse struct {
	ID       int64  `json:"id"`
	PostID   int64  `json:"post_id"`
	AuthorID int64  `json:"author_id"`
	Body     string `json:"body"`
}
