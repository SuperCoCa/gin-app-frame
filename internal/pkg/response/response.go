package response

import (
	"net/http"

	e "github.com/gin-app-frame/internal/pkg/errors"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"` // 错误码
	Msg  string      `json:"msg"`  // 返回信息
	Data interface{} `json:"data"` // 返回数据
}

// @description 返回请求成功信息
// @param ctx *gin.Context gin上下文
// @param data interface{} 返回数据
func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// @description 返回请求失败信息
// @param ctx *gin.Context gin上下文
// @param code int 错误码
// @param message string 错误信息
func Error(ctx *gin.Context, code int, message string, httpStatus int) {
	ctx.JSON(httpStatus, Response{
		Code: code,
		Msg:  message,
		Data: nil,
	})
}

// @description api请求错误
// @param ctx *gin.Context gin上下文
// @param code int 错误码
func ApiError(ctx *gin.Context, code int) {
	info, ok := e.Errors[code]
	if !ok {
		info = e.Errors[e.ErrServer]
	}
	Error(ctx, info.Code, info.Msg, http.StatusOK)
}

// @description 参数错误
// @param ctx *gin.Context gin上下文
// @param code int 错误信息
func ValidateError(ctx *gin.Context, message string) {
	info := e.Errors[e.ErrParams]
	if message == "" {
		info = e.Errors[e.ErrServer]
	} else {
		info.Msg = message
	}

	Error(ctx, info.Code, info.Msg, http.StatusOK)
}
