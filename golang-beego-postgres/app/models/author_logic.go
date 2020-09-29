package models

import (
	"context"

	"github.com/astaxie/beego/orm"

	"sketch-tech/golang-beego-postgres/app/data"
	"sketch-tech/golang-beego-postgres/app/service"
)

func authorLogicMiddleware() service.AuthorModelMiddleware {
	return func(service.AuthorModel) service.AuthorModel {
		return &authorLogic{}
	}
}

type authorLogic struct{}

func (as authorLogic) Store(ctxIn context.Context, in *data.Author) (ctxOut context.Context, out *data.Author, err error) {
	// update context
	ctxOut = ctxIn

	// update author
	out = in

	// create ormer
	o := orm.NewOrm()

	// read whether name exists
	err = o.QueryTable("author").Filter("author_name", in.Name).One(out)
	if err == nil {
		return
	}

	// insert author
	_, err = o.Insert(out)
	return
}
