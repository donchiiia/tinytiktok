package pack

import (
	"tinytiktok/kitex_gen/message"
	"tinytiktok/pkg/errno"
)

// 包装 message 的Response
func messageActionResp(err errno.ErrNo) *message.DouyinMessageActionResponse {
	return &message.DouyinMessageActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

func messageChatResp(err errno.ErrNo) *message.DouyinMessageChatResponse {
	return &message.DouyinMessageChatResponse{StatusCode: int32(err.ErrCode), StatusMsg: &err.ErrMsg}
}

// BuildMessageActionResp 由 error 构建MessageActionResponse
func BuildMessageActionResp(err error) *message.DouyinMessageActionResponse {
	if err == nil {
		return messageActionResp(errno.Success)
	}

	return messageActionResp(errno.ConvertErr(err))
}

// BuildMessageChatResp 由 error 构建MessageChatResponse
func BuildMessageChatResp(err error) *message.DouyinMessageChatResponse {
	if err == nil {
		return messageChatResp(errno.Success)
	}

	return messageChatResp(errno.ConvertErr(err))
}
