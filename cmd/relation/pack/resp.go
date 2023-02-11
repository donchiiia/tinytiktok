package pack

import (
	"tinytiktok/kitex_gen/relation"
	"tinytiktok/pkg/errno"
)

// 包装 follow 的Response
func relationActionResp(err errno.ErrNo) *relation.DouyinRelationActionResponse {
	return &relation.DouyinRelationActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

func followListResp(err errno.ErrNo) *relation.DouyinRelationFollowListResponse {
	return &relation.DouyinRelationFollowListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

func followerListResp(err errno.ErrNo) *relation.DouyinRelationFollowerListResponse {
	return &relation.DouyinRelationFollowerListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// 包装 friend 的Response
func friendListResp(err errno.ErrNo) *relation.DouyinRelationFriendListResponse {
	return &relation.DouyinRelationFriendListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildRelationActionResp 由 error 构建RelationActionResponse
func BuildRelationActionResp(err error) *relation.DouyinRelationActionResponse {
	if err == nil {
		return relationActionResp(errno.Success)
	}

	return relationActionResp(errno.ConvertErr(err))
}

// BuildFollowListResp 由 error 构建RelationFollowListResponse
func BuildFollowListResp(err error) *relation.DouyinRelationFollowListResponse {
	if err == nil {
		return followListResp(errno.Success)
	}

	return followListResp(errno.ConvertErr(err))
}

// BuildFollowerListResp 由 error 构建RelationFollowerListResponse
func BuildFollowerListResp(err error) *relation.DouyinRelationFollowerListResponse {
	if err == nil {
		return followerListResp(errno.Success)
	}

	return followerListResp(errno.ConvertErr(err))
}

// 由 error 构建RelationFriendListResponse
func BuildFriendListResp(err error) *relation.DouyinRelationFriendListResponse {
	if err == nil {
		return friendListResp(errno.Success)
	}

	return friendListResp(errno.ConvertErr(err))
}
