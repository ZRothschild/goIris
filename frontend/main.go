package main

import (
	"github.com/ZRothschild/goIris/frontend/frontendRoute"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	frontendRoute.InitRoute(app)
	_ = app.Listen(":8080")
}
