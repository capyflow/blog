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
	articleDao := dao.NewArticleDao(ctx, ptr.ToString(cfg.Group), dServer)

	// service
	userServ := service.NewUserService(ctx, cfg, userDao)
	articleServ := service.NewArticleService(ctx, cfg, articleDao)

	// handle
	userHandle := handle.NewUserHandle(ctx, userServ)
	articleHandle := handle.NewArticleHandle(ctx, articleServ)
	return appendBlogRouters(userHandle, articleHandle)
}

func appendBlogRouters(userHandle *handle.UserHandle, articleHandle *handle.ArticleHandle) []*vortex.VortexHttpRouter {
	return []*vortex.VortexHttpRouter{
		// 用户接口
		vortex.AppendHttpRouter([]string{http.MethodPost}, "/console/login/pwd", userHandle.LoginByPassword, "根据密码登录"),
		vortex.AppendHttpRouter([]string{http.MethodPost}, "/console/login/email_code", userHandle.LoginByEmailCode, "根据邮箱验证码登录"),

		// 文章接口
		vortex.AppendHttpRouter([]string{http.MethodPost}, "/console/article/publish", articleHandle.HandlePublishArticle, "发布文章"),
		vortex.AppendHttpRouter([]string{http.MethodPost}, "/console/article/update/:id", articleHandle.HandleUpdateArticle, "更新文章"),
		vortex.AppendHttpRouter([]string{http.MethodDelete}, "/console/article/delete/:id", articleHandle.HandleDeleteArticle, "删除文章"),
	}
}
