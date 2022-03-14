package e

import (
	"github.com/gin-app-frame/internal/pkg/global"
)

type ErrorInfo struct {
	Code int    // 错误码
	Msg  string // 错误信息
}

var Errors = map[int]ErrorInfo{}

// @description 注册错误码
// @param code int 错误码
// @param message string 错误信息
func register(code int, message string) {
	info := &ErrorInfo{
		Code: code,
		Msg:  message,
	}
	if _, ok := Errors[code]; ok {
		global.App.Logger.Panic("错误码 %d 已被使用", code)
	}
	Errors[code] = *info
}

// @description 注册所有错误信息
func RegisterErrorInfo() {
	// 基础错误
	registerBaseErrorInfo()
	// 数据库错误
	registerDatabaseErrorInfo()
	// 认证授权错误
	registerAuthErrorInfo()
}
