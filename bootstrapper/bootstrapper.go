package bootstrapper

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	AppName  string
	AppOwner string
}

func New(appName string, appOwner string) *Bootstrapper {
	b := &Bootstrapper{
		Application: iris.New(),
		AppName:     appName,
		AppOwner:    appOwner,
	}

}

func (this *Bootstrapper) Bootstrap(configList ...Configurator) *Bootstrapper {
	this.SetupErrorHandler()
	this.Favicon("../favicon.ico")
	//this.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	this.setupCron()
	this.Use(recover.New())
	this.Use(logger.New())
	return this
}

func (this *Bootstrapper) Configure(configList ...Configurator) *Bootstrapper {
	for _, config := range configList {
		config(b)
	}

	return this
}

func (this *Bootstrapper) SetupViews(viewDir string) {
	htmlEngine := iris.HTML(viewDir, ".html").Layout("shared/layout.html")
	//htmlEngine := iris.HTML(viewDir, ".html")

	// production 环境设置 false
	htmlEngine.Reload(true)

	htmlEngine.AddFunc("FromUnixTimeShort", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeFormShort)
	})

	htmlEngine.AddFunc("FromUnixTime", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeForm)
	})

	this.RegisterView(htmlEngine)
}
