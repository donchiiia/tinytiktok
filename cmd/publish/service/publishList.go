package service

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/feed"
	"tinytiktok/kitex_gen/publish"
	"tinytiktok/pack"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new PublishListService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

// PublishList 获取当前登录用户的发布视频列表
func (s *PublishListService) PublishList(req *publish.DouyinPublishListRequest, dbFollowSet map[int64]struct{}, dbFavoriteSet map[int64]struct{}) ([]*feed.Video, error) {
	videos, err := db.MGetPublishVideoList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	publishList, err := pack.VideoListInfo(s.ctx, videos, req.UserId, dbFollowSet, dbFavoriteSet)
	if err != nil {
		return nil, err
	}

	return publishList, nil
}
