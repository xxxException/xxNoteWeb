package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/mvc"
	"xxNoteWeb/services"
)

type NoteController struct {
	Ctx     iris.Context
	Service services.NoteService
}

var indexView = mvc.View{
	Name: "../views/index.html",
}

func (this *NoteController) Get() {
	return indexView
}

func (this *NoteController) PostNew(ctx iris.Context, symbol string) {

}

func (this *NoteController) PostDelete(ctx iris.Context, symbol string) {

}

func (this *NoteController) PostShare(ctx iris.Context, symbol string) {

}

func (this *NoteController) PostEdit(ctx iris.Context, symbol string, content string) {
	if symbol == "" || content == "" {
		return nil
	}
	println("symbol: " + symbol + "   content: " + content)
	noteSer.InsertNote(symbol, content)
	return hero.View{
		Name: "welcome.html",
		Data: map[string]interface{}{
			"symbol":  symbol,
			"content": content,
		},
	}
}

func (this *NoteController) GetQuery(ctx iris.Context, symbol string) mvc.Result {
	//query
	//if nil insert
	//else display
}
