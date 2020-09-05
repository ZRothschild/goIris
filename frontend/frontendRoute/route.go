package frontendRoute

import (
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/app/repository"
	"github.com/ZRothschild/goIris/config/conf"
	"github.com/ZRothschild/goIris/config/db"
	"github.com/ZRothschild/goIris/config/logger"
	"github.com/ZRothschild/goIris/config/viper"
	"github.com/ZRothschild/goIris/frontend/web/controller"
	"github.com/ZRothschild/goIris/frontend/web/service"
	"github.com/ZRothschild/goIris/utils/help"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	viper2 "github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

var (
	sessManager *sessions.Sessions
	dB *gorm.DB
	newViper *viper2.Viper
	logSrv *logger.Logger
)

func init() {
	// session 管理
	sessManager = sessions.New(sessions.Config{
		Cookie:  "sessionId",
		Expires: 24 * time.Hour,
	})

	// viper 设置
	newViper = viper.NewViper(conf.FrontendConfName, conf.FrontendConfType, conf.FrontendConfPath)

	// 数据库
	frontMySqlViperKey, _ := help.FrontMySqlViperKey("Frontend", newViper)
	dB = db.NewMySql(frontMySqlViperKey, newViper)

	//
	frontLogViperKey, _ := help.FrontLogViperKey("Service", newViper)

	// 获取日志
	logSrv, _ = logger.NewLog(frontLogViperKey, newViper)
}

// 初始化 路由
func InitRoute(app *iris.Application) {
	mvc.Configure(app.Party("/users"), users)
}

// user controller
func users(application *mvc.Application) {
	// user  依赖
	userModel := model.NewUser()

	// 数据仓库
	userRepository := repository.NewUser(userModel, dB)

	// 数据服务
	userSrv := service.NewUser(userRepository, logSrv)

	// 依赖注入
	application.Register(userSrv, sessManager)

	// 控制器载入
	application.Handle(new(controller.User))
}


// 初始化控制器依赖
// func InitDependent() (*sessions.Sessions, *gorm.DB, *viper2.Viper) {
// 	sessManager := sessions.New(sessions.Config{
// 		Cookie:  "sessionId",
// 		Expires: 24 * time.Hour,
// 	})
// 	newViper := viper.NewViper(conf.FrontendConfName, conf.FrontendConfType, conf.FrontendConfPath)
// 	frontMySqlViperKey, _ := help.FrontMySqlViperKey("Frontend", newViper)
// 	dB := db.NewMySql(frontMySqlViperKey, newViper)
// 	return sessManager, dB, newViper
// }
