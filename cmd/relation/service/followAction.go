package service

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/relation"
	"tinytiktok/pkg/errno"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{
		ctx: ctx,
	}
}

// FollowAction 关注操作 关注/取消关注
func (s *RelationActionService) FollowAction(req *relation.DouyinRelationActionRequest) error {
	// 1. 关注操作
	if req.ActionType == 1 {
		return db.NewFollow(s.ctx, req.UserId, req.ToUserId)
	}

	// 2. 取关操作
	if req.ActionType == 2 {
		return db.UnFollow(s.ctx, req.UserId, req.ToUserId)
	}

	return errno.ActionTypeErr
}
