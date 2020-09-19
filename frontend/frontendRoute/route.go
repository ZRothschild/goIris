package frontendRoute

import (
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/app/repository"
	"github.com/ZRothschild/goIris/config/conf"
	"github.com/ZRothschild/goIris/config/db"
	"github.com/ZRothschild/goIris/config/logger"
	"github.com/ZRothschild/goIris/config/validators"
	"github.com/ZRothschild/goIris/config/viper"
	"github.com/ZRothschild/goIris/frontend/web/controller"
	"github.com/ZRothschild/goIris/frontend/web/service"
	"github.com/ZRothschild/goIris/utils/lib/viperKey"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	viper2 "github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

var (
	viperString string
	dB          *gorm.DB
	newViper    *viper2.Viper
	trans       ut.Translator
	logSrv      *logger.Logger
	sessManager *sessions.Sessions
	validate    *validator.Validate
)

func init() {
	// session 管理
	sessManager = sessions.New(sessions.Config{
		Cookie:  "sessionId",
		Expires: 24 * time.Hour,
	})

	// viper 设置
	newViper = viper.NewViper(
		conf.FrontendConfName,
		conf.FrontendConfType,
		conf.FrontendConfPathFirst,
		conf.FrontendConfPathSecond,
	)

	// 数据库
	viperString, _ = viperKey.MySql("Frontend", newViper)
	dB, _ = db.NewMySql(viperString, newViper)

	// 获取日志
	viperString, _ = viperKey.Log("Service", newViper)
	logSrv, _ = logger.NewLog(viperString, newViper)

	// 验证工具
	viperString, _ = viperKey.Validator("Default", newViper)
	validate, trans, _ = validators.NewValidators(viperString, newViper)

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
	application.Register(trans, userSrv, sessManager, validate)

	// 控制器载入
	application.Handle(new(controller.User))
}
