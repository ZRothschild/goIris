package controller

import (
	"github.com/ZRothschild/goIris/backend/web/service"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type Role struct {
	Ctx      iris.Context
	Trans    *ut.Translator
	RoleSrv  *service.Role
	Validate *validator.Validate
}

// 初始化 role 控制器
func NewRole(ctx iris.Context, trans *ut.Translator, roleSrv *service.Role, validate *validator.Validate) *Role {
	return &Role{
		Ctx:      ctx,
		Trans:    trans,
		RoleSrv:  roleSrv,
		Validate: validate,
	}
}
