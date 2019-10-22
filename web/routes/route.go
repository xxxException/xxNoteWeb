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

	note := index.Party("/note")
	//注册service
	//注册的类型要与controller中的service字段类型一致
	note.Register(noteService)
	note.Handle(new(controllers.NoteController))

	//b.Use(func(ctx iris.Context) {
	//	print(ctx.Request().URL.String())
	//	ctx.Next()
	//})

}
