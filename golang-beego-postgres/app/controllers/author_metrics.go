package controllers

import (
	"context"
	"time"

	"github.com/astaxie/beego"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/generic"

	"sketch-tech/golang-beego-postgres/app/service"
)

func authorMetricsMiddleware(c metrics.Counter) service.AuthorControllerMiddleware {
	return func(s service.AuthorController) service.AuthorController {
		return &authorMetrics{c, s}
	}
}

type authorMetrics struct {
	counter metrics.Counter
	service service.AuthorController
}

func (as *authorMetrics) Store(ctxIn context.Context, controller *beego.Controller) (ctxOut context.Context, outcome interface{}, err error) {
	defer func() {
		// gather time values
		start := ctxOut.Value(service.TimeInitialKey).(int64)
		elapsed := time.Now().UnixNano() - start

		// store time values
		ctxOut = context.WithValue(ctxOut, service.TimeElapsedKey, elapsed)

		// update request counter
		as.counter.Add(1.0)
		value := as.counter.(*generic.Counter).Value()
		ctxOut = context.WithValue(ctxOut, service.RequestCounterKey, int64(value))
	}()

	return as.service.Store(ctxIn, controller)
}
