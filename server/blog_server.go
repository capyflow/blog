package server

import (
	"context"

	"github.com/aws/smithy-go/ptr"
	"github.com/capyflow/Allspark-go/ds"
	"github.com/capyflow/blog/pkg"
	"github.com/capyflow/vortex/v2"
)

type BlogServer struct {
	ctx context.Context
	v   *vortex.Vortex
}

// NewBlogServer 创建一个新的BlogServer实例
func NewBlogServer(ctx context.Context, cfg *pkg.Config) *BlogServer {
	bs := &BlogServer{
		ctx: ctx,
	}

	dServer := ds.InitDatabaseServer(ctx, cfg.DatabaseConfig, func(dbIdxs map[string]interface{}) {
		dbIdxs["user"] = 2
	})

	// 准备路由
	routers := PrepareBlogRouters(ctx, cfg, dServer)
	bs.v = vortex.BootStrap(ctx,
		vortex.WithPort(ptr.ToString(cfg.Port)),
		vortex.WithJwtSecretKey(cfg.Jwt.Secret),
		vortex.WithRouters(routers),
	)
	return bs
}

// 启动
func (bs *BlogServer) Start() {
	go bs.v.Start()
}

// 关闭
func (bs *BlogServer) Stop(ctx context.Context) error {
	return nil
}
