package service

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/relation"
	"tinytiktok/kitex_gen/user"
	"tinytiktok/pack"
)

type FollowListService struct {
	ctx context.Context
}

func NewFollowListService(ctx context.Context) *FollowListService {
	return &FollowListService{
		ctx: ctx,
	}
}

// FollowList 获取 关注者列表
func (s *FollowListService) FollowList(req *relation.DouyinRelationFollowListRequest) ([]*user.User, error) {
	follows, err := db.MGetFollowList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return pack.FollowList(s.ctx, follows)
}
