package dao

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/capyflow/Allspark-go/ds"
	"github.com/capyflow/Allspark-go/logx"
	"github.com/capyflow/blog/model"
	"github.com/redis/go-redis/v9"
)

// ArticleCategories 文章分类
var (
	ArticleCategories = struct {
		Default       string // 默认文章类型
		Project       string // 项目文章类型
		Entertainment string // 娱乐文章类型
	}{
		Default:       "default",
		Project:       "project",
		Entertainment: "entertainment",
	}
	categoryList = []string{
		ArticleCategories.Default,
		ArticleCategories.Project,
		ArticleCategories.Entertainment,
	}
)

// buildArticleInfoKey 构建文章信息键
func buildArticleInfoKey(group, articleID string) string {
	return fmt.Sprintf("%s:article:%s:info", group, articleID)
}

// 构建文件类别listKey
func buildArticleCategoryListKey(group, category string) string {
	return fmt.Sprintf("%s:article:category:list:%s", group, category)
}

type ArticleDao struct {
	ctx   context.Context
	group string
	rdb   *redis.Client
}

// NewArticleDao 创建文章 DAO
func NewArticleDao(ctx context.Context, group string, dServer *ds.DatabaseServer) *ArticleDao {
	rdb, ok := dServer.GetRedis("article")
	if !ok {
		panic("article redis not found")
	}
	return &ArticleDao{
		ctx: ctx,
		rdb: rdb,
	}
}

// 根据文章id查询文章
func (a *ArticleDao) QueryArticleInfoById(articleID string) (*model.ArticleInfo, error) {
	// 从 Redis 中获取文章详情
	articleJSON, err := a.rdb.Get(a.ctx, buildArticleInfoKey(a.group, articleID)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.New("article not found")
		}
		return nil, err
	}
	// 反序列化文章详情
	var article model.ArticleInfo
	if err := json.Unmarshal([]byte(articleJSON), &article); err != nil {
		return nil, err
	}

	return &article, nil
}

// 根据文件类别查询文件列表
func (a *ArticleDao) QueryArticleListByCategory(category string) ([]*model.ArticleSummary, error) {
	// 从 Redis 中获取文章列表
	key := buildArticleCategoryListKey(a.group, category)
	ids, err := a.rdb.ZRange(a.ctx, key, 0, -1).Result()
	if err != nil {
		logx.Errorf("ArticleDao|QueryArticleListByCategory|Error|%v|%s", err, category)
		if err == redis.Nil {
			return nil, errors.New("article category not found")
		}
		return nil, err
	}
	// 反序列化文章列表
	articles := make([]*model.ArticleSummary, 0, len(ids))
	for _, id := range ids {
		articleJSON, err := a.rdb.Get(a.ctx, buildArticleInfoKey(a.group, id)).Result()
		if err != nil {
			logx.Errorf("ArticleDao|QueryArticleListByCategory|Get|Error|%v|%s", err, id)
			if err == redis.Nil {
				continue
			}
			return nil, err
		}
		var article model.ArticleSummary
		if err := json.Unmarshal([]byte(articleJSON), &article); err != nil {
			logx.Errorf("ArticleDao|QueryArticleListByCategory|Unmarshal|Error|%v|%s", err, id)
			return nil, err
		}
		articles = append(articles, &article)
	}
	return articles, nil
}

// 查询所有文件类别对应的文章列表
func (a *ArticleDao) QueryAllArticleListByCategory(category string) ([]*model.ArticleCategory, error) {
	articleCategories := make([]*model.ArticleCategory, 0, len(categoryList))
	for _, category := range categoryList {
		articles, err := a.QueryArticleListByCategory(category)
		if err != nil {
			logx.Errorf("ArticleDao|QueryAllArticleListByCategory|Error|%v|%s", err, category)
			return nil, err
		}
		articleCategories = append(articleCategories, &model.ArticleCategory{
			Category: category,
			Articles: articles,
		})
	}
	return articleCategories, nil
}

// 更新文章
func (a *ArticleDao) UpdateArticleInfo(ctx context.Context, article *model.ArticleInfo) error {
	// 序列化文章详情
	articleJSON, err := json.Marshal(article)
	if err != nil {
		logx.Errorf("ArticleDao|UpdateArticleInfo|Marshal|Error|%v|%s", err, article.ID)
		return err
	}
	// 存储文章详情到 Redis
	key := buildArticleInfoKey(a.group, article.ID)
	if err := a.rdb.Set(ctx, key, articleJSON, 0).Err(); err != nil {
		logx.Errorf("ArticleDao|UpdateArticleInfo|Set|Error|%v|%s", err, key)
		return err
	}
	return nil
}

// 删除文章
func (a *ArticleDao) DeleteArticleInfo(ctx context.Context, articleId string) error {
	// 查询文章是否存在
	article, err := a.QueryArticleInfoById(articleId)
	if err != nil {
		logx.Errorf("ArticleDao|DeleteArticleInfo|QueryArticleInfoById|Error|%v|%s", err, articleId)
		return err
	}
	// 从 Redis 中删除文章详情
	key := buildArticleInfoKey(a.group, articleId)
	if err := a.rdb.Del(ctx, key).Err(); err != nil {
		logx.Errorf("ArticleDao|DeleteArticleInfo|Del|Error|%v|%s", err, key)
		return err
	}
	// 从对应文件类别列表中删除文章
	categoryKey := buildArticleCategoryListKey(a.group, article.Category)
	if err := a.rdb.ZRem(ctx, categoryKey, articleId).Err(); err != nil {
		logx.Errorf("ArticleDao|DeleteArticleInfo|ZRem|Error|%v|%s", err, categoryKey)
		return err
	}
	return nil
}
