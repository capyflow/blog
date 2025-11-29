package handle

import (
	"context"

	"github.com/capyflow/Allspark-go/logx"
	"github.com/capyflow/blog/pkg"
	"github.com/capyflow/blog/service"
	"github.com/capyflow/vortex/v2"
	"github.com/labstack/echo/v4"
)

type UserHandle struct {
	ctx      context.Context
	userServ *service.UserService
}

// NewUserHandle 创建用户处理句柄
func NewUserHandle(ctx context.Context, userServ *service.UserService) *UserHandle {
	return &UserHandle{
		ctx:      ctx,
		userServ: userServ,
	}
}

// 登录接口
func (lh *UserHandle) Login(ctx *vortex.Context) error {
	var req pkg.LoginReq
	if err := ctx.Bind(&req); nil != err {
		logx.Errorf("UserHandle|Login|Bind req failed|%v", err)
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.Success, echo.Map{
			"error": err.Error(),
		})
	}

	return nil
}
