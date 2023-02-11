package minio

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"tinytiktok/pkg/consts"
)

func CreateBucket(bucketName string) error {
	ctx := context.Background()

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: consts.Miniolocation})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			klog.Debugf("Bucket %s already exists\n", bucketName)
			return nil
		} else {
			return err
		}
	} else {
		klog.Errorf("Successfully created %s\n", bucketName)
	}
	return nil
}
