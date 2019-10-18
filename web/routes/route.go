package routes

import (
	"github.com/kataras/iris/mvc"
	"xxNoteWeb/bootstrapper"
	"xxNoteWeb/services"
	"xxNoteWeb/web/controllers"
)

func Configure(b *bootstrapper.Bootstrapper) {
	//service
	noteService := services.NewNoteService()

	index := mvc.New(b.Party("/"))

	//注册service
	index.Register(noteService)

	note := index.Party("/note")
	note.Handle(new(controllers.NoteController))

	//b.Use(func(ctx iris.Context) {
	//	print(ctx.Request().URL.String())
	//	ctx.Next()
	//})

}
