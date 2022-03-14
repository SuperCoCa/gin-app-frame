package e

const (
	ErrDatabase int = iota + 100101
)

// @description 注册数据库错误信息
func registerDatabaseErrorInfo() {
	register(ErrDatabase, "数据库错误")
}
