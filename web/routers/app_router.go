package routers

import (
	"github.com/chulinshao/rehab/service"
	"github.com/chulinshao/rehab/web/controller"
	"github.com/gin-gonic/gin"
)

// user router
func configUserRouter(r *gin.Engine) {
	ctl := controller.UserController{Service: service.NewUserService()}
	user := r.Group("/user")
	{
		user.GET("/all", ctl.ListAll)
		user.GET("/id/:id", ctl.GetById)
	}
}

// account router
func configAccountRouter(r *gin.Engine) {
	ctl := controller.AccountController{Service: service.NewAccountService()}
	account := r.Group("/account")
	{
		account.GET("/code/:code", ctl.GetByCode)
		account.POST("/updateAlipayAccount", ctl.UpdateAlipayAccount)
	}
}

func ConfigRouter(r *gin.Engine) {
	configUserRouter(r)
	configAccountRouter(r)
}
