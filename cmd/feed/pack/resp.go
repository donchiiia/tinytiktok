package pack

import (
	"tinytiktok/kitex_gen/feed"
	"tinytiktok/pkg/errno"
)

// 包装 feed 的Response
func feedResp(err errno.ErrNo) *feed.DouyinFeedResponse {
	return &feed.DouyinFeedResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// 由 error 构建FeedResponse
func BuildFeedResp(err error) *feed.DouyinFeedResponse {
	if err == nil {
		return feedResp(errno.Success)
	}

	return feedResp(errno.ConvertErr(err))
}
