package db

import (
	"context"
	"gorm.io/gorm"
	"tinytiktok/pkg/consts"
	"tinytiktok/pkg/errno"
)

type Follow struct {
	gorm.Model
	User     User  `gorm:"foreignkey:UserID;"`
	UserID   int64 `gorm:"index:idx_userid,unique;not null"`
	ToUser   User  `gorm:"foreignkey:ToUserID;"`
	ToUserID int64 `gorm:"index:idx_userid,unique;index:idx_userid_to;not null"`
}

func (Follow) TableName() string {
	return consts.FollowTableName
}

// NewFollow 新建 关注 信息
func NewFollow(ctx context.Context, userID int64, toUserID int64) error {
	// 事务操作，新建关系，修改user表的follow_count和following_count字段
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 新建关系
		follow := Follow{
			UserID:   userID,
			ToUserID: toUserID,
		}
		err := tx.Create(follow).Error
		if err != nil {
			return err
		}

		// 修改user表
		// 更新 follow_count 字段
		res := tx.Model(&User{}).Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errno.DBErr
		}

		// 更新 follower_count 字段
		res = tx.Model(&User{}).Where("id = ?", toUserID).Update("follower_count", gorm.Expr("follower_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errno.DBErr
		}
		return nil
	})
	return err
}

// UnFollow 取消关注
func UnFollow(ctx context.Context, userID int64, toUserID int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 取消关注，事务操作
		// 先判断是否存在记录
		follow := new(Follow)
		err := tx.Where("user_id = ? And to_user_id = ?", userID, toUserID).Find(&follow).Error
		if err != nil {
			return err
		}

		// 关注信息可以不用软删
		err = tx.Unscoped().Delete(&follow).Error
		if err != nil {
			return err
		}
		// 修改follow_count
		model := tx.Model(&User{})
		res := model.Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errno.DBErr
		}
		res = model.Where("id = ?", toUserID).Update("follower_count", gorm.Expr("follower_count - 1", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errno.DBErr
		}
		return nil
	})
	return err
}

// GetFollowInfo 获取 关注 信息
func GetFollowInfo(ctx context.Context, userID int64, toUserID int64) ([]*Follow, error) {
	follow := make([]*Follow, 0)
	err := DB.WithContext(ctx).Find(&follow, "user_id = ? AND to_user_id = ?", userID, toUserID).Error
	if err != nil {
		return nil, err
	}
	return follow, nil
}

// MGetFollowList 获取 follow 列表
func MGetFollowList(ctx context.Context, userID int64) ([]*Follow, error) {
	followList := make([]*Follow, 0)
	err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&followList).Error
	if err != nil {
		return nil, err
	}
	return followList, nil
}

// MGetFollowSet 获取 follow 集合
func MGetFollowSet(ctx context.Context, userID int64) (map[int64]struct{}, error) {
	followSet := make(map[int64]struct{}, 0)
	followList, err := MGetFollowList(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, v := range followList {
		followSet[v.ToUserID] = struct{}{}
	}
	return followSet, nil
}

// MGetFollowerList 获取 follower 列表
func MGetFollowerList(ctx context.Context, userID int64) ([]*Follow, error) {
	followerList := make([]*Follow, 0)
	err := DB.WithContext(ctx).Where("to_user_id = ?", userID).Find(&followerList).Error
	if err != nil {
		return nil, err
	}
	return followerList, nil
}

// MGetFollowerSet 获取 follower 集合
func MGetFollowerSet(ctx context.Context, toUserID int64) (map[int64]struct{}, error) {
	followerSet := make(map[int64]struct{}, 0)
	followerList, err := MGetFollowerList(ctx, toUserID)
	if err != nil {
		return nil, err
	}

	for _, v := range followerList {
		followerSet[v.UserID] = struct{}{}
	}
	return followerSet, nil
}
