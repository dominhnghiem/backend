package social

import (
	"context"

	"your.module/name/domain/dto"
	"your.module/name/domain/repository"
)

type ListFeedUsecase struct {
	Posts   repository.PostRepository
	Follows repository.FollowRepository
}

func NewListFeedUsecase(p repository.PostRepository, f repository.FollowRepository) *ListFeedUsecase {
	return &ListFeedUsecase{Posts: p, Follows: f}
}

func (uc *ListFeedUsecase) Execute(ctx context.Context, me int64, limit, offset int) ([]*dto.PostResponse, error) {
	ids, err := uc.Follows.ListFolloweeIDs(ctx, me, 1000, 0)
	if err != nil {
		return nil, err
	}
	ids = append(ids, me) // bao gồm bài của mình
	rows, err := uc.Posts.ListFeedByAuthorIDs(ctx, ids, limit, offset)
	if err != nil {
		return nil, err
	}
	out := make([]*dto.PostResponse, 0, len(rows))
	for _, r := range rows {
		out = append(out, &dto.PostResponse{ID: r.ID, AuthorID: r.AuthorID, Body: r.Body})
	}
	return out, nil
}
