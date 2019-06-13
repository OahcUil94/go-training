package main

import (
	"github.com/OahcUil94/go-training/iris-training/overview/web/controllers"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))

	app.Run(
		iris.Addr("localhost:9000"),
	)
}
