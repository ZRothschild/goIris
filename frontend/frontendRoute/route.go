package frontendRoute

import (
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/app/repository"
	"github.com/ZRothschild/goIris/config/conf"
	"github.com/ZRothschild/goIris/config/db"
	"github.com/ZRothschild/goIris/config/logger"
	"github.com/ZRothschild/goIris/config/validators"
	"github.com/ZRothschild/goIris/config/viper"
	"github.com/ZRothschild/goIris/frontend/docs"
	"github.com/ZRothschild/goIris/frontend/web/controller"
	"github.com/ZRothschild/goIris/frontend/web/service"
	"github.com/ZRothschild/goIris/utils/lib/viperKey"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	viper2 "github.com/spf13/viper"
	"gorm.io/gorm"
	"time"
)

var (
	err         error
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

	// 获取日志
	viperString, err = viperKey.PreKeyViper(newViper, "Log.Service")
	if err != nil {
		// do
	}

	logSrv, err = logger.NewLog(newViper, viperString)
	if err != nil {
		// do
	}

	// 数据库
	viperString, err = viperKey.PreKeyViper(newViper, "MySql.Frontend")
	if err != nil {
		// do
	}

	dB, err = db.NewMySql(newViper, viperString)
	if err != nil {
		// do
	}

	// 验证工具
	viperString, err = viperKey.PreKeyViper(newViper, "Validator.Lang.Zh")
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
	// 用户控制器
	mvc.Configure(app.Party("/users"), users)

	// API 文档生成
	SwaggerApiDoc(app)
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
	application.Register(&trans, userSrv, sessManager, validate)

	// 控制器载入
	application.Handle(new(controller.User))
}

// swagger 自动生成文档
func SwaggerApiDoc(app *iris.Application) {
	// programmatically set swagger info
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:8080"
	docs.SwaggerInfo.Title = "Go Iris Api"
	docs.SwaggerInfo.Description = "go iris api."
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	config := &swagger.Config{
		URL: "http://localhost:8080/swagger/swagger.json", // The url pointing to API definition
	}
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
}
