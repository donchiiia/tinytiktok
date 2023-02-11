package pack

import (
	"tinytiktok/kitex_gen/favorite"
	"tinytiktok/pkg/errno"
)

// 包装 favorite 的Response
func favoriteActionResp(err errno.ErrNo) *favorite.DouyinFavoriteActionResponse {
	return &favorite.DouyinFavoriteActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

func favoriteListResp(err errno.ErrNo) *favorite.DouyinFavoriteListResponse {
	return &favorite.DouyinFavoriteListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildFavoriteActionResp 由 error 构建FavoriteActionResponse
func BuildFavoriteActionResp(err error) *favorite.DouyinFavoriteActionResponse {
	if err == nil {
		return favoriteActionResp(errno.Success)
	}

	return favoriteActionResp(errno.ConvertErr(err))
}

// BuildFavoriteListResp 由 error 构建FavoriteListResponse
func BuildFavoriteListResp(err error) *favorite.DouyinFavoriteListResponse {
	if err == nil {
		return favoriteListResp(errno.Success)
	}

	return favoriteListResp(errno.ConvertErr(err))
}
