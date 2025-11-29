package model

// UserProfile 用户的详细信息
type UserProfile struct {
	Nickname string `json:"nickname"`  // 昵称
	Avatar   string `json:"avatar"`    // 头像
	UpdateTs int64  `json:"update_ts"` // 更新时间
}
