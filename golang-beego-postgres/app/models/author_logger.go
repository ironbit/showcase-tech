package models

import (
	"context"

	"github.com/astaxie/beego/logs"
	uuid "github.com/satori/go.uuid"

	"sketch-tech/golang-beego-postgres/app/data"
	"sketch-tech/golang-beego-postgres/app/service"
)

func authorLoggerMiddleware(l *logs.BeeLogger) service.AuthorModelMiddleware {
	return func(s service.AuthorModel) service.AuthorModel {
		return &authorLogger{l, s}
	}
}

type authorLogger struct {
	logger  *logs.BeeLogger
	service service.AuthorModel
}

func (as *authorLogger) Store(ctxIn context.Context, in *data.Author) (ctxOut context.Context, out *data.Author, err error) {
	// retrieve request id
	id := ctxIn.Value(service.IDKey).(uuid.UUID)

	// log of the start of the request
	as.logger.Info("Author | Model | Store | id: %s | name: %s", id.String(), in.Name)

	// log of the end of the request
	defer func(id string) {
		if err != nil {
			as.logger.Error("Author | Model | Store | id: %s | error : %v", id, err)
		} else {
			as.logger.Info("Author | Model | Store | id: %s | author: %v", id, out)
		}
	}(id.String())

	// execute request
	return as.service.Store(ctxIn, in)
}
