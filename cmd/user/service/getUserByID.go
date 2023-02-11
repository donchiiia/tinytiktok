package service

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/user"
	"tinytiktok/pack"
	"tinytiktok/pkg/errno"
)

type GetUserService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// GetUser 获取 用户信息
// 用户可能查询别的 user 信息，不一定是自己的。因此传入的 req.UserId可能是别的用户ID
func (s *GetUserService) GetUser(req *user.DouyinUserRequest, currID int64) (*user.User, error) {
	// 获取被访问用户的db对象
	users, err := db.GetUsersByID(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.UserNotExistErr
	}

	// 查询登录用户与被查询用户的关注信息
	aimUser := users[0]
	followInfo, err := db.GetFollowInfo(s.ctx, currID, aimUser.ID)
	if err != nil {
		return nil, err
	}

	// 返回结果用户信息
	isFollow := len(followInfo) == 0
	u := pack.UserInfo(aimUser, isFollow)
	return u, nil
}
