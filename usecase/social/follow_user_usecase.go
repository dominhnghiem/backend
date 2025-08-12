package social

import "context"
import "your.module/name/domain/repository"

type FollowUserUsecase struct{ Follows repository.FollowRepository }

func NewFollowUserUsecase(f repository.FollowRepository) *FollowUserUsecase {
	return &FollowUserUsecase{Follows: f}
}

func (uc *FollowUserUsecase) Follow(ctx context.Context, me, other int64) error {
	return uc.Follows.Put(ctx, me, other)
}
func (uc *FollowUserUsecase) Unfollow(ctx context.Context, me, other int64) error {
	return uc.Follows.Delete(ctx, me, other)
}
