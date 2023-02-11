package pack

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/user"
)

// FollowList 将 db.Follow列表 包装成 user.User列表，返回关注列表
func FollowList(ctx context.Context, dbFollowList []*db.Follow) ([]*user.User, error) {
	followList := make([]*user.User, 0)
	for _, v := range dbFollowList {
		u, err := db.GetUsersByID(ctx, v.ToUserID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		followList = append(followList, UserInfo(u[0], true))
	}

	return followList, nil
}

// FollowerList 将 db.Follow列表 包装成 user.User列表，返回粉丝列表
func FollowerList(ctx context.Context, dbFollowerList []*db.Follow, dbFollowSet map[int64]struct{}) ([]*user.User, error) {
	followerList := make([]*user.User, 0)
	for _, v := range dbFollowerList {
		u, err := db.GetUsersByID(ctx, v.UserID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		_, isFollow := dbFollowSet[v.ToUserID]
		followerList = append(followerList, UserInfo(u[0], isFollow))
	}

	return followerList, nil
}
