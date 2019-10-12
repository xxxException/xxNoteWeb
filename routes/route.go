package routes

import (
	"github.com/kataras/iris"
	"noteWeb/web/noteController"
)

func Route(app *iris.Application) {
	noteRouter := app.Party("/note", noteController.Index)
	noteRouter.Get("/query", noteController.Query)
	noteRouter.Post("/edit", noteController.Edit)
	noteRouter.Post("/new", noteController.New)
	noteRouter.Post("/delete", noteController.Delete)
	noteRouter.Post("/share", noteController.Share)
}
