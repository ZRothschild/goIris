package main

import (
	"github.com/ZRothschild/goIris/backend/backendRoute"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	backendRoute.InitRoute(app)
	_ = app.Listen(":8080")
}
