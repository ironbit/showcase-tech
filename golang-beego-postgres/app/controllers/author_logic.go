package controllers

import (
	"context"
	"encoding/json"

	"github.com/astaxie/beego"

	"sketch-tech/golang-beego-postgres/app/data"
	"sketch-tech/golang-beego-postgres/app/service"
)

func authorLogicMiddleware(m service.AuthorModel) service.AuthorControllerMiddleware {
	return func(service.AuthorController) service.AuthorController {
		return &authorLogic{m}
	}
}

type authorLogic struct {
	model service.AuthorModel
}

func (a *authorLogic) Store(ctxIn context.Context, controller *beego.Controller) (ctxOut context.Context, outcome interface{}, err error) {
	var author data.Author
	err = json.Unmarshal(controller.Ctx.Input.RequestBody, &author)
	if err != nil {
		return ctxIn, nil, err
	}

	defer func() {
		if err == nil {
			// casting to Author
			a := outcome.(*data.Author)

			// json marshal
			b, err := json.Marshal(a)
			if err != nil {
				return
			}

			// make the response
			controller.Data["json"] = string(b)
			controller.ServeJSON()
		}
	}()

	return a.model.Store(ctxIn, &author)
}
