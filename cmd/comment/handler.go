package main

import (
	"context"
	"tinytiktok/cmd/comment/pack"
	"tinytiktok/cmd/comment/service"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/comment"
	"tinytiktok/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	currID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildCommentActionResp(err)
		return resp, nil
	}

	if currID == 0 {
		resp = pack.BuildCommentActionResp(errno.UserNotExistErr)
		return resp, nil
	}

	if req.VideoId == 0 {
		resp = pack.BuildCommentActionResp(errno.VideoNotExistErr)
		return resp, nil
	}

	if req.ActionType != 1 && req.ActionType != 2 {
		resp = pack.BuildCommentActionResp(errno.ActionTypeErr)
		return resp, nil
	}

	if req.ActionType == 1 && len([]rune(*req.CommentText)) > 255 {
		resp = pack.BuildCommentActionResp(errno.TextLenLimitExceededErr)
		return resp, nil
	}

	err = service.NewCommentActionService(ctx).CommentAction(req, currID)
	if err != nil {
		resp = pack.BuildCommentActionResp(err)
		return resp, nil
	}
	resp = pack.BuildCommentActionResp(errno.Success)
	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	currID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildCommentListResp(err)
		return resp, nil
	}

	// 获取followSet
	followSet, err := db.MGetFollowSet(ctx, currID)
	if err != nil {
		resp = pack.BuildCommentListResp(err)
		return resp, nil
	}

	commentList, err := service.NewCommentListService(ctx).CommentList(req, currID, followSet)
	if err != nil {
		resp = pack.BuildCommentListResp(err)
	}
	resp = pack.BuildCommentListResp(errno.Success)
	resp.CommentList = commentList
	return resp, nil
}
