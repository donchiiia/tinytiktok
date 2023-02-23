package main

import (
	"context"
	"tinytiktok/cmd/favorite/pack"
	"tinytiktok/cmd/favorite/service"
	"tinytiktok/dal/db"
	favorite "tinytiktok/kitex_gen/favorite"
	"tinytiktok/pkg/errno"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.DouyinFavoriteActionRequest) (resp *favorite.DouyinFavoriteActionResponse, err error) {
	currID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(err)
		return resp, nil
	}

	if currID == 0 {
		resp = pack.BuildFavoriteActionResp(errno.UserNotExistErr)
		return resp, nil
	}

	if req.VideoId == 0 {
		resp = pack.BuildFavoriteActionResp(errno.VideoProcErr)
		return resp, nil
	}

	if req.ActionType != 1 && req.ActionType != 2 {
		resp = pack.BuildFavoriteActionResp(errno.ActionTypeErr)
		return resp, nil
	}

	err = service.NewFavoriteActionService(ctx).FavoriteAction(req, currID)
	if err != nil {
		resp = pack.BuildFavoriteActionResp(err)
		return resp, nil
	}
	resp = pack.BuildFavoriteActionResp(errno.Success)
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.DouyinFavoriteListRequest) (resp *favorite.DouyinFavoriteListResponse, err error) {
	currID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildFavoriteListResp(err)
		return resp, nil
	}

	// 获取followSet
	followSet, err := db.MGetFollowSet(ctx, currID)
	if err != nil {
		resp = pack.BuildFavoriteListResp(err)
		return resp, nil
	}

	// 获取favoriteSet
	favoriteSet, err := db.MGetFavoriteVideoSet(ctx, currID)
	if err != nil {
		resp = pack.BuildFavoriteListResp(err)
	}
	favoriteList, err := service.NewFavoriteListService(ctx).FavoriteList(req, currID, followSet, favoriteSet)
	if err != nil {
		resp = pack.BuildFavoriteListResp(err)
	}
	resp = pack.BuildFavoriteListResp(errno.Success)
	resp.VideoList = favoriteList
	return resp, nil
}
