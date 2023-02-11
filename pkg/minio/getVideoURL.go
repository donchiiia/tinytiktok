package minio

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/url"
	"tinytiktok/pkg/consts"
)

func GetVideoURL(objectName string) (*url.URL, error) {
	// Set request parameters
	reqParams := make(url.Values)

	// Gernerate presigned get object url.
	presignedURL, err := minioClient.PresignedGetObject(context.Background(), consts.VideoBucketName, objectName, consts.MinioExpires, reqParams)
	if err != nil {
		klog.Errorf("Fetch video url failed %v", err)
		return nil, nil
	}
	klog.Infof("Fetch video url successfully. [video url]%v", presignedURL)
	return presignedURL, nil
}
