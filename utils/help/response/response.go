package response

import (
	"github.com/kataras/iris/v12"
)

const (
	Suc  = 0
	Fail = 1
)

// 统一返回结构体
type Response struct {
	Status int64       `json:"status"` // 0 代表成功
	Msg    string      `json:"msg"`    // 返回提示
	Data   interface{} `json:"data"`   // 内容
}

// 响应错误
func RspErr(msg string) (res Response) {
	res = Response{
		Status: Fail,
		Msg:    msg,
		Data:   nil,
	}
	return
}

// 响应成功数据
func RspSuccess(data interface{}) (res Response) {
	res = Response{
		Status: Suc,
		Msg:    "success",
		Data:   data,
	}
	return
}

// 错误具体返回
func Err(msg string, data interface{}) (res Response) {
	res = Response{
		Status: Fail,
		Msg:    msg,
		Data:   data,
	}
	return
}

// 响应成功具体返回
func Success(msg string, data interface{}) (res Response) {
	res = Response{
		Status: Suc,
		Msg:    msg,
		Data:   data,
	}
	return
}

// 错误具体返回
func CtxErr(ctx iris.Context, msg string, data interface{}) {
	res := Response{
		Status: Fail,
		Msg:    msg,
		Data:   data,
	}
	_, _ = ctx.JSON(res)
	return
}

// 响应成功具体返回
func CtxSuccess(ctx iris.Context, msg string, data interface{}) {
	res := Response{
		Status: Suc,
		Msg:    msg,
		Data:   data,
	}
	_, _ = ctx.JSON(res)
	return
}
