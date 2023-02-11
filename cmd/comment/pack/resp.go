package pack

import (
	"tinytiktok/kitex_gen/comment"
	"tinytiktok/pkg/errno"
)

// 包装 comment 的Response
func commentActionResp(err errno.ErrNo) *comment.DouyinCommentActionResponse {
	return &comment.DouyinCommentActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}
func commentListResp(err errno.ErrNo) *comment.DouyinCommentListResponse {
	return &comment.DouyinCommentListResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildCommentActionResp 由 error 构建CommentActionResponse
func BuildCommentActionResp(err error) *comment.DouyinCommentActionResponse {
	if err == nil {
		return commentActionResp(errno.Success)
	}

	return commentActionResp(errno.ConvertErr(err))
}

// BuildCommentListResp 由 error 构建CommentListResponse
func BuildCommentListResp(err error) *comment.DouyinCommentListResponse {
	if err == nil {
		return commentListResp(errno.Success)
	}

	return commentListResp(errno.ConvertErr(err))
}
