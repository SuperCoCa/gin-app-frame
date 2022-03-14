package e

const (
	ErrTokenInvalid int = iota + 100201
	ErrTokenExpire
	ErrTokenMissing
)

// @description 注册认证授权错误信息
func registerAuthErrorInfo() {
	register(ErrTokenInvalid, "Token无效")
	register(ErrTokenExpire, "Token已过期")
	register(ErrTokenMissing, "缺少Token")
}
