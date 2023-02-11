package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"tinytiktok/kitex_gen/publish"
	"tinytiktok/kitex_gen/publish/publishservice"
	"tinytiktok/pkg/consts"
	"tinytiktok/pkg/errno"
	"tinytiktok/pkg/mw"
)

var publishClient publishservice.Client

func initPublishRPC() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := publishservice.NewClient(
		consts.PublishServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	publishClient = c
}

// 传递视频投稿请求，获取rpc响应
func PublishAction(ctx context.Context, req *publish.DouyinPublishActionRequest) (resp *publish.DouyinPublishActionResponse, err error) {
	resp, err = publishClient.PublishAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}

// 传递获取投稿列表请求，获取rpc响应
func GetPublishList(ctx context.Context, req *publish.DouyinPublishListRequest) (resp *publish.DouyinPublishListResponse, err error) {
	resp, err = publishClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(int(resp.StatusCode), *resp.StatusMsg)
	}
	return resp, nil
}
