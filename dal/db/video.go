package db

import (
	"context"
	"gorm.io/gorm"
	"time"
	"tinytiktok/pkg/consts"
)

type Video struct {
	gorm.Model
	ID            int64     `gorm:"primarykey"`
	UpdatedAt     time.Time `gorm:"column:update_time;not null;index:idx_update_time" ` // feed流需要查询视频上传时间戳，建立上传时间索引
	Author        User      `gorm:"foreignkey:AuthorID"`
	Title         string    `gorm:"type:varchar(50);not null"`
	AuthorID      int64     `gorm:"index:idx_authorid;not null"` // 业务中要用户发布视频列表，建立作者ID索引
	VideoName     string    `gorm:"type:varchar(255);not null"`
	CoverName     string    `gorm:"type:varchar(255);not null"`
	PlayUrl       string    `gorm:"type:varchar(255);not null"`
	CoverUrl      string    `gorm:"type:varchar(255)"`
	FavoriteCount int64     `gorm:"default:0"`
	CommentCount  int64     `gorm:"default:0"`
}

func (Video) TableName() string {
	return consts.VideoTableName
}

// MGetFeed 获取视频feed流
// 因为feedrequest中lastest_time为可选参数，故参数中为指针
func MGetFeed(ctx context.Context, lastestTime *int64, limit int) ([]*Video, error) {
	videos := make([]*Video, 0)

	// 判断最近时间，为空则返回当前时间
	if lastestTime == nil || *lastestTime == 0 {
		currentTime := time.Now().Unix()
		lastestTime = &currentTime
	}

	// 返回以时间倒序视频列表
	err := DB.WithContext(ctx).Limit(limit).Order("update_time desc").Find(&videos, "update_time < ?", time.UnixMilli(*lastestTime)).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// CreateVideoInfo 在数据库中创建上传的视频信息
func CreateVideoInfo(ctx context.Context, video *Video) error {
	// 涉及数据库写操作尽量使用事务，防止视频数据传输中断造成数据不一致问题（可在上层使用消息队列优化）
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(video).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// MGetPublishVideoList 获取发布视频列表
func MGetPublishVideoList(ctx context.Context, authorID int64) ([]*Video, error) {
	var videoList []*Video
	// model指定表名
	err := DB.Model(&Video{}).WithContext(ctx).Where(&Video{AuthorID: authorID}).Find(&videoList).Error
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
