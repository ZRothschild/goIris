package frontendRoute

import (
	"fmt"
	"github.com/ZRothschild/goIris/app/model"
	"github.com/ZRothschild/goIris/app/repository"
	"github.com/ZRothschild/goIris/config/conf"
	"github.com/ZRothschild/goIris/config/db"
	"github.com/ZRothschild/goIris/config/logger"
	"github.com/ZRothschild/goIris/config/validators"
	"github.com/ZRothschild/goIris/config/viper"
	"github.com/ZRothschild/goIris/frontend/web/controller"
	"github.com/ZRothschild/goIris/frontend/web/param/frontendReq"
	"github.com/ZRothschild/goIris/frontend/web/service"
	"github.com/ZRothschild/goIris/utils/lib/viperKey"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	viper2 "github.com/spf13/viper"
	"gorm.io/gorm"
	"net/http"
	"testing"
	"time"
)

var (
	viperStringTest string
	dBTest          *gorm.DB
	newViperTest    *viper2.Viper
	transTest       ut.Translator
	logSrvTest      *logger.Logger
	sessManagerTest *sessions.Sessions
	validateTest    *validator.Validate
)

func init() {
	// session 管理
	sessManagerTest = sessions.New(sessions.Config{
		Cookie:  "sessionId",
		Expires: 24 * time.Hour,
	})

	// viper 设置
	newViperTest = viper.NewViper(
		conf.FrontendConfName,
		conf.FrontendConfType,
		conf.FrontendConfPathFirst,
		conf.FrontendConfPathSecond,
	)

	// 数据库
	viperStringTest, _ = viperKey.MySql("Frontend", newViperTest)
	dBTest, _ = db.NewMySql(viperStringTest, newViperTest)

	// 获取日志
	viperStringTest, _ = viperKey.Log("Service", newViperTest)
	logSrvTest, _ = logger.NewLog(viperStringTest, newViperTest)

	// 验证工具
	viperStringTest, _ = viperKey.Validator("Default", newViperTest)
	validateTest, transTest, _ = validators.NewValidators(viperStringTest, newViperTest)
}

func TestUser_PostRegister(t *testing.T) {
	app := NewApp()
	e := httptest.New(t, app)
	req := e.Request(http.MethodPost, "/users/register")
	reqParam := frontendReq.UserRegister{
		Nickname: "",
		Email:    "873908900.com",
		Password: "124577",
	}
	// body := `{"key": "value"}`
	resp := req.WithJSON(reqParam).Expect().Status(http.StatusOK)
	// assert.Equal(t, body, resp.Body().Raw())
	fmt.Printf("返回格式  %s\n", resp.Body().Raw())
}

func NewApp() *iris.Application {
	app := iris.New()
	testInitRout(app)
	return app
}

// 初始化 路由
func testInitRout(app *iris.Application) {
	mvc.Configure(app.Party("/users"), testUsers)
}

// user controller
func testUsers(application *mvc.Application) {
	// user  依赖
	userModel := model.NewUser()

	// 数据仓库
	userRepository := repository.NewUser(userModel, dBTest)

	// 数据服务
	userSrv := service.NewUser(userRepository, logSrvTest)

	// 依赖注入
	application.Register(transTest, userSrv, sessManagerTest, validateTest)

	// 控制器载入
	application.Handle(new(controller.User))
}
