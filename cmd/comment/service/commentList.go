package service

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/comment"
	"tinytiktok/pack"
)

type CommentListService struct {
	ctx context.Context
}

func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{ctx: ctx}
}

// CommentList 获取 对应视频下的评论列表
func (s *CommentListService) CommentList(req *comment.DouyinCommentListRequest, currID int64, followSet map[int64]struct{}) (commentList []*comment.Comment, err error) {
	comments, err := db.MGetVideoComments(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	commentList, err = pack.CommentListInfo(s.ctx, comments, currID, followSet)
	if err != nil {
		return nil, err
	}
	return commentList, err
}
