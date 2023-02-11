package main

import (
	"context"
	"tinytiktok/cmd/feed/pack"
	"tinytiktok/cmd/feed/service"
	"tinytiktok/dal/db"
	feed "tinytiktok/kitex_gen/feed"
	"tinytiktok/pkg/errno"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetUserFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	// feed 可能传入空的token
	// 解析当前登录用户ID
	var currID int64
	if req.Token != nil && *req.Token != "" {
		currID, err = Jwt.GetUserIDFromToken(*req.Token)
		if err != nil {
			resp = pack.BuildFeedResp(err)
			return resp, nil
		}
	}

	// 获取followSet
	followSet, err := db.MGetFollowSet(ctx, currID)
	if err != nil {
		resp = pack.BuildFeedResp(err)
		return resp, nil
	}

	// 获取favoriteSet
	favoriteSet, err := db.MGetFavoriteVideoSet(ctx, currID)
	if err != nil {
		resp = pack.BuildFeedResp(err)
		return resp, nil
	}
	videoList, nextTime, err := service.NewGetFeedService(ctx).GetUserFeed(req, currID, followSet, favoriteSet)
	resp = pack.BuildFeedResp(errno.Success)
	resp.VideoList = videoList
	resp.NextTime = &nextTime
	return resp, nil
}
