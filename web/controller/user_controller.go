package controller

import "github.com/chulinshao/rehab/service"

type UserController struct {
	service service.UserService
}

func (c UserController) GetAll() string {
	return c.service.GetAll()
}
