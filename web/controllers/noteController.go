package controllers

import (
	"github.com/kataras/iris/hero"
	"noteWeb/services"
)

func InsertNewNote(noteSer services.NoteService, symbol string, content string) hero.Result {
	if symbol == "" || content == "" {
		return nil
	}

	noteSer.InsertNote(symbol, content)
	return hero.View{
		Name: "welcome.html",
		Data: map[string]interface{}{
			"symbol":  symbol,
			"content": content,
		},
	}
}
