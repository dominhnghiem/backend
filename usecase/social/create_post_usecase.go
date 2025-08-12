package social

import (
	"context"

	"your.module/name/domain/dto"
	"your.module/name/domain/repository"
)

type CreatePostUsecase struct{ Posts repository.PostRepository }

func NewCreatePostUsecase(p repository.PostRepository) *CreatePostUsecase {
	return &CreatePostUsecase{Posts: p}
}

func (uc *CreatePostUsecase) Execute(ctx context.Context, authorID int64, in dto.CreatePostRequest) (*dto.PostResponse, error) {
	id, err := uc.Posts.Create(ctx, authorID, in.Body)
	if err != nil {
		return nil, err
	}
	return &dto.PostResponse{ID: id, AuthorID: authorID, Body: in.Body}, nil
}
