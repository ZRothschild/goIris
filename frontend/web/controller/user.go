package controller

import (
	"github.com/ZRothschild/goIris/frontend/web/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type User struct {
	Ctx     iris.Context
	UserSrv *service.User
	Session *sessions.Session
}

// 初始化 user 控制器
func NewUser(ctx iris.Context, userSrv *service.User, session *sessions.Session) *User {
	return &User{Ctx: ctx, UserSrv: userSrv, Session: session}
}

func (c *User) GetId() {
	_, _ = c.Ctx.JSON("成功")
}
