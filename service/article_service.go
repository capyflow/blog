package service

import (
	"context"
	"time"

	"github.com/capyflow/Allspark-go/logx"
	"github.com/capyflow/blog/dao"
	"github.com/capyflow/blog/model"
	"github.com/capyflow/blog/pkg"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// randArticleId 生成随机文章ID
func randArticleId() string {
	alphabet := "abcdefghijklmnopqrstuvwxyz0123456789"
	id, _ := gonanoid.Generate(alphabet, 16) // 长度 16
	return "ai_" + id
}

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

// PublishArticle 发布文章
func (s *ArticleService) PublishArticle(ctx context.Context, title, content, category string) error {
	article := &model.ArticleInfo{
		ID:        randArticleId(),
		Title:     title,
		Content:   content,
		Category:  category,
		CreatedTs: time.Now().Unix(),
		UpdatedTs: time.Now().Unix(),
	}

	// 根据文章内容调用ai生成文章简介

	if err := s.articleDao.UpsertArticleInfo(ctx, article); err != nil {
		logx.Errorf("ArticleService|PublishArticle|UpsertArticleInfo|Error|%v", err)
		return err
	}
	if err := s.articleDao.AddArticleToCategoryList(ctx, article.ID, category); err != nil {
		logx.Errorf("ArticleService|PublishArticle|AddArticleToCategoryList|Error|%v", err)
		return err
	}
	logx.Infof("ArticleService|PublishArticle|Success|%s|%s|%s", article.ID, article.Title, article.Category)
	return nil
}

// 更新文章
func (s *ArticleService) UpdateArticle(ctx context.Context, id, title, content string) (*model.ArticleInfo, error) {
	article, err := s.articleDao.QueryArticleInfo(ctx, id)
	if err != nil {
		logx.Errorf("ArticleService|UpdateArticle|QueryArticleInfoById|Error|%v", err)
		return nil, err
	}
	if len(title) > 0 {
		article.Title = title
	}
	if len(content) > 0 {
		article.Content = content
	}
	article.UpdatedTs = time.Now().Unix()
	if err := s.articleDao.UpsertArticleInfo(ctx, article); err != nil {
		logx.Errorf("ArticleService|UpdateArticle|UpsertArticleInfo|Error|%v", err)
		return nil, err
	}
	return article, nil
}

// 删除文章
func (s *ArticleService) DeleteArticle(ctx context.Context, id string) error {
	if err := s.articleDao.DeleteArticleInfo(ctx, id); err != nil {
		logx.Errorf("ArticleService|DeleteArticle|DeleteArticleInfo|Error|%v", err)
		return err
	}
	return nil
}
