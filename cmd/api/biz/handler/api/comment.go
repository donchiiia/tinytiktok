package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tinytiktok/cmd/api/biz/model/api"
	"tinytiktok/cmd/api/biz/rpc"
	"tinytiktok/cmd/comment/pack"
	"tinytiktok/kitex_gen/comment"
	"tinytiktok/pkg/errno"
)

// CommentAction .
// @router /douyin/comment/action/ [POST]
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinCommentActionRequest
	err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildCommentActionResp(errno.ConvertErr(err)))
		return
	}

	rpcReq := &comment.DouyinCommentActionRequest{
		Token:      req.Token,
		VideoId:    req.VideoID,
		ActionType: req.ActionType,
	}
	if rpcReq.ActionType == 1 {
		rpcReq.CommentText = req.CommentText
	} else {
		rpcReq.CommentId = req.CommentID
	}

	resp, err := rpc.CommentAction(ctx, rpcReq)
	if err != nil {
		SendResponse(c, pack.BuildCommentActionResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}

// CommentList .
// @router /douyin/comment/list/ [GET]
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinCommentListRequest
	err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildCommentListResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.GetCommentList(ctx, &comment.DouyinCommentListRequest{
		Token:   req.Token,
		VideoId: req.VideoID,
	})
	if err != nil {
		SendResponse(c, pack.BuildCommentListResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}
