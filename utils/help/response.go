package help

import (
	"github.com/ZRothschild/goIris/config/conf"
)

// 响应错误
func RspErr(msg string) (res conf.Response) {
	res.Status = conf.Fail
	res.Msg = msg
	res.Data = nil
	return
}

// 响应成功数据
func RspSuccess(data interface{}) (res conf.Response) {
	res.Status = conf.Success
	res.Msg = "success"
	res.Data = data
	return
}

// 错误具体返回
func Err(msg string, data interface{}) (res conf.Response) {
	res.Status = conf.Fail
	res.Msg = msg
	res.Data = data
	return
}

// 响应成功具体返回
func Success(msg string, data interface{}) (res conf.Response) {
	res.Status = conf.Success
	res.Msg = msg
	res.Data = data
	return
}