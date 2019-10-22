package main

import (
	"xxNoteWeb/bootstrapper"
	"xxNoteWeb/dataSource"
	"xxNoteWeb/web/middleware/logger"
	"xxNoteWeb/web/routes"
)

func NewApp() *bootstrapper.Bootstrapper {
	//mysql
	dataSource.NewMysqlEngineGroup()
	dataSource.GetEngineGroup().Ping()

	//logger
	logger.NewLogger()

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
