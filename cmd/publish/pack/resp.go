package pack

import (
	"tinytiktok/kitex_gen/publish"
	"tinytiktok/pkg/errno"
)

// 包装 register 的Response
func publishResp(err errno.ErrNo) *publish.DouyinPublishActionResponse {
	return &publish.DouyinPublishActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

func publishListResp(err errno.ErrNo) *publish.DouyinPublishListResponse {
	return &publish.DouyinPublishListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildPublishResp 由 error 构建PublishResponse
func BuildPublishResp(err error) *publish.DouyinPublishActionResponse {
	if err == nil {
		return publishResp(errno.Success)
	}

	return publishResp(errno.ConvertErr(err))
}

// BuildPublishListResp 由 error 构建PublishListResponse
func BuildPublishListResp(err error) *publish.DouyinPublishListResponse {
	if err == nil {
		return publishListResp(errno.Success)
	}

	return publishListResp(errno.ConvertErr(err))
}
