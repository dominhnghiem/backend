package user

import (
	"context"
	"errors"

	"your.module/name/domain/entity"
	"your.module/name/domain/repository"
)

var ErrUserNotFound = errors.New("user not found")

type ProfileUsecase struct {
	Users repository.UserRepository
}

func NewProfileUsecase(users repository.UserRepository) *ProfileUsecase {
	return &ProfileUsecase{Users: users}
}

func (uc *ProfileUsecase) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	u, err := uc.Users.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, ErrUserNotFound
	}
	return u, nil
}
