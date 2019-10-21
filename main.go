package main

import (
	"xxNoteWeb/bootstrapper"
	"xxNoteWeb/dao"
	"xxNoteWeb/dataSource"
	"xxNoteWeb/web/routes"
)

func NewApp() *bootstrapper.Bootstrapper {
	dataSource.NewMysqlEngineGroup()
	dao.DbInit()
	app := bootstrapper.New("xxNoteWeb", "xzw")

	// route config , controller etc ..
	app.Configure(routes.Configure)
	//others  like:log, recover  etc..
	app.Bootstrap()
	//view
	app.SetupViews("web/views")

	return app
}

func main() {
	app := NewApp()
	app.Listen(":80")
}
