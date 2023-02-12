package service

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/message"
	"tinytiktok/pkg/errno"
)

type MessageActionService struct {
	ctx context.Context
}

// NewMessageActionService new MessageActionService
func NewMessageActionService(ctx context.Context) *MessageActionService {
	return &MessageActionService{ctx: ctx}
}

// MessageAction 点赞操作 点赞/取消点赞
func (s *MessageActionService) MessageAction(req *message.DouyinMessageActionRequest, fromUserID int64, createTime string) error {
	// 1-点赞
	if req.ActionType == 1 {
		return db.AddMessage(s.ctx, &db.Message{
			ToUserID:   req.ToUserId,
			FromUserID: fromUserID,
			CreateTime: createTime,
			Content:    req.Content,
		})
	}
	return errno.ActionTypeErr
}
