package service

import (
	"context"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/message"
	"tinytiktok/pack"
)

type MessageChatService struct {
	ctx context.Context
}

// NewMessageChatService creates a new MessageChatService
func NewMessageChatService(ctx context.Context) *MessageChatService {
	return &MessageChatService{
		ctx: ctx,
	}
}

// MessageChat 返回 toUserID的chat消息列表
func (s *MessageChatService) MessageChat(req *message.DouyinMessageChatRequest, fromUserID int64) ([]*message.Message, error) {
	messageList, err := db.MGetMessageList(s.ctx, req.ToUserId, fromUserID)
	if err != nil {
		return nil, err
	}

	messages, err := pack.MessageListInfo(s.ctx, messageList)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
