package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	htmlEngine := iris.HTML("./", ".html")
	app.RegisterView(htmlEngine)

	app.Get("/", func (ctx iris.Context) {
		ctx.WriteString("Hello World ! --- from iris.")
	})

	app.Get("hello", func (ctx iris.Context) {
		ctx.ViewData("Title", "测试页面")
		ctx.ViewData("Content", "Hello world! --- from iris.")
		ctx.View("hello.html")
	})

	app.Run(
		iris.Addr(":8081"),
		iris.WithCharset("UTF-8"),
	)
}