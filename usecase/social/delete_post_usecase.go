package social

import "context"
import "your.module/name/domain/repository"

type DeletePostUsecase struct{ Posts repository.PostRepository }

func NewDeletePostUsecase(p repository.PostRepository) *DeletePostUsecase {
	return &DeletePostUsecase{Posts: p}
}
func (uc *DeletePostUsecase) Execute(ctx context.Context, id, authorID int64) error {
	return uc.Posts.SoftDelete(ctx, id, authorID)
}
