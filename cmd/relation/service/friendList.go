package service

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/relation"
	"tinytiktok/kitex_gen/user"
	"tinytiktok/pack"
)

type FriendListService struct {
	ctx context.Context
}

func NewFriendListService(ctx context.Context) *FriendListService {
	return &FriendListService{
		ctx: ctx,
	}
}

// FriendList 获取 好友列表
func (s *FriendListService) FriendList(req *relation.DouyinRelationFriendListRequest) ([]*user.User, error) {
	friends, err := db.MGetFriendList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return pack.FriendList(s.ctx, friends)
}
