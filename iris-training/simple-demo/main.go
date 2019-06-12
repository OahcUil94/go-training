package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main() {
	app := iris.Default()

	// Post, Put, Delete, Patch, Head, Options
	app.Get("/ping", func (ctx iris.Context) {
		// content的长度
		length, err := ctx.JSON(iris.Map{
			"message": "pong",
		})
		fmt.Println(length)
		fmt.Println(err)
	})

	app.Get("/users/{id:uint64}", func (ctx iris.Context) {
		id := ctx.Params().GetUint64Default("id", 0)
		if _, err := ctx.JSON(iris.Map{
			"user_id": id,
		}); err != nil {
			fmt.Println("err: ", err.Error())
		}
	})

	if err := app.Run(iris.Addr(":8080")); err != nil {
		fmt.Println("error: ", err.Error())
	}
}

// 2017年Go web框架的对比: https://blog.csdn.net/dev_csdn/article/details/78740990