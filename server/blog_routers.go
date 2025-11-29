package server

import (
	"context"
	"net/http"

	"github.com/aws/smithy-go/ptr"
	"github.com/capyflow/Allspark-go/ds"
	"github.com/capyflow/blog/dao"
	"github.com/capyflow/blog/handle"
	"github.com/capyflow/blog/pkg"
	"github.com/capyflow/blog/service"
	"github.com/capyflow/vortex/v2"
)

// / 接口文件
func PrepareBlogRouters(ctx context.Context, cfg *pkg.Config, dServer *ds.DatabaseServer) []*vortex.VortexHttpRouter {

	// dao
	userDao := dao.NewNewUserDao(ctx, ptr.ToString(cfg.Group), dServer)

	// service
	userServ := service.NewUserService(ctx, cfg, userDao)

	// handle
	userHandle := handle.NewUserHandle(ctx, userServ)
	return appendBlogRouters(userHandle)
}

func appendBlogRouters(userHandle *handle.UserHandle) []*vortex.VortexHttpRouter {
	return []*vortex.VortexHttpRouter{
		vortex.AppendHttpRouter([]string{http.MethodPost}, "/console/login/pwd", userHandle.LoginByPassword, "根据密码登录"),
		vortex.AppendHttpRouter([]string{http.MethodPost}, "/console/login/email_code", userHandle.LoginByEmailCode, "根据邮箱验证码登录"),
	}
}
