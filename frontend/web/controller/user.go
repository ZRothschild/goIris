package controller

import (
	"github.com/ZRothschild/goIris/app/model"
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

func (c *User) PostCreate() {
	var (
		err          error
		rowsAffected int64
		user         model.User
	)
	if err = c.Ctx.ReadJSON(&user); err != nil {
		_, _ = c.Ctx.JSON(user)
		return
	}

	for i := 0; i < 10; i++ {
		go func(user model.User, i int) {
			user.ID += uint64(i)
			rowsAffected, err = c.UserSrv.Create(&user)
			if err != nil {
				_, _ = c.Ctx.JSON(rowsAffected)
				return
			}
		}(user, i)
	}
	_, _ = c.Ctx.JSON(user)
	return
}
