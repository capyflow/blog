package locale

var V = "{\"code_for_article_category_not_exist.en-us\":\"article category not exist\",\"code_for_article_category_not_exist.zh-cn\":\"文章分类不存在\",\"code_for_article_not_exist.en-us\":\"article not exist\",\"code_for_article_not_exist.zh-cn\":\"文章不存在\",\"code_for_email_not_match.en-us\":\"email not match\",\"code_for_email_not_match.zh-cn\":\"邮箱错误\",\"code_for_password_not_match.en-us\":\"password not match\",\"code_for_password_not_match.zh-cn\":\"密码错误\"}"

var K = struct {
	CODE_FOR_PASSWORD_NOT_MATCH string
	CODE_FOR_EMAIL_NOT_MATCH string
	CODE_FOR_ARTICLE_NOT_EXIST string
	CODE_FOR_ARTICLE_CATEGORY_NOT_EXIST string
} {
	CODE_FOR_PASSWORD_NOT_MATCH: "code_for_password_not_match",
	CODE_FOR_EMAIL_NOT_MATCH: "code_for_email_not_match",
	CODE_FOR_ARTICLE_NOT_EXIST: "code_for_article_not_exist",
	CODE_FOR_ARTICLE_CATEGORY_NOT_EXIST: "code_for_article_category_not_exist",
}
