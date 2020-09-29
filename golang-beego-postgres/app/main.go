package main

import (
	_ "github.com/lib/pq"

	_ "sketch-tech/golang-beego-postgres/app/db"
	_ "sketch-tech/golang-beego-postgres/app/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
