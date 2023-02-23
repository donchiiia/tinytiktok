package db

import (
	"context"
	"gorm.io/gorm"
	"tinytiktok/pkg/consts"
	"tinytiktok/pkg/errno"
)

type Comment struct {
	gorm.Model
	ID      int64  `gorm:"primarykey"`
	Video   Video  `gorm:"foreignkey:VideoID"`
	VideoID int64  `gorm:"index:idx_videoid;not null"`
	User    User   `gorm:"foreignkey:UserID"`
	UserID  int64  `gorm:"index:idx_userid;not null"`
	Content string `gorm:"type:varchar(255);not null"`
}

func (Comment) TableName() string {
	return consts.CommentTableName
}

// CreateComment 创建 评论
func CreateComment(ctx context.Context, comment *Comment) error {
	// 评论要改变video表中的comment_count
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 新增评论
		if err := tx.Create(comment).Error; err != nil {
			return err
		}

		// 修改video表中的comment
		res := tx.Model(&Video{}).Where("id = ?", comment.VideoID).Update("comment_count", gorm.Expr("comment_count + ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			// 数据库更新错误
			return errno.DBErr
		}
		return nil
	})
	return err
}

// MGetVideoComments 通过 视频id 获取 评论列表
func MGetVideoComments(ctx context.Context, videoID int64) ([]*Comment, error) {
	var commentList []*Comment
	err := DB.WithContext(ctx).Model(&Comment{}).Where("video_id = ?", videoID).Find(&commentList).Error
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, commentID int64, videoID int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		comment := new(Comment)

		// 如果无法找到commentID记录就直接返回
		err := tx.Find(&comment, commentID).Error
		if err != nil {
			return err
		}

		// 找到记录就开始删除
		// 1. 删除comment记录
		//软删除
		err = tx.Delete(&comment).Error
		//物理删除
		// err := tx.Unscoped().Delete(&comment).Error
		if err != nil {
			return err
		}

		// 2.修改video表comment_count字段
		res := tx.Model(&Video{}).Where("id = ?", videoID).Update("comment_count", gorm.Expr("comment_count - ?", 1))
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errno.DBErr
		}
		return nil
	})
	return err
}
