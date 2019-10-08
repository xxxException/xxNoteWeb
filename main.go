package main

import "github.com/kataras/iris"

type Note struct {
	Info       string
	Identifier string
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./web/views", ".html").Reload(true))
	//app.OnErrorCode(iris.StatusInternalServerError, )

	app.Use(func(ctx iris.Context) {
		ctx.Application().Loger().Info("begin request for path : %s", ctx.Path())
		ctx.Next()
	})

	app.Done(func(ctx iris.Context) {})

	app.Post("")
}
