package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"xxNoteWeb/bootstrapper"
	"xxNoteWeb/repositories"
	"xxNoteWeb/services"
	"xxNoteWeb/web/controllers"
	"xxNoteWeb/web/routes"
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
