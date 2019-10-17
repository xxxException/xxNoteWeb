package bootstrapper

import (
	//"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"log"
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
	return b
}

func (this *Bootstrapper) Bootstrap(configList ...Configurator) *Bootstrapper {
	//this.SetupErrorHandler()
	this.Favicon("./web/static/favicon.ico")
	//this.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	//this.setupCron()
	this.Use(recover.New())
	this.Use(logger.New())
	return this
}

func (this *Bootstrapper) Configure(configList ...Configurator) *Bootstrapper {
	for _, config := range configList {
		config(this)
	}

	return this
}

func (this *Bootstrapper) SetupViews(viewDir string) {
	//htmlEngine := iris.HTML(viewDir, ".html").Layout("shared/layout.html")
	htmlEngine := iris.HTML(viewDir, ".html")

	// production 环境设置 false
	htmlEngine.Reload(true)

	/*
		htmlEngine.AddFunc("FromUnixTimeShort", func(t int) string {
			dt := time.Unix(int64(t), int64(0))
			return dt.Format(conf.SysTimeFormShort)
		})

		htmlEngine.AddFunc("FromUnixTime", func(t int) string {
			dt := time.Unix(int64(t), int64(0))
			return dt.Format(conf.SysTimeForm)
		})
	*/

	this.RegisterView(htmlEngine)
}

func (this *Bootstrapper) Listen(addr string) {
	err := this.Run(iris.Addr(addr))
	if err != nil {
		log.Fatal("bootstrap.Listen error ", err)
	}
}
