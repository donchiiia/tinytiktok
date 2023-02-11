package pack

import (
	"tinytiktok/kitex_gen/user"
	"tinytiktok/pkg/errno"
)

// 包装 register 的Response
func registerResp(err errno.ErrNo) *user.DouyinUserRegisterResponse {
	return &user.DouyinUserRegisterResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}
func loginResp(err errno.ErrNo) *user.DouyinUserLoginResponse {
	return &user.DouyinUserLoginResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}
func userInfoResp(err errno.ErrNo) *user.DouyinUserResponse {
	return &user.DouyinUserResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildRegisterResp 由 error 构建RegisterResponse
func BuildRegisterResp(err error) *user.DouyinUserRegisterResponse {
	// 无错误，兜底返回
	if err == nil {
		return registerResp(errno.Success)
	}

	return registerResp(errno.ConvertErr(err))
}

// 由 error 构建LoginResponse
func BuildLoginResp(err error) *user.DouyinUserLoginResponse {
	// 无错误，兜底返回
	if err == nil {
		return loginResp(errno.Success)
	}

	return loginResp(errno.ConvertErr(err))
}

// 由 error 构建UserInfoResponse
func BuildUserInfoResp(err error) *user.DouyinUserResponse {
	// 无错误，兜底返回
	if err == nil {
		return userInfoResp(errno.Success)
	}

	return userInfoResp(errno.ConvertErr(err))
}
