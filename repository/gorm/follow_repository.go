package gormrepo

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause" // thêm dòng này
	"your.module/name/repository/model"
)

type FollowRepository struct{ db *gorm.DB }

func NewFollowRepository(db *gorm.DB) *FollowRepository {
	return &FollowRepository{db: db}
}

func (r *FollowRepository) Put(ctx context.Context, followerID, followeeID int64) error {
	m := model.Follow{FollowerID: followerID, FolloweeID: followeeID}
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "follower_id"}, {Name: "followee_id"}},
			DoNothing: true,
		}).
		Create(&m).Error
}

func (r *FollowRepository) Delete(ctx context.Context, followerID, followeeID int64) error {
	return r.db.WithContext(ctx).
		Where("follower_id = ? AND followee_id = ?", followerID, followeeID).
		Delete(&model.Follow{}).Error
}

func (r *FollowRepository) ListFolloweeIDs(ctx context.Context, followerID int64, limit, offset int) ([]int64, error) {
	var rows []model.Follow
	if err := r.db.WithContext(ctx).
		Where("follower_id = ?", followerID).
		Limit(limit).Offset(offset).
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]int64, 0, len(rows))
	for _, m := range rows {
		out = append(out, m.FolloweeID)
	}
	return out, nil
}
