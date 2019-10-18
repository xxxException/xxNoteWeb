package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"xxNoteWeb/services"
)

/*
todo:给note加锁，在有人打开的情况下，只能读，不能写
*/

type NoteController struct {
	Ctx     iris.Context
	Service services.NoteService
	Session sessions.Session
}

var indexView = mvc.View{
	Name: "index",
}

func (this *NoteController) Get() hero.Result {
	println("/note")
	return indexView
}

func (this *NoteController) PostOpennote() {
	println("open note")
	var symbol = this.Ctx.FormValue("symbol")
	isExist, err := this.Service.IsExistNote(symbol)
	if err != nil {
		//todo:log and html response
		println("count note fail: ", err)
		return
	}
	//存在
	if isExist {
		err = this.Service.NewNote(symbol)
		if err != nil {
			//todo:log and html err response
			println("new note file: ", err)
			return
		}
		//todo: 发信息给前端，可以开始写入，同时将symbol保存在前端
		return
	}

}

func (this *NoteController) PostDelete(symbol string) {

}

func (this *NoteController) PostShare(symbol string) {

}

func (this *NoteController) PostEdit(symbol string, content string) {

}

func (this *NoteController) GetQuery() mvc.Result {
	symbol := this.Ctx.URLParam("symbol")
	println(symbol)
	//query
	//if nil insert
	//else display
	return indexView
}
