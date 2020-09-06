package response

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
	res.Status = Fail
	res.Msg = msg
	res.Data = nil
	return
}

// 响应成功数据
func RspSuccess(data interface{}) (res Response) {
	res.Status = Suc
	res.Msg = "success"
	res.Data = data
	return
}

// 错误具体返回
func Err(msg string, data interface{}) (res Response) {
	res.Status = Fail
	res.Msg = msg
	res.Data = data
	return
}

// 响应成功具体返回
func Success(msg string, data interface{}) (res Response) {
	res.Status = Suc
	res.Msg = msg
	res.Data = data
	return
}
