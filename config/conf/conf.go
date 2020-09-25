package conf

/****************** 统一定义常量 *********************************/

const (
	/******************** 前 start  ******************************/

	BackendConfName       = "config"
	BackendConfType       = "yaml"
	BackendConfPathFirst  = "./conf"
	BackendConfPathSecond = "./../conf"

	BackendCasbinConf   = "./conf/rbac.conf"
	BackendCasbinPrefix = ""
	BackendCasbinTable  = "casbin_rule"

	// 权限类型 角色类型
	Ordinary = 1 // 普通
	Group    = 2 // 组

	/******************** 后 end  **********************************/

	// 用户状态
	Normal   uint8 = 1 // 正常用户
	Abnormal uint8 = 2 // 普通用户

	/******************** 前 start  ******************************/

	FrontendConfName       = "config"
	FrontendConfType       = "yaml"
	FrontendConfPathFirst  = "./conf"
	FrontendConfPathSecond = "./../conf"

	FrontendCasbinConf   = "./conf/rbac.conf"
	FrontendCasbinPrefix = ""
	FrontendCasbinTable  = "casbin_rule"

	/******************** 前 end  **********************************/
)

/************************  自定义结构体 返回实现  ******************************/
