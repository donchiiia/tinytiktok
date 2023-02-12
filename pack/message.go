package pack

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/message"
)

// CommentInfo 将 db.Messsage 包装成 message.Message
func MessageInfo(dbMessage *db.Message) *message.Message {
	c := &message.Message{
		ToUserId:   dbMessage.ToUserID,
		FromUserId: dbMessage.FromUserID,
		Content:    dbMessage.Content,
	}
	return c
}

// MessageListInfo 将 db.Message列表包装成 message.Message列表
func MessageListInfo(ctx context.Context, dbMessageList []*db.Message) ([]*message.Message, error) {
	messageList := make([]*message.Message, 0)
	for _, v := range dbMessageList {
		message, err := db.MGetMessageList(ctx, v.ToUserID, v.FromUserID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		messageList = append(messageList, MessageInfo(message[0]))
	}
	return messageList, nil
}
