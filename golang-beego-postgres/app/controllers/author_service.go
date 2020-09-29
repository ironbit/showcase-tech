package controllers

import (
	"context"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/go-kit/kit/metrics"
	uuid "github.com/satori/go.uuid"

	"sketch-tech/golang-beego-postgres/app/models"
	"sketch-tech/golang-beego-postgres/app/service"
)

type authorService struct{}

func (as *authorService) store(
	timeout time.Duration,
	logger *logs.BeeLogger,
	counter metrics.Counter,
	controller *beego.Controller) {

	// start time
	start := time.Now().UnixNano()

	// create service
	model := models.NewAuthorService(logger)
	authorService := newAuthorService(logger, counter, model)

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// store UUID
	ctx = context.WithValue(ctx, service.IDKey, uuid.NewV4())

	// store start time
	ctx = context.WithValue(ctx, service.TimeInitialKey, start)

	// create channel
	done := make(chan struct{})

	// goroutine service execution
	go func() {
		authorService.Store(ctx, controller)
		done <- struct{}{}
	}()

	// wait for timeout result whether happen.
	select {
	case <-done:
	case <-ctx.Done():
		logger.Error("Author | Store | Timeout: %d", timeout)
	}
}
