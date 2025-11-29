package main

import (
	"context"
	"flag"

	"github.com/capyflow/Allspark-go/conv"
	"github.com/capyflow/Allspark-go/logx"
	"github.com/capyflow/Allspark-go/system"
	"github.com/capyflow/blog/pkg"
	"github.com/capyflow/blog/server"
)

func main() {
	var cfgPath = flag.String("config", "./conf/blog.toml", "config file path")
	flag.Parse()
	cfg, err := pkg.LoadConfig(*cfgPath)
	if err != nil {
		panic(err)
	}
	logx.Infof("Config: %v", conv.ToJsonWithoutError(cfg))
	bs := server.NewBlogServer(context.Background(), cfg)
	bs.Start()
	system.GracefulShutdown(bs.Stop)
}
