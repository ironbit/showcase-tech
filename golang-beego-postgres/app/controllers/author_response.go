package controllers

import (
	"context"
	"net/http"

	"github.com/astaxie/beego"

	"sketch-tech/golang-beego-postgres/app/service"
)

func authorResponseMiddleware() service.AuthorControllerMiddleware {
	return func(s service.AuthorController) service.AuthorController {
		return &authorResponse{s}
	}
}

type authorResponse struct {
	service service.AuthorController
}

func (as *authorResponse) Store(ctxIn context.Context, controller *beego.Controller) (ctxOut context.Context, outcome interface{}, err error) {
	defer func() {
		if err != nil {
			msg := "API Error"
			http.Error(controller.Ctx.ResponseWriter, msg, http.StatusBadRequest)
		}
	}()

	return as.service.Store(ctxIn, controller)
}
