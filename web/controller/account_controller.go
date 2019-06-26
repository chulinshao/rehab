package controller

import (
	"github.com/chulinshao/rehab/models"
	"github.com/chulinshao/rehab/service"
)

type AccountController struct {
	Service service.AccountService
}

func (c AccountController) GetByCode(code string) models.DoctorAccount {
	return c.Service.GetByCode(code)
}
