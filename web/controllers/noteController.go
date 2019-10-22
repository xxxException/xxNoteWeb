package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"xxNoteWeb/services"
	"xxNoteWeb/web/middleware/logger"
)

/*
todo: 给note加锁，在有人打开的情况下，只能读，不能写
todo: 日志的添加，暂定logrus 2019.10.22 晚任务
*/

type NoteController struct {
	Ctx iris.Context
	//类型要与app.Register()中注册的类型一致，如果是指针，那都需要为指针
	Service *services.NoteService
	Session sessions.Session
}

var indexView = mvc.View{
	Name: "index",
}

func (this *NoteController) Get() hero.Result {
	println("/note")
	return indexView
}

/*
Post代表请求的类型   Opennote代表监控的path，在route时会自动将首字母改为小写，如果为OpenNote会被解析为open/note
*/
func (this *NoteController) PostOpennote() {
	var symbol = this.Ctx.FormValue("symbol")

	isExist, err := this.Service.IsExistNote(symbol)
	if err != nil {
		logger.Logger.Error("open note fail : " + err.Error())
		_, _ = this.Ctx.JSON(iris.Map{
			"status": "error",
			"msg":    "打开失败",
		})
		return
	}
	//不存在，new
	if !isExist {
		err = this.Service.NewNote(symbol)
		if err != nil {
			logger.Logger.Error("new note fail : " + err.Error())
			_, _ = this.Ctx.JSON(iris.Map{
				"status": "normal",
				"msg":    "",
			})
		}
	} else {
		_, _ = this.Ctx.JSON(iris.Map{
			"status": "error",
			"msg":    "该symbol已被占用",
		})
	}
}

func (this *NoteController) PostDelete() {
	var symbol = this.Ctx.FormValue("symbol")
	err := this.Service.DeleteNote(symbol)
	if err != nil {
		//log
		logger.Logger.Error("delete note fail : " + err.Error())
		_, _ = this.Ctx.JSON(iris.Map{
			"status": "error",
			"msg":    "删除失败",
		})
		return
	}
	_, _ = this.Ctx.JSON(iris.Map{
		"status": "normal",
		"msg":    "",
	})
}

/*
func (this *NoteController) PostShare() {
	var symbol = this.Ctx.FormValue("symbol")
	note, err := this.Service.GetNoteBySymbol(symbol)
	if err != nil {
		//log
		logger.Logger.Error("share note fail : " + err.Error())
		//return web error
		_,_ = this.Ctx.JSON(iris.Map{
			"status": "error",
			"msg": "保存失败",
		})
		return
	}
	_,_ = this.Ctx.JSON(iris.Map{
		"status": "normal",
		"msg": "",
	})
}
*/

func (this *NoteController) PostEdit() {
	var symbol = this.Ctx.FormValue("symbol")
	var content = this.Ctx.FormValue("content")
	err := this.Service.UpdateNote(symbol, content)
	if err != nil {
		//log
		logger.Logger.Error("edit note fail : " + err.Error())
		_, _ = this.Ctx.JSON(iris.Map{
			"status": "error",
			"msg":    "保存失败",
		})
		//return web error
		return
	}
	_, _ = this.Ctx.JSON(iris.Map{
		"status": "normal",
		"msg":    "",
	})
	return
}

func (this *NoteController) GetQuery() {
	symbol := this.Ctx.URLParam("symbol")
	note, err := this.Service.GetNoteBySymbol(symbol)
	if err != nil {
		logger.Logger.Error("query note fail : " + err.Error())
		_, _ = this.Ctx.JSON(iris.Map{
			"status": "error",
			"msg":    "查询失败",
			"note": iris.Map{
				"symbol":     note.Symbol,
				"content":    note.Content,
				"editTime":   note.EditTime,
				"createTime": note.CreateTime,
			},
		})
		//vue
	}
	_, _ = this.Ctx.JSON(iris.Map{
		"status": "normal",
		"msg":    "",
		"note": iris.Map{
			"symbol":     note.Symbol,
			"content":    note.Content,
			"editTime":   note.EditTime,
			"createTime": note.CreateTime,
		},
	})
	return
}
