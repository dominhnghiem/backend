package gormrepo

import (
	"context"

	"gorm.io/gorm"
	"your.module/name/domain/repository"
	"your.module/name/repository/model"
)

type CommentRepository struct{ db *gorm.DB }

func NewCommentRepository(db *gorm.DB) *CommentRepository { return &CommentRepository{db: db} }

func (r *CommentRepository) Create(ctx context.Context, postID, authorID int64, body string) (int64, error) {
	m := model.Comment{PostID: postID, AuthorID: authorID, Body: body}
	if err := r.db.WithContext(ctx).Create(&m).Error; err != nil {
		return 0, err
	}
	return m.ID, nil
}

func (r *CommentRepository) Delete(ctx context.Context, id, authorID int64) error {
	return r.db.WithContext(ctx).
		Where("id = ? AND author_id = ? AND deleted_at IS NULL", id, authorID).
		Delete(&model.Comment{}).Error
}

func (r *CommentRepository) ListByPost(ctx context.Context, postID int64, limit, offset int) ([]*repository.CommentView, error) {
	var rows []model.Comment
	if err := r.db.WithContext(ctx).
		Where("post_id = ? AND deleted_at IS NULL", postID).
		Order("created_at ASC").
		Limit(limit).Offset(offset).
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*repository.CommentView, 0, len(rows))
	for _, m := range rows {
		out = append(out, &repository.CommentView{ID: m.ID, PostID: m.PostID, AuthorID: m.AuthorID, Body: m.Body})
	}
	return out, nil
}
