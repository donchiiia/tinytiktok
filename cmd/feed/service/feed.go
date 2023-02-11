package service

import (
	"context"
	"time"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/feed"
	"tinytiktok/pack"
)

// 固定每次返回video数量最大为30
const LIMIT = 30

type GetFeedService struct {
	ctx context.Context
}

// NewGetFeedService new GetFeedService
func NewGetFeedService(ctx context.Context) *GetFeedService {
	return &GetFeedService{
		ctx: ctx,
	}
}

// GetUserFeed 获取用户的视频流
func (s *GetFeedService) GetUserFeed(req *feed.DouyinFeedRequest, currID int64, dbFollowSet map[int64]struct{}, dbFavoriteSet map[int64]struct{}) (videoList []*feed.Video, nextTime int64, err error) {
	// 获取当前登录用户ID

	videos, err := db.MGetFeed(s.ctx, req.LatestTime, LIMIT)
	if err != nil {
		return nil, 0, err
	}

	// 如果返回空，说明数据库中没有视频，但属于正常流程，nextTime依旧要正确更新，此处更新为当前时间
	if len(videos) == 0 {
		nextTime = time.Now().UnixMilli()
		return nil, nextTime, nil
	}

	videoList, err = pack.VideoListInfo(s.ctx, videos, currID, dbFollowSet, dbFavoriteSet)
	if err != nil {
		// 此时如果发生 error，说明以及查到了db.Video的数据，只是在转换过程中发生了异常，也需要将nextTime更新为当前时间
		nextTime = time.Now().UnixMilli()
		return nil, nextTime, err
	}
	// 找到数据就更新为发布最早的视频时间
	nextTime = videos[len(videos)-1].UpdatedAt.UnixMilli()
	return videoList, nextTime, nil
}
