package dto

// DTO tạo/cập nhật bài viết
type CreatePostRequest struct {
	Body string `json:"body" validate:"required,min=1,max=2000"`
}
type UpdatePostRequest struct {
	Body string `json:"body" validate:"required,min=1,max=2000"`
}
type PostResponse struct {
	ID       int64  `json:"id"`
	AuthorID int64  `json:"author_id"`
	Body     string `json:"body"`
}
