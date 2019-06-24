package controller

import (
	"github.com/chulinshao/rehab/models"
	"github.com/chulinshao/rehab/service"
)

type UserController struct {
	Service service.UserService
}

func (c UserController) GetAll() []models.User {
	return c.Service.ListAll()
}

func (c UserController) GetById(id string) models.User {
	return c.Service.FindById(id)
}
