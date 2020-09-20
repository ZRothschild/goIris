package main

import (
	_ "github.com/ZRothschild/goIris/frontend/docs"
	"github.com/ZRothschild/goIris/frontend/frontendRoute"
	"github.com/kataras/iris/v12"
)

func main() {

	// e := casbin.NewEnforcer("E:\\go-test\\test\\abac\\abac_model.conf")
	// e.Enforce("", "", "")

	app := iris.New()
	frontendRoute.InitRoute(app)
	_ = app.Listen(":8080")
}
