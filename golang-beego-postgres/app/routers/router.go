package routers

import (
	"github.com/astaxie/beego"

	"sketch-tech/golang-beego-postgres/app/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/author", &controllers.AuthorController{}, "post:Store"),
		//beego.NSRouter("/book/:id", &controllers.BookController{}, "get:GetBookByID"),
	)
	beego.AddNamespace(ns)
}
