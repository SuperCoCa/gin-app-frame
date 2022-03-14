package schema

type User struct {
	Id
	Username string `json:"username" gorm:"not null;comment:用户名"`
	Password string `json:"password" gorm:"not null;comment:密码"`
	Timestamps
	SoftDelete
}

type RegisterUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,user_password,min=8,max=16"`
}

func (registerUser RegisterUser) GetMessage() map[string]string {
	return map[string]string{
		"Username.required":      "用户名不能为空",
		"Password.required":      "密码不能为空",
		"Password.user_password": "密码必须包含数字和字母",
		"Password.min":           "密码长度为8到16位",
		"Password.max":           "密码长度为8到16位",
	}
}
