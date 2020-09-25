package controller

import (
	"github.com/ZRothschild/goIris/backend/web/service"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
)

type Permission struct {
	Ctx           iris.Context
	Trans         *ut.Translator
	PermissionSrv *service.Permission
	Validate      *validator.Validate
}

// 初始化 permission 控制器
func NewPermission(ctx iris.Context, trans *ut.Translator, permissionSrv *service.Permission, validate *validator.Validate) *Permission {
	return &Permission{
		Ctx:           ctx,
		Trans:         trans,
		PermissionSrv: permissionSrv,
		Validate:      validate,
	}
}
