package models

import (
	"github.com/astaxie/beego/logs"

	"sketch-tech/golang-beego-postgres/app/service"
)

// NewAuthorService creates the author model.
func NewAuthorService(logger *logs.BeeLogger) service.AuthorModel {
	var service service.AuthorModel

	// logic middleware
	service = authorLogicMiddleware()(service)

	// logging middleware
	service = authorLoggerMiddleware(logger)(service)

	// service - decorator pattern
	return service
}
