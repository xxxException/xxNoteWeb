package routes

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"xxNoteWeb/bootstrapper"
	"xxNoteWeb/web/controllers"
)

func Configure(b *bootstrapper.Bootstrapper) {
	//service
	noteService := services.NewNoteService()

	index := mvc.New(b.Party("/"))

	//注册service
	index.Register(noteService)
}
