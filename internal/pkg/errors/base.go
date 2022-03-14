package e

const (
	ErrServer int = iota + 100001
	ErrParams
)

// @description 注册通用错误信息
func registerBaseErrorInfo() {
	register(ErrServer, "服务器内部错误")
	register(ErrParams, "参数错误")
}
