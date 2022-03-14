package validator

import (
	"github.com/gin-app-frame/internal/pkg/global"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	GetMessage() map[string]string
}

func GetErrMessage(request interface{}, err error) string {
	if _, isValidationErrors := err.(validator.ValidationErrors); isValidationErrors {
		// 判断是否为Validator接口
		_, isValidator := request.(Validator)
		for _, v := range err.(validator.ValidationErrors) {
			if isValidator {
				if msg, ok := request.(Validator).GetMessage()[v.Field()+"."+v.Tag()]; ok {
					return msg
				}
				global.App.Logger.Error(v.Error())
				return ""
			}
		}
	}

	return "参数错误"
}

// @description 初始化自定义验证规则
func InitValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("user_password", ValidateUserPassword)
	}
}
