package gormrepo

import (
	"context"

	"gorm.io/gorm"
	//"gorm.io/gorm/clause" // thêm import clause

	"your.module/name/repository/model"
)

type LikeRepository struct{ db *gorm.DB }

func NewLikeRepository(db *gorm.DB) *LikeRepository { return &LikeRepository{db: db} }

func (r *LikeRepository) Put(ctx context.Context, userID, postID int64) error {
	// Kiểm tra xem đã like chưa
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Like{}).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		// Đã like rồi
		return nil // hoặc trả về lỗi tùy yêu cầu
	}

	// Thực hiện like
	m := model.Like{UserID: userID, PostID: postID}
	return r.db.WithContext(ctx).Create(&m).Error
}

func (r *LikeRepository) Delete(ctx context.Context, userID, postID int64) error {
	result := r.db.WithContext(ctx).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Delete(&model.Like{})

	if result.Error != nil {
		return result.Error
	}

	// Nếu không có bản ghi nào bị xóa
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *LikeRepository) CountByPost(ctx context.Context, postID int64) (int64, error) {
	var n int64
	err := r.db.WithContext(ctx).
		Model(&model.Like{}).
		Where("post_id = ?", postID).
		Count(&n).Error
	return n, err
}
