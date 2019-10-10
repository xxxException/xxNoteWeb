package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"noteWeb/repositories"
	"noteWeb/services"
	"noteWeb/web/controllers"
)

type Note struct {
	Info       string
	Identifier string
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./web/views", ".html").Reload(true))
	//app.OnErrorCode(iris.StatusInternalServerError, )

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		// ctx.Values() 是一个很有用的东西，主要用来使 处理方法与中间件 通信 记住真的很重要
		// ctx.Values().GetString("error") 获取自定义错误提示信息
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}
		ctx.Writef("(Unexpected) internal server error")
	})

	app.Use(func(ctx iris.Context) {
		println("begin request for path : ", ctx.Path())
		//ctx.Application().Loger().Info("begin request for path : %s", ctx.Path())
		ctx.Next()
	})

	app.Done(func(ctx iris.Context) {})

	//hero
	hero.Register(services.NoteService{})

	//app.Get("/welcome", welcomeHandler)
	app.Get("/newNote", hero.Handler(controllers.InsertNewNote))

	//init database
	repositories.DbInit()
	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}
