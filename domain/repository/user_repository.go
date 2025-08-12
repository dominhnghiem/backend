package repository

import (
	"context"

	"your.module/name/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, u *entity.User) (int64, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByID(ctx context.Context, id int64) (*entity.User, error)
}
