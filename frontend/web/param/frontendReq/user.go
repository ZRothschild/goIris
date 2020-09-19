package frontendReq

// 用户注册请求参数

type (
	UserRegister struct {
		Nickname string `json:"nickname" from:"nickname" validate:"required" label:"昵称"`
		Email    string `json:"email" from:"email" validate:"required,email" label:"邮箱"`
		Password string `json:"password" from:"password" validate:"required" label:"密码"`
	}
)
