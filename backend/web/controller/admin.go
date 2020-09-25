package controller

import (
	"github.com/ZRothschild/goIris/backend/web/service"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

type Admin struct {
	Ctx      iris.Context
	Trans    *ut.Translator
	AdminSrv *service.Admin
	jwt      *jwt.Middleware
	Validate *validator.Validate
}

// 初始化 admin 控制器
func NewAdmin(ctx iris.Context, adminSrv *service.Admin, trans *ut.Translator, validate *validator.Validate) *Admin {
	return &Admin{
		Ctx:      ctx,
		Trans:    trans,
		AdminSrv: adminSrv,
		Validate: validate,
	}
}
