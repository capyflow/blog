package model

// ArticleInfo 文章信息
type ArticleInfo struct {
	ID        string `json:"id"`         // 文章id
	Title     string `json:"title"`      // 文章标题
	Content   string `json:"content"`    // 文章内容
	CreatedTs int64  `json:"created_ts"` // 创建时间
	UpdatedTs int64  `json:"updated_ts"` // 更新时间
	Category  string `json:"category"`   // 文章分类
}

// 文章的摘要信息
type ArticleSummary struct {
	ID        string `json:"id"`         // 文章id
	Title     string `json:"title"`      // 文章标题
	CreatedTs int64  `json:"created_ts"` // 创建时间
	Category  string `json:"category"`   // 文章分类
}

// 文件类别对应文章
type ArticleCategory struct {
	Category string            `json:"category"` // 文章分类
	Articles []*ArticleSummary `json:"articles"` // 文章列表
}
