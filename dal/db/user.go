package db

import (
	"context"
	"gorm.io/gorm"
	"tinytiktok/pkg/consts"
)

type User struct {
	gorm.Model
	ID             int64   `gorm:"primarykey"`
	UserName       string  `gorm:"index:idx_username,unique;type:varchar(40);not null" json:"username"`
	Password       string  `gorm:"type:varchar(256);not null" json:"password"`
	FavoriteVideos []Video `gorm:"many2many:favorite" json:"favorite_videos"`
	FollowCount    int64   `gorm:"default:0" json:"follow_count"`
	FollowerCount  int64   `gorm:"default:0" json:"follower_count"`
}

func (u *User) TableName() string {
	return consts.UserTableName
}

// CreateUser 创建用户
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// GetUsersByName 通过 name 获取db.User
func GetUsersByName(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	err := DB.WithContext(ctx).Where("user_name = ?", userName).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetUsersByID 通过 id 获取db.User
func GetUsersByID(ctx context.Context, userID int64) ([]*User, error) {
	res := make([]*User, 0)
	err := DB.WithContext(ctx).Where("id = ?", userID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
