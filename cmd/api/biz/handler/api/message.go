package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"tinytiktok/cmd/api/biz/model/api"
	"tinytiktok/cmd/message/pack"
	"tinytiktok/pkg/errno"
)

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinMessageActionRequest
	err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildMessageActionResp(errno.ConvertErr(err)))
	}

	resp := new(api.DouyinMessageActionResponse)

	c.JSON(consts.StatusOK, resp)
}

// MessageChat .
// @router /douyin/message/chat/ [GET]
func MessageChat(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinMessageChatRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.DouyinMessageChatResponse)

	c.JSON(consts.StatusOK, resp)
}
