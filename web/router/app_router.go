package router

import (
	"github.com/chulinshao/rehab/service"
	"github.com/chulinshao/rehab/web/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func configUserRouter(app *iris.Application) {
	user := mvc.New(app.Party("/user"))
	user.Register(service.NewUserService())
	user.Handle(new(controller.UserController))
}

func configAccountRouter(app *iris.Application) {
	account := mvc.New(app.Party("/account"))
	account.Register(service.NewAccountService())
	account.Handle(new(controller.AccountController))
}

func ConfigRouter(app *iris.Application) {
	configUserRouter(app)
	configAccountRouter(app)
}
