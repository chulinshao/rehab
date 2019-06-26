package service

import (
	"github.com/chulinshao/rehab/dao"
	"github.com/chulinshao/rehab/datasource"
	"github.com/chulinshao/rehab/models"
)

type AccountService interface {
	GetByCode(code string) models.DoctorAccount
}

func NewAccountService() AccountService {
	return &accountService{
		dao: dao.NewAccountDao(datasource.GetEngine()),
	}
}

type accountService struct {
	dao *dao.AccountDao
}

func (s accountService) GetByCode(code string) models.DoctorAccount {
	return s.dao.GetByCode(code)
}
