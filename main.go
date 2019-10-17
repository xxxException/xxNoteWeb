package main

import (
	"xxNoteWeb/bootstrapper"
	"xxNoteWeb/repositories"
	"xxNoteWeb/web/routes"
)

func NewApp() *bootstrapper.Bootstrapper {
	app := bootstrapper.New("xxNoteWeb", "xzw")

	// route config , controller etc ..
	app.Configure(routes.Configure)
	//others  like:log, corver  etc..
	app.Bootstrap()
	//view
	app.SetupViews("web/views")

	repositories.DbInit()
	return app
}

func main() {
	app := NewApp()
	app.Listen(":80")
}
