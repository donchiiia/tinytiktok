package minio

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
	"tinytiktok/pkg/consts"
)

func FileUploader(bucketName string, objectName string, reader io.Reader, fileSize int64, contentType string) (*url.URL, error) {
	info, err := minioClient.PutObject(context.Background(), bucketName, objectName, reader, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		klog.Errorf("Upload %v of size %d to buckets %v failed: %s", objectName, info.Size, bucketName, err)
		return nil, err
	}
	klog.Infof("Successfully uploaded %s of size %d\n", objectName, info.Size)

	var urlString *url.URL
	if contentType == consts.VideoContentType {
		urlString, err = GetVideoURL(objectName)
	} else if contentType == consts.CoverContentType {
		urlString, err = GetCoverURL(objectName)
	}
	if err != nil {
		return nil, err
	}
	return urlString, nil

}
