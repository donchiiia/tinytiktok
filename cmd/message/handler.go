package main

import (
	"context"
	"time"
	"tinytiktok/cmd/message/pack"
	"tinytiktok/cmd/message/service"
	message "tinytiktok/kitex_gen/message"
	"tinytiktok/pkg/errno"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.DouyinMessageActionRequest) (resp *message.DouyinMessageActionResponse, err error) {
	fromUserID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildMessageActionResp(err)
		return resp, nil
	}

	if fromUserID == 0 || req.ToUserId == 0 {
		resp = pack.BuildMessageActionResp(errno.UserNotExistErr)
		return resp, nil
	}

	if req.ActionType != 1 {
		resp = pack.BuildMessageActionResp(errno.ActionTypeErr)
		return resp, nil
	}

	if req.ActionType == 1 && len([]rune(req.Content)) > 255 {
		resp = pack.BuildMessageActionResp(errno.TextLenLimitExceededErr)
		return resp, nil
	}

	err = service.NewMessageActionService(ctx).MessageAction(req, fromUserID, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		resp = pack.BuildMessageActionResp(err)
		return resp, nil
	}
	resp = pack.BuildMessageActionResp(errno.Success)
	return resp, nil
}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.DouyinMessageChatRequest) (resp *message.DouyinMessageChatResponse, err error) {
	fromUserID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildMessageChatResp(err)
		return resp, nil
	}

	if fromUserID == 0 || req.ToUserId == 0 {
		resp = pack.BuildMessageChatResp(errno.UserNotExistErr)
		return resp, nil
	}

	messageChat, err := service.NewMessageChatService(ctx).MessageChat(req, fromUserID)
	resp.MessageList = messageChat
	return resp, nil
}
