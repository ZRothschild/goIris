package controller

import (
	"github.com/ZRothschild/goIris/frontend/web/param/frontendReq"
	"github.com/ZRothschild/goIris/frontend/web/service"
	"github.com/ZRothschild/goIris/utils/help/response"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type User struct {
	Ctx      iris.Context
	Trans    *ut.Translator
	UserSrv  *service.User
	Session  *sessions.Session
	Validate *validator.Validate
}

// 初始化 user 控制器
func NewUser(ctx iris.Context, trans *ut.Translator, userSrv *service.User, session *sessions.Session, validate *validator.Validate) *User {
	return &User{
		Ctx:      ctx,
		Trans:    trans,
		UserSrv:  userSrv,
		Session:  session,
		Validate: validate,
	}
}

// 用户注册
// @Summary 用户注册
// @tags 用户模块
// @Description 用户注册
// @ID usersPostRegister
// @Accept  json
// @Produce  json
// @Param 请求参数 json body frontendReq.UserRegister true "请求参数json"
// @Success 200 {object} response.Response
// @Router /users/register [POST]
func (c *User) PostRegister() {
	var (
		err          error
		rowsAffected int64
		req          frontendReq.UserRegister
	)

	if err = c.Ctx.ReadJSON(&req); err != nil {
		response.CtxErr(c.Ctx, "请求错误", err)
		return
	}

	if err := c.Validate.Struct(req); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			response.CtxErr(c.Ctx, "请求错误", e.Translate(*c.Trans))
			return
		}
	}

	response.CtxSuccess(c.Ctx, "请求成功", nil)
	return
	rowsAffected, err = c.UserSrv.Register(&req)
	response.CtxSuccess(c.Ctx, "请求成功", rowsAffected)
	// for i := 0; i < 10; i++ {
	// 	go func(user model.User) {
	// 		rowsAffected, err = c.UserSrv.Create(&user)
	// 		if err != nil {
	// 			_, _ = c.Ctx.JSON(rowsAffected)
	// 			return
	// 		}
	// 	}(user)
	// }
	return
}

// 用户登录
func (c *User) PostLogin() {

}
