package service

import (
	"context"

	"github.com/capyflow/blog/dao"
	"github.com/capyflow/blog/pkg"
)

type ArticleService struct {
	ctx        context.Context
	articleDao *dao.ArticleDao
}

// NewArticleService 创建文章服务
func NewArticleService(ctx context.Context, cfg *pkg.Config, articleDao *dao.ArticleDao) *ArticleService {
	return &ArticleService{
		ctx:        ctx,
		articleDao: articleDao,
	}
}
