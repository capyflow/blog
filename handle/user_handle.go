package handle

import (
	"context"
	"errors"

	"github.com/capyflow/Allspark-go/logx"
	"github.com/capyflow/blog/api"
	"github.com/capyflow/blog/pkg"
	"github.com/capyflow/blog/service"
	"github.com/capyflow/vortex/v2"
	"github.com/labstack/echo/v4"
)

type UserHandle struct {
	ctx      context.Context
	userServ *service.UserService
}

// NewUserHandle 创建用户处理句柄
func NewUserHandle(ctx context.Context, userServ *service.UserService) *UserHandle {
	return &UserHandle{
		ctx:      ctx,
		userServ: userServ,
	}
}

// 根据密码登录接口
func (uh *UserHandle) LoginByPassword(ctx *vortex.Context) error {
	var req api.LoginReq
	if err := ctx.Bind(&req); nil != err {
		logx.Errorf("UserHandle|LoginByPassword|Bind req failed|%v", err)
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild, echo.Map{
			"error": err.Error(),
		})
	}

	token, err := uh.userServ.LoginByPwd(ctx.GetContext(), req.Username, req.Password)
	if nil != err {
		logx.Errorf("UserHandle|LoginByPassword|LoginByPwd failed|%v", err)
		if errors.Is(err, pkg.ErrorsEnum.ErrPasswordNotMatch) {
			return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild.WithSubCode(pkg.SubCodes.PasswordNotMatch), echo.Map{
				"error": err.Error(),
			})
		} else if errors.Is(err, pkg.ErrorsEnum.ErrEmailNotMatch) {
			return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild.WithSubCode(pkg.SubCodes.EmailNotMatch), echo.Map{
				"error": err.Error(),
			})
		}
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.InternalError, echo.Map{
			"error": err.Error(),
		})
	}

	userProfile, err := uh.userServ.QueryUserProfile(ctx.GetContext())
	if nil != err {
		logx.Errorf("UserHandle|LoginByPassword|QueryUserProfile failed|%v", err)
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.InternalError, echo.Map{
			"error": err.Error(),
		})
	}

	return vortex.HttpJsonResponse(ctx, vortex.Statuses.Success, api.LoginResp{
		Token:       token,
		UserProfile: userProfile,
	})
}

// 更新个人信息
func (uh *UserHandle) HandleUpdateUserProfile(ctx *vortex.Context) error {
	var req api.UpdateUserProfileReq
	if err := ctx.Bind(&req); nil != err {
		logx.Errorf("UserHandle|HandleUpdateUserProfile|Bind req failed|%v", err)
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild, echo.Map{
			"error": err.Error(),
		})
	}

	userProfile, err := uh.userServ.UpdateUserProfile(ctx.GetContext(), &req)
	if nil != err {
		logx.Errorf("UserHandle|HandleUpdateUserProfile|UpdateUserProfile failed|%v", err)
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.InternalError, echo.Map{
			"error": err.Error(),
		})
	}
	return vortex.HttpJsonResponse(ctx, vortex.Statuses.Success, api.UpdateUserProfileResp{
		UserProfile: userProfile,
	})
}

// 根据邮箱验证码登录接口
func (uh *UserHandle) LoginByEmailCode(ctx *vortex.Context) error {
	var req api.LoginReq
	if err := ctx.Bind(&req); nil != err {
		logx.Errorf("UserHandle|LoginByEmailCode|Bind req failed|%v", err)
		return vortex.HttpJsonResponse(ctx, vortex.Statuses.ParamsInvaild, echo.Map{
			"error": err.Error(),
		})
	}
	// TODO 实现邮箱验证码登录
	return nil
}
