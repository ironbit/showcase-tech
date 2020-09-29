package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/go-kit/kit/metrics"

	"sketch-tech/golang-beego-postgres/app/service"
)

func newAuthorService(
	logger *logs.BeeLogger,
	counter metrics.Counter,
	model service.AuthorModel) service.AuthorController {
	// define service variable
	var service service.AuthorController

	// logic middleware
	service = authorLogicMiddleware(model)(service)

	// response middleware
	service = authorResponseMiddleware()(service)

	// metrics middleware
	service = authorMetricsMiddleware(counter)(service)

	// logger middleware
	service = authorLoggerMiddleware(logger)(service)

	// service - decorator pattern
	return service
}
