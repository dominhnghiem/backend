package auth

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"your.module/name/domain/dto"
	"your.module/name/domain/entity"
	"your.module/name/domain/repository"
)

type SignupUsecase struct {
	Users repository.UserRepository
}

func NewSignupUsecase(users repository.UserRepository) *SignupUsecase {
	return &SignupUsecase{Users: users}
}

var ErrEmailExists = errors.New("email already registered")

func (uc *SignupUsecase) Execute(ctx context.Context, in dto.SignupRequest) (*dto.SignupResponse, error) {
	// email đã tồn tại?
	exist, err := uc.Users.FindByEmail(ctx, in.Email)
	if err != nil {
		return nil, err
	}
	if exist != nil {
		return nil, ErrEmailExists
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &entity.User{
		Email:        in.Email,
		PasswordHash: string(hash),
	}
	if in.Name != "" {
		u.Name = &in.Name
	}

	id, err := uc.Users.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return &dto.SignupResponse{
		ID:    id,
		Email: in.Email,
		Name:  in.Name,
	}, nil
}
