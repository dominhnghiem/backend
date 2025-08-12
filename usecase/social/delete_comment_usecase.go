package social

import "context"
import "your.module/name/domain/repository"

type DeleteCommentUsecase struct{ Comments repository.CommentRepository }

func NewDeleteCommentUsecase(c repository.CommentRepository) *DeleteCommentUsecase {
	return &DeleteCommentUsecase{Comments: c}
}

func (uc *DeleteCommentUsecase) Execute(ctx context.Context, id, authorID int64) error {
	return uc.Comments.DeleteSoft(ctx, id, authorID)
}
