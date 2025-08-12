package social

import (
	"context"

	"your.module/name/domain/dto"
	"your.module/name/domain/repository"
)

type CreateCommentUsecase struct{ Comments repository.CommentRepository }

func NewCreateCommentUsecase(c repository.CommentRepository) *CreateCommentUsecase {
	return &CreateCommentUsecase{Comments: c}
}

func (uc *CreateCommentUsecase) Execute(ctx context.Context, authorID, postID int64, in dto.CreateCommentRequest) (*dto.CommentResponse, error) {
	id, err := uc.Comments.Create(ctx, postID, authorID, in.Body)
	if err != nil {
		return nil, err
	}
	return &dto.CommentResponse{ID: id, PostID: postID, AuthorID: authorID, Body: in.Body}, nil
}
