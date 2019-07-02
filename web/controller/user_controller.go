package controller

import (
	"github.com/chulinshao/rehab/common"
	"github.com/chulinshao/rehab/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	Service service.UserService
}

func (ctl UserController) ListAll(c *gin.Context) {
	c.JSON(http.StatusOK, common.Result{Data: ctl.Service.ListAll()})
}

func (ctl UserController) GetById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, common.Result{Data: ctl.Service.FindById(id)})
}
