package db

import (
	"context"
	"gorm.io/gorm"
	"tinytiktok/pkg/consts"
)

type Friend struct {
	gorm.Model
	User     User  `gorm:"foreignkey:UserID;"`
	UserID   int64 `gorm:"index:idx_userid,unique;not null"`
	ToUser   User  `gorm:"foreignkey:ToUserID;"`
	ToUserID int64 `gorm:"index:idx_userid,unique;index:idx_userid_to;not null"`
}

func (Friend) TableName() string {
	return consts.FriendTableName
}

// NewFriend 新建好友信息
func NewFriend(ctx context.Context, userID int64, toUserID int64) error {
	// 事务操作，新建朋友关系
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 新建关系
		friend := Friend{
			UserID:   userID,
			ToUserID: toUserID,
		}
		err := tx.Create(friend).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// UnFriend 取消好友关系
func UnFriend(ctx context.Context, userID int64, toUserID int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 取消朋友关系，事务操作
		// 先判断是否存在记录
		friend := new(Friend)
		err := tx.Where("user_id = ? And to_user_id = ?", userID, toUserID).Find(&friend).Error
		if err != nil {
			return err
		}

		// 朋友信息可以不用软删
		err = tx.Unscoped().Delete(&friend).Error
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetFriendInfo 获取 friend 信息
func GetFriendInfo(ctx context.Context, userID int64, toUserID int64) ([]*Friend, error) {
	friend := make([]*Friend, 0)
	err := DB.WithContext(ctx).Find(&friend, "user_id = ? AND to_user_id = ?", userID, toUserID).Error
	if err != nil {
		return nil, err
	}
	return friend, nil
}

// MGetFriendList (db)获取好友列表
func MGetFriendList(ctx context.Context, userID int64) ([]*Friend, error) {
	friendList := make([]*Friend, 0)
	err := DB.WithContext(ctx).Model(&Friend{}).Where(&Friend{UserID: userID}).Find(&friendList).Error
	if err != nil {
		return nil, err
	}
	return friendList, nil
}
