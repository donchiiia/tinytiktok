package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tinytiktok/cmd/api/biz/model/api"
	"tinytiktok/cmd/api/biz/rpc"
	"tinytiktok/cmd/favorite/pack"
	"tinytiktok/kitex_gen/favorite"
	"tinytiktok/pkg/errno"
)

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFavoriteActionRequest
	err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.FavoriteAction(ctx, &favorite.DouyinFavoriteActionRequest{
		Token:      req.Token,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	})
	if err != nil {
		SendResponse(c, pack.BuildFavoriteActionResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFavoriteListRequest
	err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.GetFavoriteList(ctx, &favorite.DouyinFavoriteListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildFavoriteListResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}
