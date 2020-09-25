package backendRoute

import (
	"github.com/ZRothschild/goIris/app/middleware"
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/app/repository"
	"github.com/ZRothschild/goIris/backend/web/controller"
	"github.com/ZRothschild/goIris/backend/web/service"
	"github.com/ZRothschild/goIris/config/casbiner"
	"github.com/ZRothschild/goIris/config/conf"
	"github.com/ZRothschild/goIris/config/db"
	"github.com/ZRothschild/goIris/config/logger"
	"github.com/ZRothschild/goIris/config/validators"
	"github.com/ZRothschild/goIris/config/viper"
	"github.com/ZRothschild/goIris/utils/lib/viperKey"
	"github.com/casbin/casbin/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	jwtMiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	viper2 "github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	err         error
	viperString string
	dB          *gorm.DB
	newViper    *viper2.Viper
	trans       ut.Translator
	logSrv      *logger.Logger
	e           *casbin.Enforcer
	validate    *validator.Validate
	jwt         *jwtMiddleware.Middleware
)

func init() {
	// viper 设置
	newViper = viper.NewViper(
		conf.BackendConfName,
		conf.BackendConfType,
		conf.BackendConfPathFirst,
		conf.BackendConfPathSecond,
	)

	// 获取日志
	viperString, err = viperKey.PreKeyViper(newViper, "Service")
	if err != nil {
		// do
	}

	logSrv, err = logger.NewLog(newViper, viperString)
	if err != nil {
		// do
	}

	// 数据库
	viperString, err = viperKey.PreKeyViper(newViper, "Backend")
	if err != nil {
		// do
	}

	dB, err = db.NewMySql(newViper, viperString)
	if err != nil {
		// do
	}

	// 用户登录Jwt
	viperString, err = viperKey.PreKeyViper(newViper, "JwtKey")
	if err != nil {
		// do
	}
	jwt = middleware.Jwt(newViper, viperString)

	// 鉴权
	e, err = casbiner.NewCasbin(
		dB,
		conf.BackendCasbinConf,
		conf.BackendCasbinPrefix,
		conf.BackendCasbinTable,
	)
	if err != nil {
		// do
	}

	// 验证工具
	viperString, err = viperKey.PreKeyViper(newViper, "Zh")
	if err != nil {
		// do
	}

	validate, trans, err = validators.NewValidators(newViper, viperString)
	if err != nil {
		// do
	}

}

// 初始化 路由
func InitRoute(app *iris.Application) {
	// 管理员控制器
	mvc.Configure(app.Party("/admin"), admins)

	// 角色控制器
	mvc.Configure(app.Party("/role"), roles)

	// 权限
	mvc.Configure(app.Party("/permission"), permissions)
}

// admin controller
func admins(application *mvc.Application) {
	// 管理员  依赖
	adminModel := model.NewAdmin()

	// 数据仓库
	adminRepository := repository.NewAdmin(adminModel, dB)

	// 数据服务
	adminSrv := service.NewAdmin(adminRepository, logSrv)

	// 依赖注入
	application.Register(&trans, adminSrv, jwt, validate)

	// 控制器载入
	application.Handle(new(controller.Admin))
}

// role controller
func roles(application *mvc.Application) {
	// 角色  依赖
	roleModel := model.NewRole()

	// 数据仓库
	roleRepository := repository.NewRole(roleModel, dB)

	// 数据服务
	roleSrv := service.NewRole(roleRepository, logSrv)

	// 依赖注入
	application.Register(&trans, roleSrv, jwt, validate)

	// 控制器载入
	application.Handle(new(controller.Role))
}

// permission controller
func permissions(application *mvc.Application) {
	// 角色  依赖
	permissionModel := model.NewPermission()

	// 数据仓库
	permissionRepository := repository.NewPermission(permissionModel, dB)

	// 数据服务
	permissionSrv := service.NewPermission(permissionRepository, logSrv)

	// 依赖注入
	application.Register(&trans, permissionSrv, jwt, validate)

	// 控制器载入
	application.Handle(new(controller.Permission))
}
