package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tinytiktok/cmd/api/biz/model/api"
	"tinytiktok/cmd/api/biz/rpc"
	"tinytiktok/cmd/feed/pack"
	"tinytiktok/kitex_gen/feed"
	"tinytiktok/pkg/errno"
)

// GetUserFeed .
// @router /douyin/feed/ [GET]
func GetUserFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinFeedRequest
	err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildFeedResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.GetFeed(ctx, &feed.DouyinFeedRequest{
		LatestTime: req.LatestTime,
		Token:      req.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildFeedResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}
