package controllers

import (
	"context"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	uuid "github.com/satori/go.uuid"

	"sketch-tech/golang-beego-postgres/app/service"
)

func authorLoggerMiddleware(l *logs.BeeLogger) service.AuthorControllerMiddleware {
	return func(s service.AuthorController) service.AuthorController {
		return &authorLogger{l, s}
	}
}

type authorLogger struct {
	logger  *logs.BeeLogger
	service service.AuthorController
}

func (as *authorLogger) Store(ctxIn context.Context, controller *beego.Controller) (ctxOut context.Context, outcome interface{}, err error) {
	// read request id
	id := ctxIn.Value(service.IDKey).(uuid.UUID)

	// log at the start of the request
	as.logger.Info("Author | Store | ID: %s", id.String())

	// log of the end of the request
	defer func(id string) {
		if err != nil {
			as.logger.Error("Author | Store | ID: %s | %v", id, err)
		} else {
			// request counter
			counter := ctxOut.Value(service.RequestCounterKey).(int64)
			as.logger.Info("Author | Store | ID: %s | requests: %d", id, counter)

			// elapsed time
			elapsed := ctxOut.Value(service.TimeElapsedKey).(int64)
			as.logger.Info("Author | Store | ID: %s | time: %d µs", id, elapsed/int64(time.Microsecond))
		}
	}(id.String())

	// execute service
	return as.service.Store(ctxIn, controller)
}
