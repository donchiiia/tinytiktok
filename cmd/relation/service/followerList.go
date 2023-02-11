package service

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/relation"
	"tinytiktok/kitex_gen/user"
	"tinytiktok/pack"
)

type FollowerListService struct {
	ctx context.Context
}

func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{
		ctx: ctx,
	}
}

// FollowerList 获取 粉丝列表
func (s *FollowerListService) FollowerList(req *relation.DouyinRelationFollowerListRequest) ([]*user.User, error) {
	followers, err := db.MGetFollowerList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	followerSet, err := db.MGetFollowerSet(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.FollowerList(s.ctx, followers, followerSet)
}
