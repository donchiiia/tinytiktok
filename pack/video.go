package pack

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/feed"
)

// VideoInfo 将单独一个 db.video 包装成 feed.video
func VideoInfo(ctx context.Context, dbVideo *db.Video, currID int64, dbFollowSet map[int64]struct{}, dbFavoriteSet map[int64]struct{}) (*feed.Video, error) {
	if dbVideo == nil {
		return nil, nil
	}

	// 查询视频作者
	user, err := db.GetUsersByID(ctx, dbVideo.AuthorID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	var isFollow = false
	if currID != -1 {
		_, isFollow = dbFollowSet[dbVideo.AuthorID]
	}

	var isFavorite = false
	if currID != -1 {
		_, isFavorite = dbFavoriteSet[dbVideo.ID]
	}

	author := UserInfo(user[0], isFollow)
	favoriteCount := dbVideo.FavoriteCount
	commentCount := dbVideo.CommentCount

	return &feed.Video{
		Id:            dbVideo.ID,
		Author:        author,
		PlayUrl:       dbVideo.PlayUrl,
		CoverUrl:      dbVideo.CoverUrl,
		FavoriteCount: favoriteCount,
		CommentCount:  commentCount,
		IsFavorite:    isFavorite,
		Title:         dbVideo.Title,
	}, nil
}

// VideoListInfo 将 db.video列表 包装成 feed.video列表
func VideoListInfo(ctx context.Context, dbVideoList []*db.Video, currID int64, dbFollowSet map[int64]struct{}, dbFavoriteSet map[int64]struct{}) ([]*feed.Video, error) {
	videoList := make([]*feed.Video, 0)
	for _, v := range dbVideoList {
		video, err := VideoInfo(ctx, v, currID, dbFollowSet, dbFavoriteSet)
		if err != nil {
			return nil, err
		}

		videoList = append(videoList, video)
	}
	return videoList, nil
}
