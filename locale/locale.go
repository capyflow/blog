package locale

var V = "{\"code_for_email_not_match.en-us\":\"email not match\",\"code_for_email_not_match.zh-cn\":\"邮箱错误\",\"code_for_password_not_match.en-us\":\"password not match\",\"code_for_password_not_match.zh-cn\":\"密码错误\"}"

var K = struct {
	CODE_FOR_PASSWORD_NOT_MATCH string
	CODE_FOR_EMAIL_NOT_MATCH string
} {
	CODE_FOR_PASSWORD_NOT_MATCH: "code_for_password_not_match",
	CODE_FOR_EMAIL_NOT_MATCH: "code_for_email_not_match",
}
