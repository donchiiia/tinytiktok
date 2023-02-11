package db

import (
	"context"
	"gorm.io/gorm"
	"tinytiktok/pkg/consts"
	"tinytiktok/pkg/errno"
)

type Favorite struct {
	gorm.Model
	UserID  int64 `gorm:"index:idx_userid;not null"`
	VideoID int64 `gorm:"not null"`
}

func (Favorite) TableName() string {
	return consts.FavoriteTableName
}

// AddFavorite 为 视频 添加喜欢
func AddFavorite(ctx context.Context, userID int64, videoID int64) error {
	// 事务操作
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 添加点赞数据
		user := new(User)
		if err := tx.WithContext(ctx).Find(user, userID).Error; err != nil {
			return err
		}
		video := new(Video)
		if err := tx.WithContext(ctx).Find(video, videoID).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(&user).Association("FavoriteVideos").Append(video); err != nil {
			return err
		}

		//2.改变 video 表中的 favorite_count
		res := tx.Model(video).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
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

// UnFavorite 取消 对视频的喜欢
func UnFavorite(ctx context.Context, userID int64, videoID int64) error {
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 将视频移除点赞列表
		user := new(User)
		if err := tx.WithContext(ctx).Find(user, userID).Error; err != nil {
			return err
		}
		video, err := GetAssociatedVideo(ctx, userID, videoID)
		if err != nil {
			return err
		}

		if err := tx.Unscoped().WithContext(ctx).Model(&user).Association("FavoriteVideos").Delete(video); err != nil {
			return err
		}

		// 修改视频点赞总数
		res := tx.Model(&Video{}).Update("favorite_count", gorm.Expr("favorite - ?", 1))
		if res.Error != nil {
			return err
		}
		if res.RowsAffected != 1 {
			return errno.DBErr
		}
		return nil
	})
	return err
}

// MGetFavoriteVideoList 获取 喜欢的视频 列表
func MGetFavoriteVideoList(ctx context.Context, userID int64) ([]*Video, error) {
	videoList := make([]*Video, 0)

	err := DB.WithContext(ctx).Model(&User{ID: userID}).Association("FavoriteVideos").Find(&videoList)
	if err != nil {
		return nil, err
	}
	return videoList, nil
}

// MGetFavoriteVideoSet 获取 喜欢的视频 集合
func MGetFavoriteVideoSet(ctx context.Context, userID int64) (map[int64]struct{}, error) {
	videoList := make([]*Video, 0)
	videoSet := make(map[int64]struct{}, 0)

	err := DB.WithContext(ctx).Model(&User{ID: userID}).Association("FavoriteVideos").Find(&videoList)
	if err != nil {
		return nil, err
	}
	for _, v := range videoList {
		videoSet[v.ID] = struct{}{}
	}
	return videoSet, nil
}

// GetAssociatedVideo 获取 关联视频
func GetAssociatedVideo(ctx context.Context, userID int64, videoID int64) (*Video, error) {
	user := new(User)
	err := DB.WithContext(ctx).Where(user, userID).Error
	if err != nil {
		return nil, err
	}
	video := new(Video)
	err = DB.WithContext(ctx).Model(video).Association("FavoriteVideos").Find(&video, videoID)
	if err != nil {
		return nil, err
	}
	return video, nil
}
