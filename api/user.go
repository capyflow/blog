package api

import "github.com/capyflow/blog/model"

// API 参数，所有的接口接收的参数和返回的参数都放在这里

// LoginReq 登录请求
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResp 登录响应
type LoginResp struct {
	Token       string             `json:"token"`
	UserProfile *model.UserProfile `json:"user_profile"`
}

// 更新用户信息
type UpdateUserProfileReq struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type UpdateUserProfileResp struct {
	UserProfile *model.UserProfile `json:"user_profile"`
}
