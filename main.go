package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"xxNoteWeb/bootstrap"
	"xxNoteWeb/repositories"
	"xxNoteWeb/routes"
	"xxNoteWeb/services"
	"xxNoteWeb/web/controllers"
)

func NewApp() *bootstrapper.Bootstrapper {
	app := bootstrap.New("xxNoteWeb", "xzw")

	// route config , controller etc ..
	app.Config()
	//others  like:log, corver  etc..
	app.Bootstrap()
	//view
	app.SetupViews("xxNoteWeb/web/views")
}

func main() {
	app := NewApp()
	app.Listen(":80")
}
