package social

import "context"
import "your.module/name/domain/repository"

type UpdatePostUsecase struct{ Posts repository.PostRepository }

func NewUpdatePostUsecase(p repository.PostRepository) *UpdatePostUsecase {
	return &UpdatePostUsecase{Posts: p}
}
func (uc *UpdatePostUsecase) Execute(ctx context.Context, id, authorID int64, body string) error {
	return uc.Posts.UpdateBody(ctx, id, authorID, body)
}
