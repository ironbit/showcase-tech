package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// AuthorController handles all request related with Author.
// beego.NSRouter resets the field of controller #3300
type AuthorController struct {
	beego.Controller
}

// Store is used to store author information in the database.
// beego.NSRouter resets the field of controller #3300
func (ac *AuthorController) Store() {
	// initialize parameters
	logger := logs.NewLogger()
	timeout := timeOutValue

	// create service
	service := authorService{}

	// execute service
	service.store(timeout, logger, requestCounter, &ac.Controller)
}
