package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// @description 用户密码校验
func ValidateUserPassword(fieldLevel validator.FieldLevel) bool {
	password := fieldLevel.Field().String()
	// 必须包含数字和英文字母，且长度为8到16位
	hasNumber, _ := regexp.MatchString(`[0-9]+`, password)
	hasLetter, _ := regexp.MatchString(`[a-zA-Z]+`, password)
	return hasNumber && hasLetter
}
