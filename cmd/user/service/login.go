package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"tinytiktok/dal/db"
	"tinytiktok/kitex_gen/user"
	"tinytiktok/pkg/errno"
)

type LoginUserService struct {
	ctx context.Context
}

func NewLoginUserService(ctx context.Context) *LoginUserService {
	return &LoginUserService{
		ctx: ctx,
	}
}

// Login 登录
func (s *LoginUserService) Login(req *user.DouyinUserLoginRequest) (userID int64, err error) {
	// 计算传入 密码 的md5值
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	// 根据 username 查询db.User
	users, err := db.GetUsersByName(s.ctx, req.Username)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}

	// 密码校验
	u := users[0]
	if u.Password != passWord {
		return 0, errno.LoginErr
	}
	return u.ID, nil
}
