package gormrepo

import (
	"context"

	"gorm.io/gorm"
	"your.module/name/domain/repository"
	"your.module/name/repository/model"
)

type PostRepository struct{ db *gorm.DB }

func NewPostRepository(db *gorm.DB) *PostRepository { return &PostRepository{db: db} }

func (r *PostRepository) Create(ctx context.Context, authorID int64, body string) (int64, error) {
	m := model.Post{AuthorID: authorID, Body: body}
	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return 0, err
	}
	return m.ID, nil
}

func (r *PostRepository) Get(ctx context.Context, id int64) (*repository.PostView, error) {
	var m model.Post
	if err := r.db.WithContext(ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	return &repository.PostView{ID: m.ID, AuthorID: m.AuthorID, Body: m.Body}, nil
}

func (r *PostRepository) UpdateBody(ctx context.Context, id, authorID int64, body string) error {
	return r.db.WithContext(ctx).
		Model(&model.Post{}).
		Where("id = ? AND author_id = ? AND deleted_at IS NULL", id, authorID).
		Update("body", body).Error
}

func (r *PostRepository) SoftDelete(ctx context.Context, id, authorID int64) error {
	return r.db.WithContext(ctx).
		Where("id = ? AND author_id = ? AND deleted_at IS NULL", id, authorID).
		Delete(&model.Post{}).Error
}

func (r *PostRepository) ListFeedByAuthorIDs(ctx context.Context, authorIDs []int64, limit, offset int) ([]*repository.PostView, error) {
	var rows []model.Post
	if err := r.db.WithContext(ctx).
		Where("author_id IN ? AND deleted_at IS NULL", authorIDs).
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*repository.PostView, 0, len(rows))
	for _, m := range rows {
		out = append(out, &repository.PostView{ID: m.ID, AuthorID: m.AuthorID, Body: m.Body})
	}
	return out, nil
}
