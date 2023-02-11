package service

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/favorite"
	"tinytiktok/pkg/errno"
)

type FavoriteActionService struct {
	ctx context.Context
}

// NewFavoriteActionService new FavoriteActionService
func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{ctx: ctx}
}

// FavoriteAction 点赞操作 点赞/取消点赞
func (s *FavoriteActionService) FavoriteAction(req *favorite.DouyinFavoriteActionRequest, currID int64) error {
	// 1-点赞
	if req.ActionType == 1 {
		return db.AddFavorite(s.ctx, currID, req.VideoId)
	}
	// 2-取消点赞
	if req.ActionType == 2 {
		return db.UnFavorite(s.ctx, currID, req.VideoId)
	}
	return errno.ActionTypeErr
}
