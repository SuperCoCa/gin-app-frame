package v1

import (
	"github.com/gin-app-frame/internal/app/schema"
	"github.com/gin-app-frame/internal/pkg/response"
	"github.com/gin-app-frame/internal/pkg/validator"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var params schema.RegisterUser
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ValidateError(ctx, validator.GetErrMessage(params, err))
		return
	}
	response.Success(ctx, gin.H{
		"msg": "注册成功",
	})
}
