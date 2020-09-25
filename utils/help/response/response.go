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
func RspErr(msg string) Response {
	return Response{
		Status: Fail,
		Msg:    msg,
		Data:   nil,
	}

}

// 响应成功数据
func RspSuccess(data interface{}) Response {
	return Response{
		Status: Suc,
		Msg:    "success",
		Data:   data,
	}
}

// 错误具体返回
func Err(msg string, data interface{}) Response {
	return Response{
		Status: Fail,
		Msg:    msg,
		Data:   data,
	}
}

// 响应成功具体返回
func Success(msg string, data interface{}) Response {
	return Response{
		Status: Suc,
		Msg:    msg,
		Data:   data,
	}
}

// 错误具体返回
func CtxErr(ctx iris.Context, msg string, data interface{}) {
	_, _ = ctx.JSON(Response{
		Status: Fail,
		Msg:    msg,
		Data:   data,
	})
	return
}

// 响应成功具体返回
func CtxSuccess(ctx iris.Context, msg string, data interface{}) {
	_, _ = ctx.JSON(Response{
		Status: Suc,
		Msg:    msg,
		Data:   data,
	})
	return
}
