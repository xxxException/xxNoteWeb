package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"noteWeb/services"
)

func New(ctx iris.Context, symbol string) {

}

func Delete(ctx iris.Context, symbol string) {

}

func Share(ctx iris.Context, symbol string) {

}

func Edit(ctx iris.Context, symbol string, content string) {
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

func Query(ctx iris.Context, symbol string) mvc.Result {
	//query
	//if nil insert
	//else display
}

func Index() hero.Result {
	return hero.View{
		Name: "note.html",
	}
}
