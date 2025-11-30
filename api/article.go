package api

import "github.com/capyflow/blog/model"

// PublishArticleReq 发布文章请求参数
type PublishArticleReq struct {
	Title    string `json:"title"`    // 文章标题
	Content  string `json:"content"`  // 文章内容
	Category string `json:"category"` // 文章分类
}

// UpdateArticleReq 更新文章请求参数
type UpdateArticleReq struct {
	ID      string `json:"id"`      // 文章id
	Title   string `json:"title"`   // 文章标题
	Content string `json:"content"` // 文章内容
}

// UpdateArticleResp 更新文章响应参数
type UpdateArticleResp struct {
	ArticleInfo *model.ArticleInfo `json:"article_info"`
}

// 删除文章请求参数
type DeleteArticleReq struct {
	ID string `param:"id"` // 文章id
}
