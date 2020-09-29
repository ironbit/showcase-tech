package controllers

import (
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

const timeOutKey = "request::timeout"

var timeOutValue time.Duration

func init() {
	t, err := beego.AppConfig.Int64(timeOutKey)
	if err != nil {
		logs.Critical("controllers | init | cannot read: %v", err)
		os.Exit(1)
	}
	timeOutValue = time.Duration(t) * time.Second
}
