package service

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"strings"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/publish"
	"tinytiktok/pkg/consts"
	"tinytiktok/pkg/ffmpeg"
	"tinytiktok/pkg/minio"
)

type PublishActionService struct {
	ctx context.Context
}

// NewPublishActionService new PublishActionService
func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

// PublishAction 发布视频
func (s *PublishActionService) PublishAction(req *publish.DouyinPublishActionRequest, currID int64) error {
	minioVideoBucket := consts.VideoBucketName
	minioCoverBucket := consts.CoverBucketName

	videoData := req.Data
	reader := bytes.NewReader(videoData)

	// 创建保存在bucket中的视频名称
	vN, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	vName := vN.String() + consts.VideoSuffix

	// 上传视频，并获得视频链接
	videoURL, err := minio.FileUploader(minioVideoBucket, vName, reader, int64(len(videoData)), consts.VideoContentType)
	if err != nil {
		return err
	}

	// 创建封面名称
	cN, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	cName := cN.String() + consts.CoverSuffix

	// 获取视频封面
	coverData, err := ffmpeg.ReadFrameAsJpeg(videoURL.String(), 1)
	if err != nil {
		return err
	}

	// 上传封面
	coverReader := bytes.NewReader(coverData)
	coverURL, err := minio.FileUploader(minioCoverBucket, cName, coverReader, int64(len(coverData)), consts.CoverContentType)
	if err != nil {
		return err
	}

	return db.CreateVideoInfo(s.ctx, &db.Video{
		Title:         req.Title,
		AuthorID:      currID,
		VideoName:     vName,
		CoverName:     cName,
		PlayUrl:       strings.Split(videoURL.String(), "?")[0],
		CoverUrl:      strings.Split(coverURL.String(), "?")[0],
		FavoriteCount: 0,
		CommentCount:  0,
	})

}
