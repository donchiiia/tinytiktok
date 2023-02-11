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

type RegisterUserService struct {
	ctx context.Context
}

func NewRegisterUserService(ctx context.Context) *RegisterUserService {
	return &RegisterUserService{
		ctx: ctx,
	}
}

// Register 注册
func (s *RegisterUserService) Register(req *user.DouyinUserRegisterRequest) error {
	users, err := db.GetUsersByName(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	// 将密码md5加密后再存入数据库
	h := md5.New()
	if _, err = io.WriteString(h, req.Password); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.Username,
		Password: password,
	}})
}
