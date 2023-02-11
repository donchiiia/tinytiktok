package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	"tinytiktok/dal"
	"tinytiktok/kitex_gen/favorite/favoriteservice"
	"tinytiktok/pkg/consts"
	"tinytiktok/pkg/jwt"
	"tinytiktok/pkg/mw"
)

var Jwt *jwt.JWT

func InitJWT() {
	Jwt = jwt.NewJWT([]byte(consts.SecretKey))
}

func Init() {
	// klog init
	klog.SetLogger(logrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)

	dal.Init()
	InitJWT()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr(consts.TCP, consts.FavoriteServiceAddr)
	if err != nil {
		panic(err)
	}

	Init()

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.FavoriteServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)

	svr := favoriteservice.NewServer(new(FavoriteServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.FavoriteServiceName}),
	)

	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
