package gormrepo

import (
	"context"

	"gorm.io/gorm"
	"your.module/name/repository/model"
)

type LikeRepository struct{ db *gorm.DB }

func NewLikeRepository(db *gorm.DB) *LikeRepository { return &LikeRepository{db: db} }

func (r *LikeRepository) Put(ctx context.Context, userID, postID int64) error {
	m := model.Like{UserID: userID, PostID: postID}
	return r.db.WithContext(ctx).Clauses(
		// upsert on conflict composite PK
		gorm.Clauses{OnConflict: gorm.OnConflict{DoNothing: true}},
	).Create(&m).Error
}

func (r *LikeRepository) Delete(ctx context.Context, userID, postID int64) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Delete(&model.Like{}).Error
}

func (r *LikeRepository) CountByPost(ctx context.Context, postID int64) (int64, error) {
	var n int64
	err := r.db.WithContext(ctx).Model(&model.Like{}).
		Where("post_id = ?", postID).
		Count(&n).Error
	return n, err
}
