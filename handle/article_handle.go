package handle

import (
	"context"
	"errors"

	"github.com/capyflow/Allspark-go/conv"
	"github.com/capyflow/Allspark-go/logx"
	"github.com/capyflow/blog/api"
	"github.com/capyflow/blog/pkg"
	"github.com/capyflow/blog/service"
	"github.com/capyflow/vortex/v2"
	"github.com/labstack/echo/v4"
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

// 发布文章
func (h *ArticleHandle) HandlePublishArticle(ctx *vortex.Context) error {
	var req api.PublishArticleReq
	if err := ctx.Bind(&req); err != nil {
		logx.Errorf("ArticleHandle|HandlePublishArticle|Error|%v", err)
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild, echo.Map{"msg": err.Error()})
	}

	// 校验一下参数
	if len(req.Title) == 0 || len(req.Content) == 0 || len(req.Category) == 0 {
		logx.Errorf("ArticleHandle|HandlePublishArticle|Error|title, content, category are required|%s", conv.ToJsonWithoutError(req))
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild, echo.Map{"msg": "title, content, category are required"})
	}

	if err := h.articleServ.PublishArticle(h.ctx, req.Title, req.Content, req.Category); err != nil {
		logx.Errorf("ArticleHandle|HandlePublishArticle|Error|%v", err)
		if errors.Is(err, pkg.ErrorsEnum.ErrArticleCategoryNotExist) {
			return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild.WithSubCode(pkg.SubCodes.ArticleCategoryNotExist),
				echo.Map{"msg": "article category not exist"})
		}
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild, echo.Map{"msg": err.Error()})
	}
	return vortex.HttpJsonResponse(ctx, vortex.Statuses.Success, echo.Map{"msg": "publish article success"})
}

// 更新文章
func (h *ArticleHandle) HandleUpdateArticle(ctx *vortex.Context) error {
	var req api.UpdateArticleReq
	if err := ctx.Bind(&req); err != nil {
		logx.Errorf("ArticleHandle|HandleUpdateArticle|Error|%v", err)
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild, echo.Map{"msg": err.Error()})
	}

	if len(req.ID) == 0 {
		logx.Errorf("ArticleHandle|HandleUpdateArticle|Error|id is required|%s", conv.ToJsonWithoutError(req))
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild, echo.Map{"msg": "id is required"})
	}

	article, err := h.articleServ.UpdateArticle(h.ctx, req.ID, req.Title, req.Content)
	if err != nil {
		logx.Errorf("ArticleHandle|HandleUpdateArticle|Error|%v", err)
		if errors.Is(err, pkg.ErrorsEnum.ErrArticleNotExist) {
			return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild.WithSubCode(pkg.SubCodes.ArticleNotExist), echo.Map{"msg": err.Error()})
		} else {
			return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild, echo.Map{"msg": err.Error()})
		}
	}
	return vortex.HttpJsonResponse(ctx, vortex.Statuses.Success, api.UpdateArticleResp{ArticleInfo: article})
}

// 删除文章
func (h *ArticleHandle) HandleDeleteArticle(ctx *vortex.Context) error {
	var req api.DeleteArticleReq
	if err := ctx.Bind(&req); err != nil {
		logx.Errorf("ArticleHandle|HandleDeleteArticle|Error|%v", err)
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild, echo.Map{"msg": err.Error()})
	}

	if len(req.ID) == 0 {
		logx.Errorf("ArticleHandle|HandleDeleteArticle|Error|id is required|%s", conv.ToJsonWithoutError(req))
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild, echo.Map{"msg": "id is required"})
	}

	if err := h.articleServ.DeleteArticle(h.ctx, req.ID); err != nil {
		logx.Errorf("ArticleHandle|HandleDeleteArticle|Error|%v", err)
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild, echo.Map{"msg": err.Error()})
	}
	return vortex.HttpJsonResponse(ctx, vortex.Statuses.Success, echo.Map{"msg": "delete article success"})
}
