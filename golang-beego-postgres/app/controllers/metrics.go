package controllers

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/generic"
)

const requestCounterName = "Request Counter"

var requestCounter metrics.Counter

func init() {
	requestCounter = generic.NewCounter(requestCounterName)
}
