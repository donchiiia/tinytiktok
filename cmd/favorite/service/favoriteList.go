package service

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/favorite"
	"tinytiktok/kitex_gen/feed"
	"tinytiktok/pack"
)

type FavoriteListService struct {
	ctx context.Context
}

// NewFavoriteListService creates a new FavoriteListService
func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}

// FavoriteList 返回 currID对应用户的点赞视频列表
func (s *FavoriteListService) FavoriteList(req *favorite.DouyinFavoriteListRequest, currID int64, dbFollowSet map[int64]struct{}, dbFavoriteSet map[int64]struct{}) ([]*feed.Video, error) {
	favoriteVideoList, err := db.MGetFavoriteVideoList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	videos, err := pack.FavoriteVideoListInfo(s.ctx, favoriteVideoList, currID, dbFollowSet, dbFavoriteSet)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
