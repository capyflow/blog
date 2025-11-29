package handle

import (
	"context"

	"github.com/capyflow/blog/service"
)

type ArticleHandle struct {
	ctx         context.Context
	articleServ *service.ArticleService
}

// NewArticleHandle 创建文章处理层
func NewArticleHandle(ctx context.Context, articleServ *service.ArticleService) *ArticleHandle {
	return &ArticleHandle{
		ctx:         ctx,
		articleServ: articleServ,
	}
}
