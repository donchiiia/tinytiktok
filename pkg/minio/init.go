package minio

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"tinytiktok/pkg/consts"
)

var minioClient *minio.Client

func init() {
	// Initialize minio client object.
	client, err := minio.New(consts.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(consts.MinioAccessKeyID, consts.MinioSecretAccessKey, ""),
		Secure: consts.MinioUseSSL,
	})
	if err != nil {
		klog.Debugf("Minio Client failed to init:", err)
	}
	klog.Debugf("Minio Client init successfully!")
	minioClient = client

	err = CreateBucket(consts.VideoBucketName)
	if err != nil {
		klog.Errorf("Minio Client can't create bucket [%v]: %#v\n", consts.VideoBucketName, client)
	}
	klog.Debugf("Minio Client create bucket [%v] successfully", consts.VideoBucketName)
	err = CreateBucket(consts.CoverBucketName)
	if err != nil {
		klog.Errorf("Minio Client can't create bucket [%v]: %#v\n", consts.CoverBucketName, client)
	}
	klog.Debugf("Minio Client create bucket [%v] successfully", consts.CoverBucketName)
}
