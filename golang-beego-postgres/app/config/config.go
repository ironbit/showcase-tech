package config

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"sketch-tech/golang-beego-postgres/app/tool"
)

const (
	cPort = 8000
)

func init() {
	host := tool.GetEnv("APP_HOST")
	port, err := strconv.Atoi(tool.GetEnv("APP_PORT"))
	if err != nil {
		logs.Error("APP_PORT not found")
		port = cPort
	}

	beego.BConfig.Listen.HTTPAddr = host
	logs.Info("host: %s", host)

	beego.BConfig.Listen.HTTPPort = port
	logs.Info("port: %d", port)
}
