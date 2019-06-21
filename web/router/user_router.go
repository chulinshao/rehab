package router

import (
	"github.com/chulinshao/rehab/service"
	"github.com/chulinshao/rehab/web/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func ConfigUserRouter(app *iris.Application) {

	//mvc.Configure(app.Party("/user"), func(app *mvc.Application) {
	//	userService := service.NewUserService()
	//	app.Register(userService)
	//	app.Handle(new(controller.UserController))
	//})
	user := mvc.New(app.Party("/user"))
	user.Register(service.NewUserService())
	user.Handle(new(controller.UserController))
}
