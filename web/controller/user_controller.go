package controller

import "github.com/chulinshao/rehab/service"

type UserController struct {
	Service service.UserService
}

func (c UserController) GetAll() string {
	return c.Service.ListAll()
}
