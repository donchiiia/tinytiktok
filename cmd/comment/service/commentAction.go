package service

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/comment"
	"tinytiktok/pkg/errno"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

// CommentAction 评论操作：添加/删除
func (s *CommentActionService) CommentAction(req *comment.DouyinCommentActionRequest, currID int64) error {
	// 发布评论
	if req.ActionType == 1 {
		return db.CreateComment(s.ctx, &db.Comment{
			VideoID: req.VideoId,
			UserID:  currID,
			Content: *req.CommentText,
		})
	}
	// 删除评论
	if req.ActionType == 2 {
		return db.DeleteComment(s.ctx, *req.CommentId, req.VideoId)
	}
	return errno.ActionTypeErr
}
