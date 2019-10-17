package main

import (
	"xxNoteWeb/bootstrapper"
)

func NewApp() *bootstrapper.Bootstrapper {
	app := bootstrapper.New("xxNoteWeb", "xzw")

	// route config , controller etc ..
	app.Configure()
	//others  like:log, corver  etc..
	app.Bootstrap()
	//view
	app.SetupViews("xxNoteWeb/web/views")

	return app
}

func main() {
	app := NewApp()
	app.Application.Listen(":80")
}
