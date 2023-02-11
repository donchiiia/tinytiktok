package pack

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/user"
)

// FriendList 将 db.Friend列表 包装成 user.User列表
func FriendList(ctx context.Context, dbFriendList []*db.Friend) (friendList []*user.User, err error) {
	friendList = make([]*user.User, 0)
	for _, v := range dbFriendList {
		friend, err := db.GetUsersByID(ctx, v.ToUserID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		friendList = append(friendList, UserInfo(friend[0], true))
	}
	return friendList, nil
}
