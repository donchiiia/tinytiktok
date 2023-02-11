package main

import (
	"context"
	"tinytiktok/cmd/user/pack"
	"tinytiktok/cmd/user/service"
	"tinytiktok/kitex_gen/user"
	"tinytiktok/pkg/errno"
	"tinytiktok/pkg/jwt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	if err = req.IsValid(); err != nil {
		resp = pack.BuildRegisterResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewRegisterUserService(ctx).Register(req)
	if err != nil {
		resp = pack.BuildRegisterResp(err)
		return resp, nil
	}

	resp = pack.BuildRegisterResp(errno.Success)

	// 当用户注册成功，直接跳转登录
	loginResponse, err := s.Login(ctx, (*user.DouyinUserLoginRequest)(req))
	if err != nil {
		resp = (*user.DouyinUserRegisterResponse)(pack.BuildLoginResp(err))
		return resp, nil
	}
	return (*user.DouyinUserRegisterResponse)(loginResponse), nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	if err = req.IsValid(); err != nil {
		resp = pack.BuildLoginResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewLoginUserService(ctx).Login(req)
	if err != nil {
		resp = pack.BuildLoginResp(err)
		return resp, nil
	}

	token, err := Jwt.CreateToken(jwt.CustomClaims{
		Id: uid,
	})
	if err != nil {
		resp = pack.BuildLoginResp(errno.SignatureInvalidErr)
		return resp, nil
	}

	resp = pack.BuildLoginResp(errno.Success)
	resp.UserId = uid
	resp.Token = token
	return resp, nil
}

// GetUserByID implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserByID(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	currID, err := Jwt.GetUserIDFromToken(req.Token)
	if err != nil {
		resp = pack.BuildUserInfoResp(err)
	}

	if err = req.IsValid(); err != nil {
		resp = pack.BuildUserInfoResp(errno.ParamErr)
		return resp, nil
	}

	users, err := service.NewGetUserService(ctx).GetUser(req, currID)
	if err != nil {
		resp = pack.BuildUserInfoResp(err)
		return resp, nil
	}

	resp = pack.BuildUserInfoResp(errno.Success)
	resp.User = users
	return resp, nil
}
