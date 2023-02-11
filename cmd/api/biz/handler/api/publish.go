package api

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"io"
	"mime/multipart"
	"tinytiktok/cmd/api/biz/model/api"
	"tinytiktok/cmd/api/biz/rpc"
	"tinytiktok/cmd/publish/pack"
	"tinytiktok/kitex_gen/publish"
	"tinytiktok/pkg/errno"
)

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var err error
	//var req api.DouyinPublishActionRequest
	// idl不支持生成文件内容，需要手动绑定传递而来的data
	//err = c.BindAndValidate(&req)

	token := c.PostForm("token")
	title := c.PostForm("title")

	fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ConvertErr(err)))
		return
	}

	file, err := fileHeader.Open()
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			SendResponse(c, pack.BuildPublishResp(errno.ServiceErr))
		}
	}(file)
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ConvertErr(err)))
		return
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		SendResponse(c, pack.BuildPublishResp(err))
		return
	}

	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.PublishAction(ctx, &publish.DouyinPublishActionRequest{
		Token: token,
		Data:  buf.Bytes(),
		Title: title,
	})
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}

// PublishList .
// @router /douyin/publish/list/ [GET]
func PublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinPublishListRequest
	err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildPublishListResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.GetPublishList(ctx, &publish.DouyinPublishListRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildPublishListResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}
