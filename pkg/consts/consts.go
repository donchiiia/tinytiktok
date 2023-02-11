package consts

import "time"

const (
	// Service Name
	ApiServiceName      = "api"
	UserServiceName     = "user"
	FeedServiceName     = "feed"
	PublishServiceName  = "publish"
	CommentServiceName  = "comment"
	FavoriteServiceName = "favorite"
	RelationServiceName = "relation"

	// SQL Table Name
	UserTableName     = "user"
	VideoTableName    = "video"
	CommentTableName  = "comment"
	FollowTableName   = "follow"
	FriendTableName   = "friend"
	FavoriteTableName = "favorite"

	ETCDAddress     = "127.0.0.1:2379"
	ExportEndpoint  = ":4317"
	TCP             = "tcp"
	UserServiceAddr = ":8081"

	// 各服务地址
	FeedServiceAddr     = ":8082"
	FavoriteServiceAddr = ":8083"
	CommentServiceAddr  = ":8084"
	PublishServiceAddr  = ":8085"
	RelationServiceAddr = ":8086"

	SecretKey       = "tinytiktok"
	IdentityKey     = "id"
	MySQLDefaultDSN = "tiktokadmin:tiktokadmin@tcp(localhost:3306)/tinytiktok?charset=utf8&parseTime=True&loc=Local"
	DefaultLimit    = 10

	// minio 配置
	MinioEndpoint        = "localhost:9000"
	MinioAccessKeyID     = "minioadmin"
	MinioSecretAccessKey = "minioadmin"
	MinioUseSSL          = false
	Miniolocation        = "cn-south-1"
	MinioExpires         = time.Second * 60 * 60 * 24
	// ContentType
	VideoContentType = "video/mp4"
	CoverContentType = "image/jpeg"
	VideoBucketName  = "tiktok"
	CoverBucketName  = "cover"
	VideoSuffix      = ".mp4"
	CoverSuffix      = ".jpeg"

	//MySQL配置
	MySQLMaxIdleConns    = 10        //空闲连接池中连接的最大数量
	MySQLMaxOpenConns    = 100       //打开数据库连接的最大数量
	MySQLConnMaxLifetime = time.Hour //连接可复用的最大时间

)
