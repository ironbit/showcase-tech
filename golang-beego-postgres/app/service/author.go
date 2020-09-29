package service

import (
	"context"

	"github.com/astaxie/beego"

	"sketch-tech/golang-beego-postgres/app/data"
)

// AuthorControllerMiddleware is used to create the services using the
// decorator pattern for the controller package.
type AuthorControllerMiddleware func(AuthorController) AuthorController

// AuthorController establishes the functionality in the controller package.
type AuthorController interface {
	Store(ctxIn context.Context, controller *beego.Controller) (ctxOut context.Context, outcome interface{}, err error)
}

// AuthorModelMiddleware is used to create the services using the
// decorator pattern for the model package.
type AuthorModelMiddleware func(AuthorModel) AuthorModel

// AuthorModel establishes the functionality in the model package.
type AuthorModel interface {
	Store(ctxIn context.Context, in *data.Author) (ctxOut context.Context, out *data.Author, err error)
}
