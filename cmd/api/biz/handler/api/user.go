package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tinytiktok/cmd/api/biz/model/api"
	"tinytiktok/cmd/api/biz/rpc"
	"tinytiktok/cmd/user/pack"
	"tinytiktok/kitex_gen/user"
	"tinytiktok/pkg/errno"
)

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRegisterRequest

	err = c.BindAndValidate(&req)
	//err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildRegisterResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.Register(ctx, &user.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		SendResponse(c, pack.BuildRegisterResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserLoginRequest

	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c, pack.BuildLoginResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.Login(ctx, &user.DouyinUserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		SendResponse(c, pack.BuildLoginResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}

// GetUserByID .
// @router /douyin/user/ [GET]
func GetUserByID(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.DouyinUserRequest
	err = c.Bind(&req)
	if err != nil {
		SendResponse(c, pack.BuildUserInfoResp(errno.ConvertErr(err)))
		return
	}

	resp, err := rpc.GetUserByID(ctx, &user.DouyinUserRequest{
		UserId: req.UserID,
		Token:  req.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildUserInfoResp(errno.ConvertErr(err)))
		return
	}

	SendResponse(c, resp)
}
