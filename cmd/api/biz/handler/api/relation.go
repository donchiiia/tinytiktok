package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tinytiktok/cmd/api/biz/model/api"
	"tinytiktok/cmd/api/biz/rpc"
	"tinytiktok/cmd/relation/pack"
	"tinytiktok/kitex_gen/relation"
	"tinytiktok/pkg/errno"
)

// RelationAction .
// @router /douyin/relation/action/ [POST]
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationActionRequest

	err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildRelationActionResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.FollowAction(ctx, &relation.DouyinRelationActionRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserID,
		ActionType: req.ActionType,
	})
	if err != nil {
		SendResponse(c, pack.BuildRelationActionResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}

// RelationFollowList .
// @router /douyin/relation/follow/list/ [GET]
func RelationFollowList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFollowListRequest
	err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildFollowListResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.GetFollowList(ctx, &relation.DouyinRelationFollowListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildFollowListResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}

// RelationFollowerList .
// @router /douyin/relation/follower/list/ [GET]
func RelationFollowerList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFollowerListRequest
	err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildFollowerListResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.GetFollowerList(ctx, &relation.DouyinRelationFollowerListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildFollowerListResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}

// RelationFriendList .
// @router /douyin/relation/friend/list/ [GET]
func RelationFriendList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinRelationFriendListRequest
	err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildFriendListResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.GetFriendList(ctx, &relation.DouyinRelationFriendListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildFriendListResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}
