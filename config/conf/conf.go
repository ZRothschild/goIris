package conf

/****************** 统一定义常量 *********************************/

const (
	Success          = 0
	Fail             = 1
	Glob             = "Global"
	Frontend         = "Frontend"
	Backend          = "Backend"
	FrontendConfName = "config"
	FrontendConfType = "yaml"
	FrontendConfPath = "../conf"
)

/************************  自定义结构体 返回实现去 utils/tool  ******************************/

type (
	// 统一返回结构体
	Response struct {
		Status int64       `json:"status"` // 0 代表成功
		Msg    string      `json:"msg"`    // 返回提示
		Data   interface{} `json:"data"`   // 内容
	}

	// 页码结构体
	Pagination struct {
		Page      int   `json:"page" example:"0"`      // 当前页
		PageSize  int   `json:"pageSize" example:"20"` // 每页条数
		TotalPage int   `json:"totalPage"`             // 总页数
		Total     int64 `json:"total"`                 // 总条数
	}

	// 查询条件
	Where struct {
		Type  string        // where group having select order
		Key   interface{}   // 表达式
		Value []interface{} // 值
	}

	// config 配置文件 Global 结构体
	Global struct {
		App      string
		Env      string
		Backend  string
		Frontend string
	}
)
