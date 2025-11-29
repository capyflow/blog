package service

import (
	"context"

	"github.com/capyflow/blog/dao"
	"github.com/capyflow/blog/pkg"
)

type UserService struct {
	ctx     context.Context
	userDao *dao.UserDao
}

// NewUserService 创建用户服务
func NewUserService(ctx context.Context, cfg *pkg.Config, userDao *dao.UserDao) *UserService {
	return &UserService{
		ctx:     ctx,
		userDao: userDao,
	}
}
