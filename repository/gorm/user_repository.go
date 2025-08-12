package gormrepo

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"your.module/name/domain/entity"
	"your.module/name/repository/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository { return &UserRepository{db: db} }

func (r *UserRepository) Create(ctx context.Context, u *entity.User) (int64, error) {
	m := model.User{
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
	}
	if u.Name != nil {
		m.Name = u.Name
	}
	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return 0, err
	}
	return m.ID, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var m model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:              m.ID,
		Email:           m.Email,
		PasswordHash:    m.PasswordHash,
		Name:            m.Name,
		EmailVerifiedAt: m.EmailVerifiedAt,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	var m model.User
	err := r.db.WithContext(ctx).First(&m, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:              m.ID,
		Email:           m.Email,
		PasswordHash:    m.PasswordHash,
		Name:            m.Name,
		EmailVerifiedAt: m.EmailVerifiedAt,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}, nil
}
