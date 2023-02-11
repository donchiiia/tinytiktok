package main

import (
	"context"
	"tinytiktok/cmd/publish/pack"
	"tinytiktok/cmd/publish/service"
	"tinytiktok/dal/db"
	publish "tinytiktok/kitex_gen/publish"
	"tinytiktok/pkg/errno"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	currID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildPublishResp(err)
		return resp, nil
	}

	if len(req.Title) == 0 || len(req.Data) == 0 {
		resp = pack.BuildPublishResp(errno.ParamParseErr)
		return resp, nil
	}

	err = service.NewPublishActionService(ctx).PublishAction(req, currID)
	if err != nil {
		resp = pack.BuildPublishResp(err)
		return resp, nil
	}

	resp = pack.BuildPublishResp(errno.Success)
	return resp, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	currID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildPublishListResp(err)
		return resp, nil
	}

	// 获取followSet
	followSet, err := db.MGetFollowSet(ctx, currID)
	if err != nil {
		resp = pack.BuildPublishListResp(err)
		return resp, nil
	}

	// 获取favoriteSet
	favoriteSet, err := db.MGetFavoriteVideoSet(ctx, currID)
	if err != nil {
		resp = pack.BuildPublishListResp(err)
		return resp, nil
	}

	publishList, err := service.NewPublishListService(ctx).PublishList(req, followSet, favoriteSet)
	if err != nil {
		resp = pack.BuildPublishListResp(err)
		return resp, nil
	}
	resp = pack.BuildPublishListResp(errno.Success)
	resp.VideoList = publishList
	return resp, nil
}
