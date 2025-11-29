package model

// UserProfile 用户的详细信息
type UserProfile struct {
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像
	Email    string `json:"email"`    // 邮箱
	Password string `json:"password"` // 密码
}
