package main

import (
	_ "github.com/lib/pq"

	_ "sketch-tech/golang-beego-postgres/app/config"
	_ "sketch-tech/golang-beego-postgres/app/routers"

	"github.com/astaxie/beego"

	"sketch-tech/golang-beego-postgres/app/db"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Initialize database
	db.Init()

	beego.Run()
}
