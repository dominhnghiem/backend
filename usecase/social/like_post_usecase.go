package social

import "context"
import "your.module/name/domain/repository"

type LikePostUsecase struct{ Likes repository.LikeRepository }

func NewLikePostUsecase(l repository.LikeRepository) *LikePostUsecase {
	return &LikePostUsecase{Likes: l}
}

func (uc *LikePostUsecase) Like(ctx context.Context, userID, postID int64) error {
	return uc.Likes.Put(ctx, userID, postID)
}
func (uc *LikePostUsecase) Unlike(ctx context.Context, userID, postID int64) error {
	return uc.Likes.Delete(ctx, userID, postID)
}
