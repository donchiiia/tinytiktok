package db

import (
	"context"
	"gorm.io/gorm"
	"tinytiktok/pkg/consts"
)

type Message struct {
	gorm.Model
	ID         int64  `gorm:"primarykey"`
	ToUserID   int64  `gorm:"index:idx_to_user_id;not null"`
	FromUserID int64  `gorm:"index:idx_from_user_id;not null"`
	CreateTime string `gorm:"column:create_time;index:idx_create_time" `
	Content    string `gorm:"type:varchar(255);not null"`
}

func (Message) TableName() string {
	return consts.MessageTableName
}

// AddMessage 添加消息至数据库
func AddMessage(ctx context.Context, message *Message) error {
	// 涉及数据库写操作尽量使用事务，防止视频数据传输中断造成数据不一致问题（可在上层使用消息队列优化）
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(message).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// MGetMessageList 根据id获取消息列表
func MGetMessageList(ctx context.Context, toUserID int64, fromUserID int64) ([]*Message, error) {
	var messageList []*Message
	// model指定表名
	err := DB.Model(&Message{}).WithContext(ctx).Where(&Message{
		ToUserID:   toUserID,
		FromUserID: fromUserID,
	}).Find(&messageList).Error
	if err != nil {
		return nil, err
	}
	return messageList, nil
}
