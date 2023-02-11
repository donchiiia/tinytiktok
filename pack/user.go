package pack

import (
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/user"
)

// UserInfo 将 db.User 包装成 user.User
func UserInfo(dbUser *db.User, isFollow bool) *user.User {
	return &user.User{
		Id:            dbUser.ID,
		Name:          dbUser.UserName,
		FollowCount:   &dbUser.FollowCount,
		FollowerCount: &dbUser.FollowerCount,
		IsFollow:      isFollow,
	}
}
