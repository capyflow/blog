package pkg

import (
	"github.com/capyflow/blog/locale"
	"github.com/capyflow/vortex/v2"
)

// SubCodes 子错误码
var SubCodes = struct {
	PasswordNotMatch vortex.SubCode
	EmailNotMatch    vortex.SubCode

	ArticleNotExist         vortex.SubCode
	ArticleCategoryNotExist vortex.SubCode
}{
	PasswordNotMatch:        vortex.SubCode{SubCode: 10403, I18nKey: locale.K.CODE_FOR_PASSWORD_NOT_MATCH},
	EmailNotMatch:           vortex.SubCode{SubCode: 10404, I18nKey: locale.K.CODE_FOR_EMAIL_NOT_MATCH},
	ArticleNotExist:         vortex.SubCode{SubCode: 20404, I18nKey: locale.K.CODE_FOR_ARTICLE_NOT_EXIST},
	ArticleCategoryNotExist: vortex.SubCode{SubCode: 20405, I18nKey: locale.K.CODE_FOR_ARTICLE_CATEGORY_NOT_EXIST},
}
