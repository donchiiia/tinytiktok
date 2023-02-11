package pack

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/feed"
)

// FavoriteVideoListInfo 将 db.video列表 包装成 feed.video列表，返回 用户点赞视频列表
func FavoriteVideoListInfo(ctx context.Context, dbVideoList []*db.Video, currID int64, dbFollowSet map[int64]struct{}, dbFavoriteSet map[int64]struct{}) ([]*feed.Video, error) {
	videoList := make([]*feed.Video, 0)
	var err error

	videoList, err = VideoListInfo(ctx, dbVideoList, currID, dbFollowSet, dbFavoriteSet)
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
